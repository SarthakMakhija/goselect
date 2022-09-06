package executor

import (
	"goselect/parser"
	"goselect/parser/context"
	"io/fs"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

type SelectQueryExecutor struct {
	options *Options
	query   *parser.SelectQuery
	context *context.ParsingApplicationContext
}

func NewSelectQueryExecutor(query *parser.SelectQuery, context *context.ParsingApplicationContext, options *Options) *SelectQueryExecutor {
	return &SelectQueryExecutor{
		query:   query,
		context: context,
		options: options,
	}
}

func (selectQueryExecutor *SelectQueryExecutor) Execute() (*EvaluatingRows, error) {
	source := selectQueryExecutor.query.Source

	var limit uint32 = math.MaxInt32
	if selectQueryExecutor.query.IsLimitDefined() {
		limit = selectQueryExecutor.query.Limit.Limit
	}
	rows, err := selectQueryExecutor.executeFrom(source.Directory, limit)
	if err != nil {
		return nil, err
	}
	newOrdering(selectQueryExecutor.query.Order).doOrder(rows)
	return rows, nil
}

func (selectQueryExecutor SelectQueryExecutor) executeFrom(directory string, maxLimit uint32) (*EvaluatingRows, error) {
	var pathSeparator = string(os.PathSeparator)
	var execute func(directory string) error

	shouldTraverseDirectory := func(file fs.FileInfo) bool {
		return file.IsDir() &&
			selectQueryExecutor.options.traverseNestedDirectories &&
			!selectQueryExecutor.options.IsDirectoryTraversalIgnored(file.Name())
	}

	rows := emptyRows(selectQueryExecutor.context.AllFunctions(), maxLimit)
	execute = func(directory string) error {
		files, err := ioutil.ReadDir(directory)
		if err != nil {
			return err
		}
		for _, file := range files {
			if shouldTraverseDirectory(file) {
				newPath := directory + pathSeparator + file.Name()
				if strings.HasSuffix(directory, pathSeparator) {
					newPath = directory + file.Name()
				}
				if err := execute(newPath); err != nil {
					return err
				}
			}
			if rows.Count() >= maxLimit && !selectQueryExecutor.query.IsOrderDefined() {
				return nil
			}
			fileAttributes := context.ToFileAttributes(directory, file, selectQueryExecutor.context)
			shouldChoose, err := selectQueryExecutor.shouldChoose(fileAttributes)
			if err != nil {
				return err
			}
			if shouldChoose {
				values, fullyEvaluated, expressions, err := selectQueryExecutor.query.Projections.EvaluateWith(
					fileAttributes,
					selectQueryExecutor.context.AllFunctions(),
				)
				if err != nil {
					return err
				}
				rows.addRow(values, fullyEvaluated, expressions)
			}
		}
		return nil
	}
	if err := execute(directory); err != nil {
		return nil, err
	}
	return rows, nil
}

func (selectQueryExecutor SelectQueryExecutor) shouldChoose(fileAttributes *context.FileAttributes) (bool, error) {
	if selectQueryExecutor.query.IsWhereDefined() {
		if passesWhere, err := selectQueryExecutor.query.Where.EvaluateWith(fileAttributes, selectQueryExecutor.context.AllFunctions()); err != nil {
			return false, err
		} else if passesWhere {
			return true, nil
		} else {
			return false, nil
		}
	}
	return true, nil
}

package executor

import (
	"goselect/parser"
	"goselect/parser/context"
	"io/fs"
	"math"
	"os"
	"strings"
)

const pathSeparator = string(os.PathSeparator)

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
	if selectQueryExecutor.query.Projections.HasAllAggregates() {
		limit = 1
	} else {
		if selectQueryExecutor.query.IsLimitDefined() {
			limit = selectQueryExecutor.query.Limit.Limit
		}
	}
	rows, err := selectQueryExecutor.executeFrom(source.Directory, limit)
	if err != nil {
		return nil, err
	}
	newOrdering(selectQueryExecutor.query.Order).doOrder(rows)
	return rows, nil
}

func (selectQueryExecutor SelectQueryExecutor) executeFrom(directory string, maxLimit uint32) (*EvaluatingRows, error) {
	rows := emptyRows(selectQueryExecutor.context.AllFunctions(), maxLimit)
	if err := selectQueryExecutor.execute(directory, maxLimit, rows); err != nil {
		return nil, err
	}
	return rows, nil
}

func (selectQueryExecutor SelectQueryExecutor) execute(directory string, maxLimit uint32, rows *EvaluatingRows) error {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		file, err := entry.Info()
		if err != nil {
			return err
		}
		if selectQueryExecutor.shouldTraverseDirectory(file) {
			newPath := selectQueryExecutor.childDirectoryName(directory, entry)
			if err := selectQueryExecutor.execute(newPath, maxLimit, rows); err != nil {
				return err
			}
		}
		if selectQueryExecutor.haveCollectedEnough(rows, maxLimit) {
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

func (selectQueryExecutor SelectQueryExecutor) shouldTraverseDirectory(file fs.FileInfo) bool {
	return file.IsDir() &&
		selectQueryExecutor.options.traverseNestedDirectories &&
		!selectQueryExecutor.options.IsDirectoryTraversalIgnored(file.Name())

}

func (selectQueryExecutor SelectQueryExecutor) childDirectoryName(directory string, entry os.DirEntry) string {
	newPath := directory + pathSeparator + entry.Name()
	if strings.HasSuffix(directory, pathSeparator) {
		newPath = directory + entry.Name()
	}
	return newPath
}

func (selectQueryExecutor SelectQueryExecutor) haveCollectedEnough(rows *EvaluatingRows, maxLimit uint32) bool {
	return rows.Count() >= maxLimit &&
		!selectQueryExecutor.query.IsOrderDefined() &&
		selectQueryExecutor.query.Projections.AggregationCount() == 0
}

func (selectQueryExecutor SelectQueryExecutor) shouldChoose(fileAttributes *context.FileAttributes) (bool, error) {
	if passesWhere, err := selectQueryExecutor.query.Where.EvaluateWith(fileAttributes, selectQueryExecutor.context.AllFunctions()); err != nil {
		return false, err
	} else if passesWhere {
		return true, nil
	} else {
		return false, nil
	}
}

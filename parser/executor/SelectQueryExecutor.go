package executor

import (
	"goselect/parser"
	"goselect/parser/context"
	"io/ioutil"
	"math"
)

type SelectQueryExecutor struct {
	query   *parser.SelectQuery
	context *context.ParsingApplicationContext
}

func NewSelectQueryExecutor(query *parser.SelectQuery, context *context.ParsingApplicationContext) *SelectQueryExecutor {
	return &SelectQueryExecutor{
		query:   query,
		context: context,
	}
}

func (selectQueryExecutor *SelectQueryExecutor) Execute() (*EvaluatingRows, error) {
	source := selectQueryExecutor.query.Source
	files, err := ioutil.ReadDir(source.Directory)
	if err != nil {
		return nil, err
	}

	var limit uint32 = math.MaxInt32
	if selectQueryExecutor.query.IsLimitDefined() {
		limit = selectQueryExecutor.query.Limit.Limit
	}

	var rowCount uint32 = 0
	rows := emptyRows(selectQueryExecutor.context.AllFunctions())

	for _, file := range files {
		if rowCount >= limit && !selectQueryExecutor.query.IsOrderDefined() {
			break
		}
		fileAttributes := context.ToFileAttributes(file, selectQueryExecutor.context)
		shouldChoose, err := selectQueryExecutor.shouldChoose(fileAttributes)
		if err != nil {
			return nil, err
		}
		if shouldChoose {
			values, fullyEvaluated, expressions, err := selectQueryExecutor.query.Projections.EvaluateWith(fileAttributes, selectQueryExecutor.context.AllFunctions())
			if err != nil {
				return nil, err
			}
			rows.addRow(values, fullyEvaluated, expressions)
			rowCount = rowCount + 1
		}
		//handle recursion
	}
	newOrdering(selectQueryExecutor.query.Order).doOrder(rows)
	//handle limit
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

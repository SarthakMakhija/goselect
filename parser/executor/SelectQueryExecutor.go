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

func (selectQueryExecutor *SelectQueryExecutor) Execute() ([][]context.Value, error) {
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
	var rows [][]context.Value
	for _, file := range files {
		//assume no order by
		if rowCount >= limit {
			break
		}
		fileAttributes := context.ToFileAttributes(file, selectQueryExecutor.context)
		row, err := selectQueryExecutor.query.Projections.EvaluateWith(fileAttributes, selectQueryExecutor.context.AllFunctions())
		if err != nil {
			return nil, err
		}
		rows, rowCount = append(rows, row), rowCount+1
		//handle recursion
	}
	return rows, nil
}

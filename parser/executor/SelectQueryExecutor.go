package executor

import (
	"goselect/parser"
	"goselect/parser/context"
	"io/ioutil"
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

func (selectQueryExecutor *SelectQueryExecutor) Execute() ([][]interface{}, error) {
	source := selectQueryExecutor.query.Source
	files, err := ioutil.ReadDir(source.Directory)
	if err != nil {
		return nil, err
	}

	var rows [][]interface{}
	for _, file := range files {
		fileAttributes := context.ToFileAttributes(file, selectQueryExecutor.context)
		row, err := selectQueryExecutor.query.Projections.EvaluateWith(fileAttributes, selectQueryExecutor.context.AllFunctions())
		if err != nil {
			return nil, err
		}
		rows = append(rows, row)
		//handle recursion
	}
	return rows, nil
}

package executor

import (
	"goselect/parser"
	"goselect/parser/context"
	"io/ioutil"
)

type SelectQueryExecutor struct {
	query   *parser.SelectQuery
	context *context.Context
}

func NewSelectQueryExecutor(query *parser.SelectQuery, context *context.Context) *SelectQueryExecutor {
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
		row := selectQueryExecutor.query.Projections.EvaluateWith(fileAttributes, selectQueryExecutor.context.AllFunctions())
		rows = append(rows, row)
		//handle recursion
	}
	return rows, nil
}

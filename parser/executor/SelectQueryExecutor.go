package executor

import (
	"goselect/parser"
	"goselect/parser/projection"
	"io/ioutil"
)

type SelectQueryExecutor struct {
	query *parser.SelectQuery
}

func NewSelectQueryExecutor(query *parser.SelectQuery) *SelectQueryExecutor {
	return &SelectQueryExecutor{
		query: query,
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
		fileAttributes := projection.FromFile(file)
		row := selectQueryExecutor.query.Projections.EvaluateWith(fileAttributes)
		rows = append(rows, row)
		//handle recursion
	}
	return rows, nil
}

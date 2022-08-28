package executor

import (
	"goselect/parser"
	"goselect/parser/projection"
	"io/ioutil"
)

func ExecuteSelect(query *parser.SelectQuery) [][]interface{} {
	source := query.Source
	files, err := ioutil.ReadDir(source.Directory)
	if err != nil {
		return nil
	}
	var rows [][]interface{}
	for _, file := range files {
		fileAttributes := projection.FromFile(file)
		row := query.Projections.AllExpressions().ExecuteWith(fileAttributes)
		rows = append(rows, row)
		//handle recursion
	}
	return rows
}

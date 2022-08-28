package executor

import (
	"fmt"
	"goselect/parser"
	"goselect/parser/projection"
	"io/ioutil"
)

func ExecuteSelect(query *parser.SelectQuery) ([][]interface{}, error) {
	source := query.Source
	files, err := ioutil.ReadDir(source.Directory)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var rows [][]interface{}
	for _, file := range files {
		fmt.Println("file name ", file.Name())
		fileAttributes := projection.FromFile(file)
		row := query.Projections.AllExpressions().ExecuteWith(fileAttributes)
		rows = append(rows, row)
		//handle recursion
	}
	return rows, nil
}

package executor

import (
	"goselect/parser/context"
	"goselect/parser/projection"
)

type EvaluatingRows struct {
	rows      []*EvaluatingRow
	functions *context.AllFunctions
}

type RowsIterator struct {
	currentIndex int
	rows         []*EvaluatingRow
}

type EvaluatingRow struct {
	attributeValues []context.Value
	fullyEvaluated  []bool
	expressions     []*projection.Expression
	functions       *context.AllFunctions
}

type AttributeIterator struct {
	currentIndex    int
	attributeValues []context.Value
	fullyEvaluated  []bool
	expressions     []*projection.Expression
}

func emptyRows(functions *context.AllFunctions) *EvaluatingRows {
	return &EvaluatingRows{functions: functions}
}

func (rows *EvaluatingRows) addRow(attributeValues []context.Value, fullyEvaluated []bool, expressions []*projection.Expression) *EvaluatingRow {
	row := &EvaluatingRow{
		attributeValues: attributeValues,
		fullyEvaluated:  fullyEvaluated,
		expressions:     expressions,
		functions:       rows.functions,
	}
	rows.rows = append(rows.rows, row)
	return row
}

func (rows EvaluatingRows) Count() int {
	return len(rows.rows)
}

func (rows *EvaluatingRows) RowIterator() *RowsIterator {
	return &RowsIterator{currentIndex: 0, rows: rows.rows}
}

func (rows EvaluatingRows) atIndex(index int) *EvaluatingRow {
	if index < len(rows.rows) {
		return rows.rows[index]
	}
	return &EvaluatingRow{}
}

func (rowsIterator *RowsIterator) HasNext() bool {
	return rowsIterator.currentIndex < len(rowsIterator.rows)
}

func (rowsIterator *RowsIterator) Next() *EvaluatingRow {
	row := rowsIterator.rows[rowsIterator.currentIndex]
	rowsIterator.currentIndex = rowsIterator.currentIndex + 1
	return row
}

func (row *EvaluatingRow) AllAttributes() []context.Value {
	var values []context.Value
	for index, attributeValue := range row.attributeValues {
		if row.fullyEvaluated[index] {
			values = append(values, attributeValue)
		} else {
			value, err := row.expressions[index].FullyEvaluate(row.functions)
			if err != nil {
				value = context.StringValue(err.Error())
			}
			values = append(values, value)
			row.attributeValues[index] = value
			row.fullyEvaluated[index] = true
		}
	}
	return values
}

func (row EvaluatingRow) TotalAttributes() int {
	return len(row.attributeValues)
}

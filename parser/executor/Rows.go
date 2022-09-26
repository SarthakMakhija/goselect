package executor

import (
	"goselect/parser/context"
	"goselect/parser/expression"
)

type EvaluatingRows struct {
	rows      []*EvaluatingRow
	functions *context.AllFunctions
	limit     uint32
}

type RowsIterator struct {
	currentIndex uint32
	limit        uint32
	rows         []*EvaluatingRow
}

type EvaluatingRow struct {
	attributeValues []context.Value
	fullyEvaluated  []bool
	expressions     []*expression.Expression
	functions       *context.AllFunctions
}

func emptyRows(functions *context.AllFunctions, limit uint32) *EvaluatingRows {
	return &EvaluatingRows{functions: functions, limit: limit}
}

func (rows *EvaluatingRows) addRow(attributeValues []context.Value, fullyEvaluated []bool, expressions []*expression.Expression) *EvaluatingRow {
	row := &EvaluatingRow{
		attributeValues: attributeValues,
		fullyEvaluated:  fullyEvaluated,
		expressions:     expressions,
		functions:       rows.functions,
	}
	rows.rows = append(rows.rows, row)
	return row
}

func (rows EvaluatingRows) Count() uint32 {
	minOf := func(a, b uint32) uint32 {
		if a < b {
			return a
		}
		return b
	}
	return minOf(uint32(len(rows.rows)), rows.limit)
}

func (rows *EvaluatingRows) RowIterator() *RowsIterator {
	return &RowsIterator{currentIndex: 0, limit: rows.limit, rows: rows.rows}
}

func (rows EvaluatingRows) AtIndex(index int) *EvaluatingRow {
	if index < len(rows.rows) {
		return rows.rows[index]
	}
	return &EvaluatingRow{}
}

func (rowsIterator *RowsIterator) HasNext() bool {
	return rowsIterator.currentIndex < uint32(len(rowsIterator.rows)) &&
		rowsIterator.currentIndex+1 <= rowsIterator.limit
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

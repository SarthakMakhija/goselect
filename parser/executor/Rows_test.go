//go:build unit
// +build unit

package executor

import (
	"goselect/parser/context"
	"goselect/parser/expression"
	"reflect"
	"strings"
	"testing"
)

func TestEvaluatingRowAllAttributesThatAreFullyEvaluated(t *testing.T) {
	_ = context.NewContext(context.NewFunctions(), context.NewAttributes())
	rows := emptyRows(context.NewFunctions(), 1)
	rows.addRow([]context.Value{context.StringValue("someValue")}, []bool{true}, []*expression.Expression{})

	attributes := rows.atIndex(0).AllAttributes()
	expected := []context.Value{context.StringValue("someValue")}

	if !reflect.DeepEqual(expected, attributes) {
		t.Fatalf("Expected attributes to be %v, received %v", expected, attributes)
	}
}

func TestEvaluatingRowAtAnIndexGreaterThanTotalNumberOfRows(t *testing.T) {
	_ = context.NewContext(context.NewFunctions(), context.NewAttributes())
	rows := emptyRows(context.NewFunctions(), 1)
	rows.addRow([]context.Value{context.StringValue("someValue")}, []bool{true}, []*expression.Expression{})

	row := rows.atIndex(1)
	if len(row.attributeValues) != 0 {
		t.Fatalf("Expected no attributes given row is accessed at an index beyond the total number of rows")
	}
}

func TestEvaluatingRowCount(t *testing.T) {
	_ = context.NewContext(context.NewFunctions(), context.NewAttributes())
	rows := emptyRows(context.NewFunctions(), 1)
	rows.addRow([]context.Value{context.StringValue("someValue")}, []bool{true}, []*expression.Expression{})

	count := rows.Count()
	expected := uint32(1)

	if count != expected {
		t.Fatalf("Expected count to be %v, received %v", expected, count)
	}
}

func TestEvaluatingRowIterator(t *testing.T) {
	_ = context.NewContext(context.NewFunctions(), context.NewAttributes())
	rows := emptyRows(context.NewFunctions(), 1)
	rows.addRow([]context.Value{context.StringValue("someValue")}, []bool{true}, []*expression.Expression{})

	attributes := rows.RowIterator().Next().AllAttributes()
	expected := []context.Value{context.StringValue("someValue")}

	if !reflect.DeepEqual(expected, attributes) {
		t.Fatalf("Expected attributes to be %v, received %v", expected, attributes)
	}
}

func TestEvaluatingRowIteratorHasNextWithAnAvailableRow(t *testing.T) {
	_ = context.NewContext(context.NewFunctions(), context.NewAttributes())
	rows := emptyRows(context.NewFunctions(), 1)
	rows.addRow([]context.Value{context.StringValue("someValue")}, []bool{true}, []*expression.Expression{})

	hasNext := rows.RowIterator().HasNext()
	if hasNext != true {
		t.Fatalf("Expected a row using row iterator but hasNext returned false")
	}
}

func TestEvaluatingRowIteratorHasNextWithLimit(t *testing.T) {
	_ = context.NewContext(context.NewFunctions(), context.NewAttributes())
	rows := emptyRows(context.NewFunctions(), 2)
	rows.addRow([]context.Value{context.StringValue("someValue")}, []bool{true}, []*expression.Expression{})
	rows.addRow([]context.Value{context.StringValue("someValue")}, []bool{true}, []*expression.Expression{})
	rows.addRow([]context.Value{context.StringValue("someValue")}, []bool{true}, []*expression.Expression{})

	iterator := rows.RowIterator()

	hasNext := iterator.HasNext()
	if hasNext != true {
		t.Fatalf("Expected a row using row iterator but hasNext returned false")
	}
	iterator.Next()

	hasNext = iterator.HasNext()
	if hasNext != true {
		t.Fatalf("Expected a row using row iterator but hasNext returned false")
	}
	iterator.Next()

	hasNext = iterator.HasNext()
	if hasNext != false {
		t.Fatalf("Expected no row given limit is reached but hasNext returned true")
	}
}

func TestEvaluatingRowIteratorHasNextWithNoAvailableRows(t *testing.T) {
	_ = context.NewContext(context.NewFunctions(), context.NewAttributes())
	rows := emptyRows(context.NewFunctions(), 1)

	hasNext := rows.RowIterator().HasNext()
	if hasNext != false {
		t.Fatalf("Expected no rows using row iterator but hasNext returned true")
	}
}

func TestEvaluatingRowTotalAttributes(t *testing.T) {
	_ = context.NewContext(context.NewFunctions(), context.NewAttributes())
	rows := emptyRows(context.NewFunctions(), 1)
	rows.addRow([]context.Value{context.StringValue("someValue")}, []bool{true}, []*expression.Expression{})

	totalAttributes := rows.RowIterator().Next().TotalAttributes()
	if totalAttributes != 1 {
		t.Fatalf("Expected total attrobutes to be %v, received %v", 1, totalAttributes)
	}
}

func TestEvaluatingRowFullyEvaluationOfAnAttribute(t *testing.T) {
	_ = context.NewContext(context.NewFunctions(), context.NewAttributes())
	rows := emptyRows(context.NewFunctions(), 1)
	rows.addRow(
		[]context.Value{context.StringValue("someValue"), context.StringValue("test")},
		[]bool{true, false},
		[]*expression.Expression{
			expression.WithValue("someValue"),
			expression.WithFunctionInstance(expression.FunctionInstanceWith(
				"min",
				nil,
				&context.FunctionState{Initial: context.StringValue("test")},
				true,
			)),
		},
	)

	values := rows.RowIterator().Next().AllAttributes()
	if values[0].GetAsString() != "someValue" {
		t.Fatalf("Expected value after full evaluation to be %v, received %v", "someValue", values[0].GetAsString())
	}

	if !strings.Contains(values[1].GetAsString(), "expected 1 parameter(s) in the function min") {
		t.Fatalf(
			"Expected error string %v to contained in the values[1] but was not, received %v",
			"expected 1 parameter(s) in the function min",
			values[1].GetAsString(),
		)
	}
}

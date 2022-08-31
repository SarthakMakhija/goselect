package executor

import (
	"goselect/parser/context"
	"goselect/parser/projection"
	"reflect"
	"testing"
)

func TestEvaluatingRowAllAttributesThatAreFullyEvaluated(t *testing.T) {
	_ = context.NewContext(context.NewFunctions(), context.NewAttributes())
	rows := emptyRows(context.NewFunctions())
	rows.addRow([]context.Value{context.StringValue("someValue")}, []bool{true}, []*projection.Expression{})

	attributes := rows.atIndex(0).AllAttributes()
	expected := []context.Value{context.StringValue("someValue")}

	if !reflect.DeepEqual(expected, attributes) {
		t.Fatalf("Expected attributes to be %v, received %v", expected, attributes)
	}
}

func TestEvaluatingRowCount(t *testing.T) {
	_ = context.NewContext(context.NewFunctions(), context.NewAttributes())
	rows := emptyRows(context.NewFunctions())
	rows.addRow([]context.Value{context.StringValue("someValue")}, []bool{true}, []*projection.Expression{})

	count := rows.Count()
	expected := 1

	if count != expected {
		t.Fatalf("Expected count to be %v, received %v", expected, count)
	}
}

func TestEvaluatingRowIterator(t *testing.T) {
	_ = context.NewContext(context.NewFunctions(), context.NewAttributes())
	rows := emptyRows(context.NewFunctions())
	rows.addRow([]context.Value{context.StringValue("someValue")}, []bool{true}, []*projection.Expression{})

	attributes := rows.RowIterator().Next().AllAttributes()
	expected := []context.Value{context.StringValue("someValue")}

	if !reflect.DeepEqual(expected, attributes) {
		t.Fatalf("Expected attributes to be %v, received %v", expected, attributes)
	}
}

func TestEvaluatingRowIteratorHasNextWithAnAvailableRow(t *testing.T) {
	_ = context.NewContext(context.NewFunctions(), context.NewAttributes())
	rows := emptyRows(context.NewFunctions())
	rows.addRow([]context.Value{context.StringValue("someValue")}, []bool{true}, []*projection.Expression{})

	hasNext := rows.RowIterator().HasNext()
	if hasNext != true {
		t.Fatalf("Expected a row using row iterator but hasNext returned false")
	}
}

func TestEvaluatingRowIteratorHasNextWithNoAvailableRows(t *testing.T) {
	_ = context.NewContext(context.NewFunctions(), context.NewAttributes())
	rows := emptyRows(context.NewFunctions())

	hasNext := rows.RowIterator().HasNext()
	if hasNext != false {
		t.Fatalf("Expected no rows using row iterator but hasNext returned true")
	}
}

func TestEvaluatingRowTotalAttributes(t *testing.T) {
	_ = context.NewContext(context.NewFunctions(), context.NewAttributes())
	rows := emptyRows(context.NewFunctions())
	rows.addRow([]context.Value{context.StringValue("someValue")}, []bool{true}, []*projection.Expression{})

	totalAttributes := rows.RowIterator().Next().TotalAttributes()
	if totalAttributes != 1 {
		t.Fatalf("Expected total attrobutes to be %v, received %v", 1, totalAttributes)
	}
}

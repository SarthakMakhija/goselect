package executor

import (
	"goselect/parser/context"
	"goselect/parser/order"
	"sort"
)

type Ordering struct {
	order *order.Order
}

func newOrdering(order *order.Order) *Ordering {
	return &Ordering{order: order}
}

func (ordering *Ordering) doOrder(rows *EvaluatingRows) {
	if ordering.order != nil {
		sort.SliceStable(rows.rows, func(i, j int) bool {
			return ordering.isOrdered(rows.atIndex(i).AllAttributes(), rows.atIndex(j).AllAttributes())
		})
	}
}
func (ordering Ordering) isOrdered(first, second []context.Value) bool {
	for index, orderingAttributeRef := range ordering.order.Attributes {
		firstAttributeValue := first[orderingAttributeRef.ProjectionPosition-1]
		secondAttributeValue := second[orderingAttributeRef.ProjectionPosition-1]

		comparisonResult := firstAttributeValue.CompareTo(secondAttributeValue)
		if comparisonResult == 0 {
			continue
		}

		if ordering.order.IsAscendingAt(index) {
			return comparisonResult < 0
		} else {
			return comparisonResult >= 0
		}
	}
	return false
}

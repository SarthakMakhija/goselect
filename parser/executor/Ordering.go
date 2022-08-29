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

func (ordering *Ordering) doOrder(rows [][]context.Value) {
	if ordering.order != nil {
		sort.SliceStable(rows, func(i, j int) bool {
			return ordering.isOrdered(rows[i], rows[j])
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
			if comparisonResult < 0 {
				return true
			}
			return false
		} else {
			if comparisonResult < 0 {
				return false
			}
			return true
		}
	}
	return false
}

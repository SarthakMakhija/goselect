package parser

import (
	"goselect/parser/limit"
	"goselect/parser/order"
	"goselect/parser/where"
	"testing"
)

func TestSelectQueryWithWhereNotDefined(t *testing.T) {
	query := SelectQuery{Where: nil}
	whereDefined := query.IsWhereDefined()

	if whereDefined != false {
		t.Fatalf("Expected where to be undefined but was defined")
	}
}

func TestSelectQueryWithWhereDefined(t *testing.T) {
	query := SelectQuery{Where: &where.Where{}}
	whereDefined := query.IsWhereDefined()

	if whereDefined != true {
		t.Fatalf("Expected where to be defined but was not")
	}
}

func TestSelectQueryWithOrderByNotDefined(t *testing.T) {
	query := SelectQuery{Order: nil}
	orderDefined := query.IsOrderDefined()

	if orderDefined != false {
		t.Fatalf("Expected order by to be undefined but was defined")
	}
}

func TestSelectQueryWithOrderByDefined(t *testing.T) {
	query := SelectQuery{Order: &order.Order{}}
	orderDefined := query.IsOrderDefined()

	if orderDefined != true {
		t.Fatalf("Expected order by to be defined but was not")
	}
}

func TestSelectQueryWithLimitNotDefined(t *testing.T) {
	query := SelectQuery{Limit: nil}
	limitDefined := query.IsLimitDefined()

	if limitDefined != false {
		t.Fatalf("Expected limit to be undefined but was defined")
	}
}

func TestSelectQueryWithLimitDefined(t *testing.T) {
	query := SelectQuery{Limit: &limit.Limit{}}
	limitDefined := query.IsLimitDefined()

	if limitDefined != true {
		t.Fatalf("Expected limit to be defined but was not")
	}
}

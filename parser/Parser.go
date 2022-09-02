package parser

import (
	"errors"
	"goselect/parser/context"
	"goselect/parser/error/messages"
	"goselect/parser/limit"
	"goselect/parser/order"
	"goselect/parser/projection"
	"goselect/parser/source"
	"goselect/parser/tokenizer"
	"goselect/parser/where"
	"strings"
)

type SelectQuery struct {
	Projections *projection.Projections
	Source      *source.Source
	Order       *order.Order
	Limit       *limit.Limit
	Where       *where.Where
}

func (selectQuery *SelectQuery) IsLimitDefined() bool {
	return selectQuery.Limit != nil
}

func (selectQuery *SelectQuery) IsWhereDefined() bool {
	return selectQuery.Where != nil
}

func (selectQuery *SelectQuery) IsOrderDefined() bool {
	return selectQuery.Order != nil
}

type Parser struct {
	query   string
	context *context.ParsingApplicationContext
}

func NewParser(query string, context *context.ParsingApplicationContext) (*Parser, error) {
	if len(query) == 0 {
		return nil, errors.New(messages.ErrorMessageEmptyQuery)
	}
	return &Parser{query: strings.TrimSpace(query), context: context}, nil
}

func (parser *Parser) Parse() (*SelectQuery, error) {
	tokens := tokenizer.NewTokenizer(parser.query).Tokenize()
	iterator := tokens.Iterator()

	if iterator.HasNext() && !iterator.Peek().Equals("select") {
		return nil, errors.New(messages.ErrorMessageNonSelectQuery)
	}
	projections, err := projection.NewProjections(iterator, parser.context)
	if err != nil {
		return nil, err
	}
	fileSource, err := source.NewSource(iterator)
	if err != nil {
		return nil, err
	}
	whereClause, err := where.NewWhere(iterator, parser.context)
	if err != nil {
		return nil, err
	}
	orderBy, err := order.NewOrder(iterator, projections.Count())
	if err != nil {
		return nil, err
	}
	limitResults, err := limit.NewLimit(iterator)
	if err != nil {
		return nil, err
	}

	return &SelectQuery{
		Projections: projections,
		Source:      fileSource,
		Where:       whereClause,
		Order:       orderBy,
		Limit:       limitResults,
	}, nil
}

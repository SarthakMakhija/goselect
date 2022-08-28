package parser

import (
	"errors"
	"goselect/parser/error/messages"
	"goselect/parser/limit"
	"goselect/parser/order"
	"goselect/parser/projection"
	"goselect/parser/source"
	"goselect/parser/tokenizer"
	"strings"
)

type SelectQuery struct {
	Projections *projection.Projections
	Source      *source.Source
	Order       *order.Order
	Limit       *limit.Limit
}

type Parser struct {
	query string
}

func NewParser(query string) (*Parser, error) {
	if len(query) == 0 {
		return nil, errors.New(messages.ErrorMessageEmptyQuery)
	}
	return &Parser{query: strings.TrimSpace(strings.ToLower(query))}, nil
}

func (parser *Parser) Parse() (*SelectQuery, error) {
	tokens := tokenizer.NewTokenizer(parser.query).Tokenize()
	iterator := tokens.Iterator()

	if iterator.HasNext() && !iterator.Peek().Equals("select") {
		return nil, errors.New(messages.ErrorMessageNonSelectQuery)
	}
	projections, err := projection.NewProjections(iterator)
	if err != nil {
		return nil, err
	}
	fileSource, err := source.NewSource(iterator)
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
		Order:       orderBy,
		Limit:       limitResults,
	}, nil
}

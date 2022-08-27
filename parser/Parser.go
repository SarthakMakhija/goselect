package parser

import (
	"goselect/parser/limit"
	"goselect/parser/order"
	"goselect/parser/projection"
	"goselect/parser/source"
	"goselect/parser/tokenizer"
	"strings"
)

type SelectQuery struct {
	projections *projection.Projections
	source      *source.Source
	order       *order.Order
	limit       *limit.Limit
}

type Parser struct {
	query string
}

func NewParser(query string) *Parser {
	return &Parser{query: query}
}

func (parser *Parser) Parse() (*SelectQuery, error) {
	tokens := tokenizer.NewTokenizer(strings.ToLower(parser.query)).Tokenize()
	iterator := tokens.Iterator()

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
		projections: projections,
		source:      fileSource,
		order:       orderBy,
		limit:       limitResults,
	}, nil
}

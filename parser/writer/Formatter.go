package writer

import (
	"goselect/parser/context"
	"goselect/parser/projection"
)

type Formatter interface {
	Format(projections *projection.Projections, rows [][]context.Value) string
}

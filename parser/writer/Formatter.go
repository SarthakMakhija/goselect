package writer

import (
	"goselect/parser/executor"
	"goselect/parser/projection"
)

type Formatter interface {
	Format(projections *projection.Projections, rows *executor.EvaluatingRows) string
}

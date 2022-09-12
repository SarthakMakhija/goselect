package context

import (
	"github.com/gabriel-vasile/mimetype"
)

type AttributeLazyEvaluationBlock interface {
	evaluate(filePath string) (Value, error)
}

type MimeTypeAttributeEvaluationBlock struct{}

func (m MimeTypeAttributeEvaluationBlock) evaluate(filePath string) (Value, error) {
	mime, err := mimetype.DetectFile(filePath)
	if err != nil {
		return EmptyValue, nil
	}
	return StringValue(mime.String()), nil
}

package context

import (
	"github.com/gabriel-vasile/mimetype"
)

type AttributeLazyEvaluationBlock interface {
	evaluate(filePath string) Value
}

type MimeTypeAttributeEvaluationBlock struct{}

func (m MimeTypeAttributeEvaluationBlock) evaluate(filePath string) Value {
	mime, err := mimetype.DetectFile(filePath)
	if err != nil {
		return StringValue("NA")
	}
	return StringValue(mime.String())
}

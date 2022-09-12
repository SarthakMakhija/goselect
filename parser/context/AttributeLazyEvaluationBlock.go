package context

import (
	"github.com/gabriel-vasile/mimetype"
	"os"
)

type AttributeLazyEvaluationBlock interface {
	evaluate(filePath string) (Value, error)
}

type MimeTypeAttributeEvaluationBlock struct{}

func (m MimeTypeAttributeEvaluationBlock) evaluate(filePath string) (Value, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return EmptyValue, nil
	}
	header := make([]byte, 261)
	_, err = file.Read(header)
	if err != nil {
		return EmptyValue, err
	}
	return StringValue(mimetype.Detect(header).String()), nil
}

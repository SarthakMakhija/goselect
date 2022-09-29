package context

import (
	"github.com/gabriel-vasile/mimetype"
	"os"
	"strings"
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

type ContentsAttributeEvaluationBlock struct{}

func (m ContentsAttributeEvaluationBlock) evaluate(filePath string) Value {
	mimeType, _ := mimetype.DetectFile(filePath)
	if !strings.HasPrefix(mimeType.String(), "text") {
		return StringValue("")
	}
	contents, err := os.ReadFile(filePath)
	if err != nil {
		return StringValue("")
	}

	return StringValue(string(contents))
}

package context

import (
	"github.com/dustin/go-humanize"
	"github.com/gabriel-vasile/mimetype"
	"os"
	"strings"
)

var contentsEvaluationSize = "20 Mib"

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
	if !shouldEvaluateContents(filePath) {
		return StringValue("")
	}
	contents, err := os.ReadFile(filePath)
	if err != nil {
		return StringValue("")
	}

	return StringValue(string(contents))
}

func shouldEvaluateContents(filePath string) bool {
	mimeType, _ := mimetype.DetectFile(filePath)
	if !strings.HasPrefix(mimeType.String(), "text") {
		return false
	}
	fileInfo, errInReadingStats := os.Stat(filePath)
	if errInReadingStats != nil {
		return false
	}
	sizeInBytes, _ := humanize.ParseBytes(contentsEvaluationSize)

	if fileInfo.Size() > int64(sizeInBytes) {
		return false
	}

	return true
}

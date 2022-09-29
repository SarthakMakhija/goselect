package context

import (
	"github.com/gabriel-vasile/mimetype"
	"io"
	"os"
	"strings"
)

const (
	twentyMb = 20 * 1024 * 1024
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

type ContentsAttributeLazyEvaluationBlock struct {
	maxFileSizeInBytesSupported int64
}

func (c ContentsAttributeLazyEvaluationBlock) evaluate(filePath string) Value {
	mime, err := mimetype.DetectFile(filePath)
	if err != nil {
		return blankStringValue
	}

	if strings.HasPrefix(mime.String(), "text/") {
		lstat, err := os.Lstat(filePath)
		if err == nil && lstat.Size() <= c.maxFileSizeInBytesSupported {
			file, err := os.Open(filePath)
			if err != nil {
				return blankStringValue
			}
			defer file.Close()
			bytes, err := io.ReadAll(file)
			if err != nil {
				return blankStringValue
			}
			return StringValue(string(bytes))
		}
	}
	return blankStringValue
}

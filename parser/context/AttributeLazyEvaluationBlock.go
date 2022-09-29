package context

import (
	"github.com/gabriel-vasile/mimetype"
	"io"
	"os"
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

type ContentsAttributeEvaluationBlock struct {
	MaxFileSizeInBytesSupported int64
}

func (m ContentsAttributeEvaluationBlock) evaluate(filePath string) Value {
	mime, err := mimetype.DetectFile(filePath)
	if err != nil {
		return StringValue("")
	}

	if mime.Is("text/plain") {
		lstat, _ := os.Lstat(filePath)
		if err == nil && lstat.Size() <= m.MaxFileSizeInBytesSupported {
			file, err := os.Open(filePath)
			if err != nil {
				return StringValue("")
			}
			defer file.Close()
			bytes, err := io.ReadAll(file)
			if err != nil {
				return StringValue("")
			}
			return StringValue(string(bytes))
		}
	}
	return StringValue("")
}

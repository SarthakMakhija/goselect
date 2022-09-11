package source

import (
	"errors"
	"fmt"
	"goselect/parser/error/messages"
	"goselect/parser/tokenizer"
	"os"
)

type Source struct {
	Directory string
}

func NewSource(tokenIterator *tokenizer.TokenIterator) (*Source, error) {
	if directory, err := getDirectory(tokenIterator); err != nil {
		return nil, err
	} else {
		if file, err := os.Stat(directory); err != nil {
			if os.IsNotExist(err) {
				return nil, fmt.Errorf(messages.ErrorMessageInaccessibleSource, directory)
			} else {
				return nil, err
			}
		} else {
			if !file.IsDir() {
				return nil, errors.New(messages.ErrorMessageSourceNotADirectory)
			}
		}
		return &Source{Directory: directory}, nil
	}
}

func getDirectory(tokenIterator *tokenizer.TokenIterator) (string, error) {
	if tokenIterator.HasNext() && tokenIterator.Peek().Equals("from") {
		tokenIterator.Next()
	}
	if tokenIterator.HasNext() && !tokenIterator.Peek().Equals("where") {
		token := tokenIterator.Next()
		path := token.TokenValue
		return ExpandDirectoryPath(path)
	}
	return "", errors.New(messages.ErrorMessageMissingSource)
}

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
	directory, err := getDirectory(tokenIterator)
	if err != nil {
		return nil, err
	}
	file, err := os.Stat(directory)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf(messages.ErrorMessageInaccessibleSource, directory)
		}
		return nil, err
	}
	if !file.IsDir() {
		return nil, errors.New(messages.ErrorMessageSourceNotADirectory)
	}
	return &Source{Directory: directory}, nil
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

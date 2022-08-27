package source

import (
	"errors"
	"goselect/parser/error/messages"
	"goselect/parser/tokenizer"
	"os/user"
	"strings"
)

type Source struct {
	directory string
}

func NewSource(tokenIterator *tokenizer.TokenIterator) (*Source, error) {
	if directory, err := getDirectory(tokenIterator); err != nil {
		return nil, err
	} else {
		return &Source{directory: directory}, nil
	}
}

func getDirectory(tokenIterator *tokenizer.TokenIterator) (string, error) {
	if tokenIterator.HasNext() && tokenIterator.Peek().Equals("from") {
		tokenIterator.Next()
	}
	if tokenIterator.HasNext() && !tokenIterator.Peek().Equals("where") {
		token := tokenIterator.Next()
		path := token.TokenValue
		if strings.HasPrefix(path, "~") {
			if currentUser, err := user.Current(); err != nil {
				return "", err
			} else {
				directory := currentUser.HomeDir + path[1:]
				return directory, nil
			}
		}
		return path, nil
	}
	return "", errors.New(messages.ErrorMessageMissingSource)
}

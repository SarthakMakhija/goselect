package source

import (
	"errors"
	"fmt"
	"goselect/parser/error/messages"
	"goselect/parser/tokenizer"
	"os"
	"os/user"
	"strings"
)

type Source struct {
	Directory string
}

func NewSource(tokenIterator *tokenizer.TokenIterator) (*Source, error) {
	if directory, err := getDirectory(tokenIterator); err != nil {
		return nil, err
	} else {
		if _, err := os.Stat(directory); err != nil {
			if os.IsNotExist(err) {
				return nil, fmt.Errorf(messages.ErrorMessageInaccessibleSource, directory)
			} else {
				return nil, err
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

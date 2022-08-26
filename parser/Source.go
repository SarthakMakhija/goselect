package parser

import (
	"errors"
	"os/user"
	"strings"
)

type Source struct {
	tokenIterator *TokenIterator
	directory     string
}

func newSource(tokenIterator *TokenIterator) (*Source, error) {
	source := &Source{tokenIterator: tokenIterator}
	if err := source.setDirectory(); err != nil {
		return nil, err
	}
	return source, nil
}

func (source *Source) setDirectory() error {
	if source.tokenIterator.hasNext() && !source.tokenIterator.peek().equals("where") {
		token := source.tokenIterator.next()
		path := token.tokenValue
		if strings.HasPrefix(path, "~") {
			if currentUser, err := user.Current(); err != nil {
				return err
			} else {
				source.directory = currentUser.HomeDir + path[1:]
				return nil
			}
		}
		source.directory = path
		return nil
	}
	return errors.New("incomplete query, need a source path")
}

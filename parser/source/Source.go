package source

import (
	"errors"
	"goselect/parser/tokenizer"
	"os/user"
	"strings"
)

type Source struct {
	tokenIterator *tokenizer.TokenIterator
	directory     string
}

func newSource(tokenIterator *tokenizer.TokenIterator) (*Source, error) {
	source := &Source{tokenIterator: tokenIterator}
	if err := source.setDirectory(); err != nil {
		return nil, err
	}
	return source, nil
}

func (source *Source) setDirectory() error {
	if source.tokenIterator.HasNext() && !source.tokenIterator.Peek().Equals("where") {
		token := source.tokenIterator.Next()
		path := token.TokenValue
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

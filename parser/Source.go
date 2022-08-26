package parser

import (
	"os/user"
	"strings"
)

type Source struct {
	directory string
}

func New(path string) *Source {
	if strings.HasPrefix(path, "~") {
		if currentUser, err := user.Current(); err != nil {
			return nil
		} else {
			return &Source{directory: currentUser.HomeDir + path[1:]}
		}
	}
	return &Source{directory: path}
}

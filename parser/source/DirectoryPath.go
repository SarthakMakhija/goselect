package source

import (
	"os/user"
	"strings"
)

func ExpandDirectoryPath(path string) (string, error) {
	if strings.HasPrefix(path, "~") {
		if currentUser, err := currentUser(); err != nil {
			return "", err
		} else {
			directory := currentUser.HomeDir + path[1:]
			return directory, nil
		}
	}
	return path, nil
}

var currentUserFunc = func() (*user.User, error) {
	return user.Current()
}

func currentUser() (*user.User, error) {
	return currentUserFunc()
}

func resetCurrentUser() {
	currentUserFunc = func() (*user.User, error) {
		return user.Current()
	}
}

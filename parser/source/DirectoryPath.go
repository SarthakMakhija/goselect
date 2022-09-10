package source

import (
	"os/user"
	"strings"
)

func ExpandDirectoryPath(path string) (string, error) {
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

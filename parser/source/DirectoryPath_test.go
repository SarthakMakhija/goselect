package source

import (
	"errors"
	"os/user"
	"testing"
)

func TestExpandDirectoryPathWithTilda(t *testing.T) {
	path, _ := ExpandDirectoryPath("~")
	currentUser, _ := user.Current()

	if path != currentUser.HomeDir {
		t.Fatalf("Expected directory path to be %v, received %v", currentUser.HomeDir, path)
	}
}

func TestExpandDirectoryPathWithAnError(t *testing.T) {
	currentUserFunc = func() (*user.User, error) {
		return nil, errors.New("error in test")
	}
	defer resetCurrentUser()

	_, err := ExpandDirectoryPath("~")

	if err == nil {
		t.Fatalf("Expected an error while expanding the directory path but received none")
	}
}

func TestExpandDirectoryPathWithoutTilda(t *testing.T) {
	path, _ := ExpandDirectoryPath("../source")

	if path != "../source" {
		t.Fatalf("Expected directory path to be %v, received %v", "../source", path)
	}
}

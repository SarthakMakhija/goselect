package parser

import (
	"os/user"
	"testing"
)

func TestCreatesANewSourceFromCurrentDirectory(t *testing.T) {
	path := "."
	source := New(path)
	if source.directory != path {
		t.Fatalf("Expected directory path to be %v, received %v", path, source.directory)
	}
}

func TestCreatesANewSourceWithHomeDirectorySymbol1(t *testing.T) {
	path := "~"
	source := New(path)
	expectedPath := homeDirectory()

	if source.directory != expectedPath {
		t.Fatalf("Expected directory path to be %v, received %v", expectedPath, source.directory)
	}
}

func TestCreatesANewSourceWithHomeDirectorySymbol2(t *testing.T) {
	path := "~/apps"
	source := New(path)
	expectedPath := homeDirectory() + "/apps"

	if source.directory != expectedPath {
		t.Fatalf("Expected directory path to be %v, received %v", expectedPath, source.directory)
	}
}

func homeDirectory() string {
	currentUser, err := user.Current()
	if err == nil {
		return currentUser.HomeDir
	}
	return ""
}

package parser

import (
	"os/user"
	"testing"
)

func TestCreatesANewSourceFromCurrentDirectory1(t *testing.T) {
	tokens := newEmptyTokens()
	tokens.add(newToken(RawString, "."))

	source, _ := newSource(tokens.iterator())
	if source.directory != "." {
		t.Fatalf("Expected directory path to be %v, received %v", ".", source.directory)
	}
}

func TestCreatesANewSourceFromCurrentDirectory2(t *testing.T) {
	tokens := newEmptyTokens()
	tokens.add(newToken(RawString, "."))
	tokens.add(newToken(RawString, "where"))

	source, _ := newSource(tokens.iterator())
	if source.directory != "." {
		t.Fatalf("Expected directory path to be %v, received %v", ".", source.directory)
	}
}

func TestCreatesANewSourceWithHomeDirectorySymbol1(t *testing.T) {
	tokens := newEmptyTokens()
	tokens.add(newToken(RawString, "~"))

	source, _ := newSource(tokens.iterator())
	expectedPath := homeDirectory()

	if source.directory != expectedPath {
		t.Fatalf("Expected directory path to be %v, received %v", expectedPath, source.directory)
	}
}

func TestCreatesANewSourceWithHomeDirectorySymbol2(t *testing.T) {
	tokens := newEmptyTokens()
	tokens.add(newToken(RawString, "~/apps"))

	source, _ := newSource(tokens.iterator())
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

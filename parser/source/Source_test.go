package source

import (
	"goselect/parser/tokenizer"
	"os/user"
	"testing"
)

func TestCreatesANewSourceFromCurrentDirectory1(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "."))

	source, _ := NewSource(tokens.Iterator())
	if source.directory != "." {
		t.Fatalf("Expected directory path to be %v, received %v", ".", source.directory)
	}
}

func TestCreatesANewSourceFromCurrentDirectory2(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "."))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "where"))

	source, _ := NewSource(tokens.Iterator())
	if source.directory != "." {
		t.Fatalf("Expected directory path to be %v, received %v", ".", source.directory)
	}
}

func TestCreatesANewSourceWithHomeDirectorySymbol1(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "~"))

	source, _ := NewSource(tokens.Iterator())
	expectedPath := homeDirectory()

	if source.directory != expectedPath {
		t.Fatalf("Expected directory path to be %v, received %v", expectedPath, source.directory)
	}
}

func TestCreatesANewSourceWithHomeDirectorySymbol2(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "~/apps"))

	source, _ := NewSource(tokens.Iterator())
	expectedPath := homeDirectory() + "/apps"

	if source.directory != expectedPath {
		t.Fatalf("Expected directory path to be %v, received %v", expectedPath, source.directory)
	}
}

func TestThrowsAnErrorWithoutAnyTokens(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	_, err := NewSource(tokens.Iterator())

	if err == nil {
		t.Fatalf("Expected errors to be non-nil when creating a source without any tokens")
	}
}

func homeDirectory() string {
	currentUser, err := user.Current()
	if err == nil {
		return currentUser.HomeDir
	}
	return ""
}

//go:build unit
// +build unit

package alias

import (
	"os"
	"reflect"
	"testing"
)

func TestAddsAQueryAlias(t *testing.T) {
	queryAliasReference := NewQueryAlias()
	defer os.Remove(queryAliasReference.FilePath)

	_ = queryAliasReference.Add(Alias{Query: "select * from .", Alias: "lsCurrent"})
	query, _, _ := queryAliasReference.GetQueryBy("lsCurrent")
	expected := "select * from ."

	if query != expected {
		t.Fatalf("Expected query to be %v, received %v", expected, query)
	}
}

func TestGetsANonExistingAlias(t *testing.T) {
	queryAliasReference := NewQueryAlias()
	defer os.Remove(queryAliasReference.FilePath)

	_ = queryAliasReference.Add(Alias{Query: "select * from .", Alias: "lsCurrent"})
	query, _, _ := queryAliasReference.GetQueryBy("unknown")
	expected := ""

	if query != expected {
		t.Fatalf("Expected query to be %v, received %v", expected, query)
	}
}

func TestGetsANonExistingAliasInNonExistingAliasFile(t *testing.T) {
	queryAliasReference := NewQueryAlias()
	defer os.Remove(queryAliasReference.FilePath)

	query, _, _ := queryAliasReference.GetQueryBy("unknown")
	expected := ""

	if query != expected {
		t.Fatalf("Expected query to be %v, received %v", expected, query)
	}
}

func TestGetsANonExistingAliasInNonExistingAliasFileEnsuringThereIsNoError(t *testing.T) {
	queryAliasReference := NewQueryAlias()
	defer os.Remove(queryAliasReference.FilePath)

	_, _, err := queryAliasReference.GetQueryBy("unknown")
	if err != nil {
		t.Fatalf("Expected no error while getting the query by an alias in a non-existing file")
	}
}

func TestAddsAQueryAliasToAnExistingFile(t *testing.T) {
	queryAliasReference := NewQueryAlias()
	defer os.Remove(queryAliasReference.FilePath)

	_ = queryAliasReference.Add(Alias{Query: "select * from .", Alias: "lsCurrent"})
	_ = queryAliasReference.Add(Alias{Query: "select * from ~/Downloads", Alias: "lsDownloads"})

	query, _, _ := queryAliasReference.GetQueryBy("lsDownloads")
	expected := "select * from ~/Downloads"

	if query != expected {
		t.Fatalf("Expected query to be %v, received %v", expected, query)
	}
}

func TestAttemptsToAddTheSameAlias(t *testing.T) {
	queryAliasReference := NewQueryAlias()
	defer os.Remove(queryAliasReference.FilePath)

	_ = queryAliasReference.Add(Alias{Query: "select * from .", Alias: "lsCurrent"})
	err := queryAliasReference.Add(Alias{Query: "select * from ~/Downloads", Alias: "lsCurrent"})

	if err == nil {
		t.Fatalf("Expected an error while adding the same alias but received none")
	}
}

func TestAttemptsToAddTheSamePreExistingAlias(t *testing.T) {
	queryAliasReference := NewQueryAlias()
	defer os.Remove(queryAliasReference.FilePath)

	err := queryAliasReference.Add(Alias{Query: "select * from ~/Downloads", Alias: "lsCurrent"})

	if err == nil {
		t.Fatalf("Expected an error while adding the same alias but received none")
	}
}

func TestAttemptsToReadACorruptFile(t *testing.T) {
	queryAliasReference := NewQueryAlias()
	defer os.Remove(queryAliasReference.FilePath)

	_ = os.WriteFile(queryAliasReference.FilePath, []byte("Hello"), 0644)
	err := queryAliasReference.Add(Alias{Query: "select * from .", Alias: "lsCurrent"})
	if err == nil {
		t.Fatalf("Expected an error while dealing with a corrupt alias file, received none")
	}
}

func TestGetsAllTheAliases(t *testing.T) {
	queryAliasReference := NewQueryAlias()
	defer os.Remove(queryAliasReference.FilePath)

	_ = queryAliasReference.Add(Alias{Query: "select * from .", Alias: "lsCurrent"})
	_ = queryAliasReference.Add(Alias{Query: "select * from ~/Downloads", Alias: "lsDownloads"})

	added := map[string]string{
		"lsCurrent":   "select * from .",
		"lsDownloads": "select * from ~/Downloads",
	}
	expected := queryAliasReference.merge(added, preDefinedAliases)
	aliases, _ := queryAliasReference.All()
	if !reflect.DeepEqual(expected, aliases) {
		t.Fatalf("Expected all the aliases to be %v, received %v", expected, aliases)
	}
}

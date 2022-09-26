//go:build integration
// +build integration

package test

import (
	"bytes"
	"goselect/cmd"
	"goselect/parser/context"
	"strings"
	"testing"
)

func TestTimeFormats(t *testing.T) {
	cmd.GetRootCommand().SetArgs([]string{"listTimeFormats"})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()
	contents := buffer.String()

	formats := context.SupportedFormats()
	for formatId, definition := range formats {
		if !strings.Contains(contents, formatId) {
			t.Fatalf("Expected format id %v to be contained in the supported formats but was not, received %v", formatId, contents)
		}
		if !strings.Contains(contents, definition.Format) {
			t.Fatalf("Expected format %v to be contained in the supported formats but was not, received %v", definition.Format, contents)
		}
	}
}

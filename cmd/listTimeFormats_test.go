package cmd

import (
	"bytes"
	"goselect/parser/context"
	"strings"
	"testing"
)

func TestTimeFormats(t *testing.T) {
	rootCmd.SetArgs([]string{"listTimeFormats"})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()
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
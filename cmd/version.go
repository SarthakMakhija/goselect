package cmd

import (
	_ "embed"
	"encoding/json"
	"github.com/spf13/cobra"
)

//go:embed version.info
var contents []byte

type Versions struct {
	Versions []Version `json:"versions"`
}

type Version struct {
	Version   string `json:"version"`
	IsCurrent bool   `json:"isCurrent"`
	Changes   string `json:"changes"`
}

func newVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Return the current version of goselect",
		Example: `
1. goselect version
2. goselect v
`,
		Run: func(cmd *cobra.Command, args []string) {
			versions := parseVersion()
			colorVersion := "\033[36m"
			colorChanges := "\033[33m"
			colorReset := "\033[0m"

			for _, version := range versions.Versions {
				if version.IsCurrent {
					cmd.Print(colorVersion)
					cmd.Println(version.Version)
					cmd.Println(colorChanges)
					cmd.Println("Changes")
					cmd.Println(version.Changes)
					cmd.Println(colorReset)
					break
				}
			}
		},
	}
}

func parseVersion() Versions {
	var versions Versions
	_ = json.Unmarshal(contents, &versions)
	return versions
}

func init() {
	rootCmd.AddCommand(newVersionCommand())
}

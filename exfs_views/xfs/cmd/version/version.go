package version

import (
	"github.com/spf13/cobra"
)

// Cmd is the version subcommand.
var Cmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  `Print the version, commit hash, and build time of the exfs CLI.`,
	Run:   run,
}

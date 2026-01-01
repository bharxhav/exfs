package init

import (
	"github.com/spf13/cobra"
)

// Cmd is the init subcommand.
var Cmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize EXFS repository",
	Long:  `Initialize a new EXFS repository in the current directory.`,
	RunE:  run,
}

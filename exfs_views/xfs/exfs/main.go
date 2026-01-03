/*
Root command for exfs CLI.
This is the entry point - runs when user types "exfs" with no subcommand.
Registers all top-level subcommands and global persistent flags.
*/

package exfs

import (
	"github.com/spf13/cobra"

	initcmd "xfs.exfs.org/exfs/init"
)

var commandName = "exfs"

var commandShortDescription = "Extensible File System CLI"

var commandLongDescription = `A multifaceted framework to manage non-functional characteristics of code that provides programmatic self reference, single source of truth, and separation of concerns.`

var Cmd = &cobra.Command{
	Use:                commandName,
	Aliases:            []string{"xfs"},
	Short:              commandShortDescription,
	Long:               commandLongDescription,
	Version:            "0.0.1",
	Args:               cobra.NoArgs,
	SilenceUsage:       true,
	PersistentPreRunE:  persistentPreRun,
	PreRunE:            preRun,
	RunE:               run,
	PostRunE:           postRun,
	PersistentPostRunE: persistentPostRun,
}

func init() {
	Cmd.AddCommand(initcmd.Cmd)
}

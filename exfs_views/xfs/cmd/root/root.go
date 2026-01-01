package root

import (
	"github.com/spf13/cobra"

	"github.com/bharxhav/exfs/cmd/version"
	"github.com/bharxhav/exfs/internal/buildinfo"
)

// Cmd is the root command for exfs.
// Both "exfs" and "xfs" are valid invocations (via symlink: ln -s exfs xfs).
var Cmd = &cobra.Command{
	Use:     "exfs",
	Aliases: []string{"xfs"},
	Short:   "EXFS - Extensible File System",
	Long:    `A multifaceted framework to manage non-functional characteristics of code that provides programmatic self reference, single source of truth, and separation of concerns.`,
	Version: buildinfo.Get().Version,
}

func init() {
	Cmd.AddCommand(version.Cmd)
	registerGlobalFlags()
}

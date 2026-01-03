/*
Initialize EXFS repository.
Creates .exfs/ directory structure and CSV database files.
*/

package init

import (
	"github.com/spf13/cobra"
)

var commandName = "init"

var commandShortDescription = "Initialize an EXFS repository"

var commandLongDescription = `Initialize a new EXFS repository in the current directory.
Creates .exfs/, exfs/, _exfs/, exfs_views/, exfs_removed/ directories and CSV database files.`

var Cmd = &cobra.Command{
	Use:   commandName,
	Short: commandShortDescription,
	Long:  commandLongDescription,
	Args:  cobra.NoArgs,

	PersistentPreRunE:  persistentPreRun,
	PreRunE:            preRun,
	RunE:               run,
	PostRunE:           postRun,
	PersistentPostRunE: persistentPostRun,
}

func init() {
	Cmd.Flags().BoolP("force", "f", false, "Reinitialize even if already exists")
}

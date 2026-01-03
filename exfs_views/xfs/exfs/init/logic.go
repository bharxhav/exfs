/*
Init command hooks.
Creates directory structure and CSV files.
*/

package init

import (
	"github.com/spf13/cobra"
)

func persistentPreRun(cmd *cobra.Command, args []string) error {
	return nil
}

func preRun(cmd *cobra.Command, args []string) error {
	return nil
}

func run(cmd *cobra.Command, args []string) error {
	return nil
}

func postRun(cmd *cobra.Command, args []string) error {
	return nil
}

func persistentPostRun(cmd *cobra.Command, args []string) error {
	return nil
}

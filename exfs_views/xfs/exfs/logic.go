/*
Root command hooks.
Defines PersistentPreRunE, PreRunE, RunE, PostRunE, PersistentPostRunE for root.
*/

package exfs

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

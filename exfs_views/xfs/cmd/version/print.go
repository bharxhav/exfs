package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/bharxhav/exfs/internal/buildinfo"
)

func run(cmd *cobra.Command, args []string) {
	info := buildinfo.Get()
	printVersion(info)
}

func printVersion(info buildinfo.Info) {
	fmt.Printf("exfs version %s\n", info.Version)
	if info.Commit != "" {
		fmt.Printf("  commit: %s\n", info.Commit)
	}
	if info.Time != "" {
		fmt.Printf("  built:  %s\n", info.Time)
	}
	if info.Modified {
		fmt.Println("  dirty:  yes (uncommitted changes)")
	}
}

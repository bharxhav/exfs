package main

import (
	"os"

	"xfs.exfs.org/exfs"
)

func main() {
	if err := exfs.Cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

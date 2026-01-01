package root

import (
	"fmt"
	"os"
)

// Execute runs the root command and exits on error.
func Execute() {
	if err := Cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

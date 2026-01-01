package init

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bharxhav/exfs/internal/config"
	"github.com/bharxhav/exfs/internal/fs"
	"github.com/spf13/cobra"
)

func run(cmd *cobra.Command, args []string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	for _, dir := range config.Directories {
		path := filepath.Join(cwd, dir)
		if err := fs.CreateDirectory(path); err != nil {
			return err
		}
		fmt.Printf("+ %s/\n", dir)
	}

	for _, file := range config.CSVFiles {
		path := filepath.Join(cwd, file.Path)
		if err := fs.CreateFile(path, file.Schema); err != nil {
			return err
		}
		fmt.Printf("+ %s\n", file.Path)
	}

	fmt.Println("\nEXFS repository initialized!")
	return nil
}

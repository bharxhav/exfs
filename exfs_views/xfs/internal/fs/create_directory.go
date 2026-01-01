package fs

import (
	"fmt"
	"os"
)

// CreateDirectory creates a directory at the given path (relative to cwd).
// Uses MkdirAll so parent directories are created as needed.
func CreateDirectory(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", path, err)
	}
	return nil
}

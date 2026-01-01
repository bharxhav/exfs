package fs

import (
	"fmt"
	"os"
)

// CreateFile creates a file at the given path with the given content.
func CreateFile(path string, content string) error {
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to create file %s: %w", path, err)
	}
	return nil
}

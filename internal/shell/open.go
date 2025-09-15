package shell

import (
	"errors"
	"os/exec"
	"path/filepath"
)

// OpenWindowsTerminal opens Windows Terminal (wt.exe) at the specified directory.
func OpenWindowsTerminal(dir string) error {
	if dir == "" {
		return errors.New("directory is empty")
	}
	// Use wt.exe -d <dir> to set starting directory. Use start to avoid blocking? Here we exec directly.
	cmd := exec.Command("wt.exe", "-d", filepath.Clean(dir))
	return cmd.Start()
}

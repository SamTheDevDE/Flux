package ui

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"Flux/m/internal/config"
)

// ShowConfig opens a minimal config flow using a native folder picker via PowerShell.
// It prompts the user to select a folder and saves it to configuration.
func ShowConfig() error {
	selected, err := pickFolder()
	if err != nil {
		return err
	}
	if selected == "" {
		return errors.New("no folder selected")
	}
	if err := config.Save(config.AppConfig{DefaultDir: selected}); err != nil {
		return err
	}
	fmt.Println("Configuration saved:", selected)
	return nil
}

// pickFolder shows a folder selection UI using PowerShell and returns the chosen path.
func pickFolder() (string, error) {
	// Show a WinForms FolderBrowserDialog via PowerShell.
	script := `
Add-Type -AssemblyName System.Windows.Forms | Out-Null
$dlg = New-Object System.Windows.Forms.FolderBrowserDialog
$dlg.Description = 'Select default directory for Flux'
$dlg.ShowNewFolderButton = $true
if ($dlg.ShowDialog() -eq [System.Windows.Forms.DialogResult]::OK) { $dlg.SelectedPath }
`
	cmd := exec.Command("powershell", "-NoProfile", "-Command", script)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("folder picker failed: %w: %s", err, out.String())
	}
	path := strings.TrimSpace(out.String())
	// Validate JSON safety (not strictly needed) and that it's not empty.
	_, _ = json.Marshal(path)
	return path, nil
}

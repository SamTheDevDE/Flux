package main

import (
	"flag"
	"fmt"
	"os"

	"Flux/m/internal/config"
	"Flux/m/internal/shell"
	ui "Flux/m/internal/ui"
)

func main() {
	var showConfig bool
	flag.BoolVar(&showConfig, "config", false, "open configuration UI")
	flag.Parse()

	if showConfig {
		// Open the configuration GUI.
		if err := ui.ShowConfig(); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		return
	}

	// Normal launch: attempt to load config and open terminal.
	cfg, err := config.Load()
	if err != nil {
		// If we cannot load, fall back to config UI.
		_ = ui.ShowConfig()
		return
	}
	dir := cfg.DefaultDir
	if dir == "" {
		// If not configured, open UI to set it.
		_ = ui.ShowConfig()
		return
	}
	if err := shell.OpenWindowsTerminal(dir); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to open Windows Terminal:", err)
		os.Exit(1)
	}
}

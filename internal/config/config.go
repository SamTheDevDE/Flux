package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// AppConfig represents the persisted application configuration.
type AppConfig struct {
	DefaultDir string `json:"defaultDir"`
}

// configDir returns the directory under %AppData% where we store config.
func configDir() (string, error) {
	appData := os.Getenv("APPDATA")
	if appData == "" {
		return "", errors.New("APPDATA not set")
	}
	dir := filepath.Join(appData, "Flux")
	return dir, nil
}

// Path returns the full path to the config file.
func Path() (string, error) {
	dir, err := configDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "config.json"), nil
}

// Load reads the configuration from disk. If the file doesn't exist, returns a zero-value config and no error.
func Load() (AppConfig, error) {
	var cfg AppConfig
	p, err := Path()
	if err != nil {
		return cfg, err
	}
	b, err := os.ReadFile(p)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return cfg, nil
		}
		return cfg, err
	}
	if err := json.Unmarshal(b, &cfg); err != nil {
		return cfg, fmt.Errorf("invalid config JSON: %w", err)
	}
	return cfg, nil
}

// Save writes the configuration to disk, creating the directory if needed.
func Save(cfg AppConfig) error {
	if cfg.DefaultDir == "" {
		return errors.New("default directory cannot be empty")
	}
	info, err := os.Stat(cfg.DefaultDir)
	if err != nil || !info.IsDir() {
		return fmt.Errorf("default directory is not a valid directory: %s", cfg.DefaultDir)
	}
	p, err := Path()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(p, b, 0o644)
}

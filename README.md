# Flux

A tiny Windows helper that opens Windows Terminal in your preferred directory.

- Double-click the app to open Windows Terminal at your configured default directory.
- Run with `--config` to open a minimal folder picker and set/update the default directory.

## Requirements

- Windows 10/11 with Windows Terminal installed (provides `wt.exe` in PATH).
- Go 1.21+ to build from source.

## How it works

- Configuration is saved to `%AppData%/Flux/config.json` with a single field `defaultDir`.
- `--config` launches a native folder picker (via PowerShell WinForms) and writes the config.
- Normal launch reads config and runs `wt.exe -d <defaultDir>`.

## Build

From the repository root:

```powershell
# Standard console build (shows console window)
go build -ldflags "-H=windowsgui" -o .\flux.exe .\cmd\flux

# Windows GUI build (no console window on double-click)
go build -o .\flux.exe .\cmd\flux
```

If you see an error that `wt.exe` is not found, ensure Windows Terminal is installed and `wt.exe` is available on PATH.

## Usage

```powershell
# First-time: set your default folder
.\flux.exe --config

# Thereafter: double-click or run to open Windows Terminal in that folder
.\flux.exe
```

## Notes

- The folder picker uses PowerShell and WinForms. If PowerShell is restricted by policy, you may need to allow running the embedded command or set the config JSON manually at `%AppData%\Flux\config.json`.
- To change the default directory later, rerun with `--config`.
- Currently targets Windows Terminal (`wt.exe`). You could adapt `internal/shell/open.go` to support other terminals.

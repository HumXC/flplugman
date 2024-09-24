package main

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/HumXC/flplugman/nfo"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
func (a *App) GetPluginDBPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return home + "\\Documents\\Image-Line\\FL Studio\\Presets\\Plugin database\\Installed", nil
}
func (a *App) ScanPluginDB(dir string) ([]nfo.Plugin, error) {
	result := make([]nfo.Plugin, 0)
	e := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !(!info.IsDir() && strings.HasSuffix(info.Name(), ".nfo")) {
			return nil
		}
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		p, err := nfo.Unmarshal(b)
		if err != nil {
			return err
		}
		result = append(result, p)
		return nil
	})
	return result, e
}

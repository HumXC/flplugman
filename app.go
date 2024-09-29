package main

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/HumXC/flplugman/config"
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

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetConfig() (config.Config, error) {
	return config.Get()
}

func (a *App) SaveConfig(c config.Config) error {
	return config.Save(c)
}

func (a *App) RemoveConfig() error {
	return config.Remove()
}

func (a *App) GetPluginDBPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return home + "\\Documents\\Image-Line\\FL Studio\\Presets\\Plugin database", nil
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

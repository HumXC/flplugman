package main

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/HumXC/flplugman/config"
	"github.com/HumXC/flplugman/nfo"
)

type Plugin struct {
	nfo.Plugin
	PresetPath string // 相对于 PluginDBPath 的路径
	Name       string // .fst 和 .nfo 的无后缀文件名
	Fst        string // .fst 的完整路径
	Nfo        string // .nfo 的完整路径
}

// App struct
type App struct {
	ctx    context.Context
	config *config.Config
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetConfig() (config.Config, error) {
	if a.config == nil {
		c, err := config.Get()
		if err != nil {
			return config.Config{}, err
		}
		a.config = &c
	}
	return *a.config, nil
}

func (a *App) SaveConfig(c config.Config) error {
	a.config = &c
	return config.Save(c)
}

func (a *App) GetConfigPath() string {
	return config.ConfigPath
}

func (a *App) RemoveConfig() error {
	return config.Remove()
}

func (a *App) GetPluginDBPath() (string, error) {
	c, err := a.GetConfig()
	if err != nil {
		return "", err
	}
	return filepath.Join(c.FLDataDir, "FL Studio\\Presets\\Plugin database"), nil
}

// 移动 .nfo 和 .fst 文件到 path 目录，返回修改后的 Plugin
func (a *App) MovePlugin(plug *Plugin, path string) (Plugin, error) {
	if plug == nil {
		return Plugin{}, fmt.Errorf("plugin is nil")
	}
	p := *plug
	if filepath.IsAbs(path) {
		return p, fmt.Errorf("path \"%s\" must be relative path", path)
	}
	flPluginDB, err := a.GetPluginDBPath()
	if err != nil {
		return p, err
	}

	dist := filepath.Join(flPluginDB, path, p.Name)

	distFst := dist + ".fst"
	distNfo := dist + ".nfo"

	if _, err := os.Stat(filepath.Dir(dist)); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(dist), 0644); err != nil {
			return p, err
		}
	}

	if err := os.Rename(p.Fst, distFst); err != nil {
		return p, err
	}
	p.Fst = distFst
	p.PS.PresetFilename = p.Fst
	if err := os.WriteFile(distNfo, nfo.Marshal(p.Plugin), 0644); err != nil {
		return p, err
	}
	if err := os.Remove(p.Nfo); err != nil {
		return p, err
	}
	p.Nfo = distNfo

	srcDir := filepath.Join(flPluginDB, p.PresetPath)
	p.PresetPath = path
	// TODO: 让文件夹也是一种可管理的资源, 而不是直接删除
	if err := deleteEmptyDir(srcDir); err != nil {
		return p, err
	}

	return p, nil
}

func deleteEmptyDir(dir string) error {
	fmt.Println("删除", dir)
	es, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	if len(es) == 0 {
		if err := os.Remove(dir); err != nil {
			return deleteEmptyDir(filepath.Dir(dir))
		}
	}
	return nil
}
func (a *App) ScanPluginDB() ([]Plugin, error) {
	PluginDBPath, err := a.GetPluginDBPath()
	if err != nil {
		return nil, err
	}
	result := make([]Plugin, 0)
	e := filepath.Walk(PluginDBPath, func(path string, info fs.FileInfo, err error) error {
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
		if p.PS.Name == "" {
			return nil
		}
		if strings.Contains(p.PS.PresetFilename, "%FLPluginDBPath%\\") {
			p.PS.PresetFilename = filepath.Join(PluginDBPath, strings.ReplaceAll(p.PS.PresetFilename, "%FLPluginDBPath%\\", ""))
		}
		rel, err := filepath.Rel(PluginDBPath, p.PS.PresetFilename)
		if err != nil {
			return err
		}
		base := filepath.Base(p.PS.PresetFilename)
		pp := Plugin{
			Plugin:     p,
			PresetPath: filepath.Dir(rel),
			Name:       base[:len(base)-4],
			Fst:        p.PS.PresetFilename,
			Nfo:        p.PS.PresetFilename[:len(p.PS.PresetFilename)-4] + ".nfo",
		}
		result = append(result, pp)
		return nil
	})
	return result, e
}

package main

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/HumXC/flplugman/config"
	"github.com/HumXC/flplugman/nfo"
	"github.com/cascax/colorthief-go"
	"github.com/reujab/wallpaper"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Plugin struct {
	Nfo        nfo.Plugin
	PresetPath string // 相对于 PluginDBPath 的路径
	Name       string // .fst 和 .nfo 的无后缀文件名
	FstName    string // .fst 的完整路径
	NfoName    string // .nfo 的完整路径
	Vendorname string
}

// App struct
type App struct {
	ctx           context.Context
	config        *config.Config
	isSavedConfig bool
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	go func(ctx context.Context) {
		modTime := time.Now() // 通过修改时间来判断更改
		for {
			w, err := wallpaper.Get()
			if err != nil {
				runtime.LogError(ctx, err.Error())
				return
			}
			stat, err := os.Stat(w)
			if err != nil {
				runtime.LogError(ctx, err.Error())
				return
			}

			if stat.ModTime() != modTime {
				modTime = stat.ModTime()
				cs, err := colorthief.GetPaletteFromFile(w, 6)
				if err != nil {
					runtime.LogError(ctx, err.Error())
					return
				}
				result := make([]string, 0, len(cs))
				for _, c := range cs {
					r, g, b, _ := c.RGBA()
					result = append(result, fmt.Sprintf("#%02x%02x%02x", r>>8, g>>8, b>>8))
				}
				runtime.EventsEmit(ctx, "wallpaper-color-changed", result)
			}
			time.Sleep(2 * time.Second)
		}
	}(ctx)
}

func (a *App) GetWallpaperColor() ([]string, error) {
	w, err := wallpaper.Get()
	if err != nil {
		return nil, err
	}
	cs, err := colorthief.GetPaletteFromFile(w, 6)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0, len(cs))
	for _, c := range cs {
		r, g, b, _ := c.RGBA()
		result = append(result, fmt.Sprintf("#%02x%02x%02x", r>>8, g>>8, b>>8))
	}
	return result, nil
}

func (a *App) GetConfig() (config.Config, error) {
	if a.config == nil || a.isSavedConfig {
		c, err := config.Get()
		if err != nil {
			return config.Config{}, err
		}
		a.config = &c
	}
	return *a.config, nil
}

func (a *App) SaveConfig(c config.Config) error {
	if err := config.Save(c); err != nil {
		return err
	}
	a.config = &c
	a.isSavedConfig = true
	return nil
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
	runtime.LogInfof(a.ctx, "MovePlugin %s: %s to %s", plug.Name, plug.PresetPath, path)
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

	if err := os.Rename(p.FstName, distFst); err != nil {
		return p, err
	}
	p.FstName = distFst
	p.Nfo.PS.PresetFilename = p.FstName
	if err := os.WriteFile(distNfo, nfo.Marshal(p.Nfo), 0644); err != nil {
		return p, err
	}
	if err := os.Remove(p.NfoName); err != nil {
		return p, err
	}
	p.NfoName = distNfo

	srcDir := filepath.Join(flPluginDB, p.PresetPath)
	p.PresetPath = path
	// TODO: 让文件夹也是一种可管理的资源, 而不是直接删除
	if err := deleteEmptyDir(srcDir); err != nil {
		return p, err
	}

	return p, nil
}

func deleteEmptyDir(dir string) error {
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
		vendornames := []string{}
		for _, v := range p.PS.File {
			vendornames = append(vendornames, v.Vendorname)
		}
		pp := Plugin{
			Nfo:        p,
			PresetPath: filepath.Dir(rel),
			Name:       base[:len(base)-4],
			FstName:    p.PS.PresetFilename,
			NfoName:    p.PS.PresetFilename[:len(p.PS.PresetFilename)-4] + ".nfo",
			Vendorname: strings.Join(vendornames, ","),
		}
		result = append(result, pp)
		return nil
	})
	return result, e
}

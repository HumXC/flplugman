package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/HumXC/flplugman/config"
	"github.com/HumXC/flplugman/log"
	"github.com/HumXC/flplugman/nfo"
	"github.com/cascax/colorthief-go"
	"github.com/reujab/wallpaper"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gopkg.in/yaml.v3"
)

type Plugin struct {
	Nfo           nfo.Plugin
	PresetPath    string // 相对于 PluginDBPath 的路径
	Name          string // .fst 和 .nfo 的无后缀文件名
	FstName       string // .fst 的完整路径
	NfoName       string // .nfo 的完整路径
	Vendorname    string
	Cover         string // 封面的 Base64 编码
	CoverMimeType string // 封面的类型
}

// App struct
type App struct {
	ctx    context.Context
	config *config.Config
	logger log.Logger
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.logger.Debug("Starting up...")
	a.GetConfig()
	c, _ := yaml.Marshal(a.config)
	a.logger.Info("Condif: \n", string(c))
	go func(ctx context.Context) {
		modTime := time.Now() // 通过修改时间来判断更改
		a.logger.Debug("Start monitoring wallpaper event: \"wallpaper-color-changed\"...")
		for {
			w, err := wallpaper.Get()
			if err != nil {
				a.logger.Error(err)
				return
			}
			stat, err := os.Stat(w)
			if err != nil {
				a.logger.Error(err)
				return
			}

			if stat.ModTime() != modTime {
				modTime = stat.ModTime()
				cs, err := colorthief.GetPaletteFromFile(w, 6)
				if err != nil {
					a.logger.Error(err)
					return
				}
				result := make([]string, 0, len(cs))
				for _, c := range cs {
					r, g, b, _ := c.RGBA()
					result = append(result, fmt.Sprintf("#%02x%02x%02x", r>>8, g>>8, b>>8))
				}
				a.logger.Debug("event: \"wallpaper-color-changed\" emit")
				runtime.EventsEmit(ctx, "wallpaper-color-changed", result)
			}
			time.Sleep(2 * time.Second)
		}
	}(ctx)
}

func (a *App) GetWallpaperColor() ([]string, error) {
	w, err := wallpaper.Get()
	if err != nil {
		a.logger.Error(err)
		return nil, err
	}
	cs, err := colorthief.GetPaletteFromFile(w, 6)
	if err != nil {
		a.logger.Error(err)
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
	a.logger.Debugw("Get config")
	if a.config == nil {
		a.logger.Debugw("Load config file", "configFile", config.ConfigPath)
		c, isInit, err := config.Get()
		a.logger.Debugw("Config loaded", "isInit", isInit)
		if err != nil {
			a.logger.Error(err)
			return config.Config{}, err
		}
		a.config = &c
	}
	return *a.config, nil
}

func (a *App) SaveConfig(c config.Config) error {
	a.logger.Debugw("Saving config", "configFile", config.ConfigPath)
	if err := config.Save(c); err != nil {
		a.logger.Error(err)
		return err
	}
	a.config = &c
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
		a.logger.Error(err)
		return "", err
	}
	return filepath.Join(c.FLDataDir, "FL Studio\\Presets\\Plugin database"), nil
}

// 移动 .nfo 和 .fst 文件到 path 目录，返回修改后的 Plugin
func (a *App) MovePlugin(plug *Plugin, path string) (Plugin, error) {
	if plug == nil {
		err := fmt.Errorf("plugin is nil")
		a.logger.Error(err)
		return Plugin{}, err
	}
	runtime.LogInfof(a.ctx, "Moving plugin [%s] from %s to %s", plug.Name, plug.PresetPath, path)
	p := *plug
	if filepath.IsAbs(path) {
		err := fmt.Errorf("path \"%s\" must be relative path", path)
		a.logger.Error(err)
		return p, err
	}
	flPluginDB, err := a.GetPluginDBPath()
	if err != nil {
		a.logger.Error(err)
		return p, err
	}
	path = filepath.Clean(path)
	if path == "" {
		err := fmt.Errorf("path is empty")
		a.logger.Error(err)
		return p, err
	}

	dist := filepath.Join(flPluginDB, path, p.Name)
	distFst := dist + ".fst"
	distNfo := dist + ".nfo"

	if _, err := os.Stat(filepath.Dir(dist)); os.IsNotExist(err) {
		a.logger.Debugw("Create directory", "path", filepath.Dir(dist))
		if err := os.MkdirAll(filepath.Dir(dist), 0644); err != nil {
			a.logger.Error(err)
			return p, err
		}
	}
	// 移动 .fst
	a.logger.Debug("Rename", "src", p.FstName, "dist", distFst)
	if err := os.Rename(p.FstName, distFst); err != nil {
		a.logger.Error(err)
		return p, err
	}
	p.FstName = distFst
	p.Nfo.PS.PresetFilename = p.FstName

	// 移动封面
	if p.Cover != "" {
		bitmap, err := base64.StdEncoding.DecodeString(p.Cover)
		if err != nil {
			a.logger.Error(err)
			return p, err
		}
		a.logger.Debug("Remove", "path", filepath.Join(filepath.Dir(p.NfoName), p.Nfo.Bitmap))
		if err := os.Remove(filepath.Join(filepath.Dir(p.NfoName), p.Nfo.Bitmap)); err != nil {
			a.logger.Error(err)
			return p, err
		}
		coverSuffix := "." + strings.Split(p.CoverMimeType, "/")[1]
		a.logger.Debug("Write", "path", dist+coverSuffix)
		if err := os.WriteFile(dist+coverSuffix, bitmap, 0644); err != nil {
			a.logger.Error(err)
			return p, err
		}
		p.Nfo.Bitmap = p.Name + coverSuffix
	}
	// 移动 .nfo
	a.logger.Debug("Write", "path", distNfo)
	if err := os.WriteFile(distNfo, nfo.Marshal(p.Nfo), 0644); err != nil {
		a.logger.Error(err)
		return p, err
	}
	a.logger.Debug("Remove", "path", p.NfoName)
	if err := os.Remove(p.NfoName); err != nil {
		a.logger.Error(err)
		return p, err
	}
	p.NfoName = distNfo

	srcDir := filepath.Join(flPluginDB, p.PresetPath)
	p.PresetPath = path
	// TODO: 让文件夹也是一种可管理的资源, 而不是直接删除
	if ds, err := deleteEmptyDir(srcDir); err != nil {
		a.logger.Error(err)
		return p, err
	} else {
		a.logger.Debugw("Delete empty directory", "dirs", ds)
	}

	return p, nil
}

func deleteEmptyDir(dir string) ([]string, error) {
	deleted := []string{}
	es, err := os.ReadDir(dir)
	if err != nil || len(es) != 0 {
		return deleted, err
	}

	if err := os.Remove(dir); err != nil {
		return deleted, err
	}
	deleted = append(deleted, dir)
	d, err := deleteEmptyDir(filepath.Dir(dir))
	return append(deleted, d...), err
}
func (a *App) ScanPluginDB() ([]Plugin, error) {
	PluginDBPath, err := a.GetPluginDBPath()
	a.logger.Infow("Scan plugin db", "path", PluginDBPath)
	if err != nil {
		a.logger.Error(err)
		return nil, err
	}
	result := make([]Plugin, 0)
	e := filepath.Walk(PluginDBPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			a.logger.Error(err)
			return err
		}
		if !(!info.IsDir() && strings.HasSuffix(info.Name(), ".nfo")) {
			a.logger.Error(err)
			return nil
		}
		b, err := os.ReadFile(path)
		if err != nil {
			a.logger.Error(err)
			return err
		}
		p, err := nfo.Unmarshal(b)
		if err != nil {
			a.logger.Error(err)
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
			a.logger.Error(err)
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
		if p.Bitmap != "" {
			bitmap := filepath.Join(filepath.Dir(path), p.Bitmap)
			f, err := os.ReadFile(bitmap)
			if err != nil {
				a.logger.Error(err)
			} else {
				pp.Cover = base64.StdEncoding.EncodeToString(f)
				pp.CoverMimeType = http.DetectContentType(f)
			}
		}
		result = append(result, pp)
		return nil
	})
	if e != nil {
		a.logger.Error(e)
	}
	return result, e
}

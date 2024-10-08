package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

var ConfigPath = "./"

func init() {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		// TODO: log error
		return
	}
	ConfigPath = filepath.Join(cfgDir, "flplugman", "config.json")
}

type Config struct {
	IsGreeted bool   `json:"is_greeted"`
	FLDataDir string `json:"fl_data_dir"`
}

func Get() (Config, error) {
	var c Config
	if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
		key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Image-Line\Shared\Paths`, registry.QUERY_VALUE)
		if err != nil {
			return c, err
		}
		val, _, err := key.GetStringValue("Shared data")
		if err != nil {
			return c, err
		}
		c.FLDataDir = val
		return c, nil
	}
	b, err := os.ReadFile(ConfigPath)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(b, &c)
	return c, err
}

func Remove() error {
	return os.Remove(ConfigPath)
}

func Save(c Config) error {
	if _, err := os.Stat(filepath.Dir(ConfigPath)); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(ConfigPath), 0644); err != nil {
			return err
		}
	}
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(ConfigPath, b, 0644)
}

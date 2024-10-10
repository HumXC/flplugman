package main

import (
	"embed"

	"github.com/HumXC/flplugman/log"
	"github.com/wailsapp/wails/v2"
	wailsLogger "github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS
var LogLevel = wailsLogger.DEBUG

func main() {
	app := NewApp()
	app.logger = log.NewGoLogger(LogLevel)
	err := wails.Run(&options.App{
		Title:  "flplugman",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},

		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			log.NewJSLogger(LogLevel),
		},
		Logger:   log.NewWailsLogger(LogLevel),
		LogLevel: LogLevel,
		Windows: &windows.Options{
			WebviewIsTransparent:              false,
			WindowIsTranslucent:               true,
			BackdropType:                      windows.Mica,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			Theme:                             windows.SystemDefault,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

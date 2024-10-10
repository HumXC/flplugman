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
	// Create an instance of the app structure
	app := NewApp()
	logger := log.NewLogger(LogLevel)
	// Create application with options
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
			&log.JSLogger{SugaredLogger: logger},
		},
		Logger:   &log.WailsLogger{SugaredLogger: logger},
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

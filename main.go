package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

type SwitchHosts struct {
	log logger.Logger
}

func NewSwitchHosts(log *logger.Logger) *SwitchHosts {
	return &SwitchHosts{
		log: *log,
	}
}

var assets embed.FS

func main() {

	app := NewApp()

	err := wails.Run(&options.App{
		Width:  1024,
		Height: 768,
		Title:  "SwitchHosts",
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Frameless:        true,
		DisableResize:    true,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

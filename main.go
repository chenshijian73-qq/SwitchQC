package main

import (
	"embed"

	_ "github.com/mattn/go-sqlite3"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	// "github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/runtime"
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

	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("File")
	// FileMenu.AddText("&Open", keys.CmdOrCtrl("o"), openFile)
	FileMenu.AddSeparator()
	FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(app.ctx)
	})

	AppMenu.Append(menu.EditMenu()) // on macos platform, we should append EditMenu to enable Cmd+C,Cmd+V,Cmd+Z... shortcut

	// mac.TitleBarDefault().UseToolbar = true

	err := wails.Run(&options.App{
		Width:  1024,
		Height: 768,
		Title:  "QuickCMD",
		Menu:   AppMenu,
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
		// Mac: &mac.Options{
		// 	TitleBar: mac.TitleBarDefault(),
		// },
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

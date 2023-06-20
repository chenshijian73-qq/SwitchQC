package main

import (
	"changeme/internal/config"
	"changeme/internal/db"
	"changeme/pkg/log"
	"embed"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	dbPool *sync.Pool
)

var assets embed.FS

func main() {
	if err := config.Init(); err != nil { //初始化config
		panic(err)
	}
	if err := log.Init(); err != nil {
		panic(err)
	}
	if err := db.Init(); err != nil { //初始化db
		panic(err)
	}
	//if err := dataInit.Init(); err != nil { //初始化db
	//	panic(err)
	//}

	//app := NewApp()
	//
	//AppMenu := menu.NewMenu()
	//FileMenu := AppMenu.AddSubmenu("File")
	//// FileMenu.AddText("&Open", keys.CmdOrCtrl("o"), openFile)
	//FileMenu.AddSeparator()
	//FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
	//	runtime.Quit(app.ctx)
	//})
	//
	//AppMenu.Append(menu.EditMenu()) // on macos platform, we should append EditMenu to enable Cmd+C,Cmd+V,Cmd+Z... shortcut
	//
	//// mac.TitleBarDefault().UseToolbar = true
	//
	//err := wails.Run(&options.App{
	//	Width:  1024,
	//	Height: 768,
	//	Title:  "QuickCMD",
	//	Menu:   AppMenu,
	//	AssetServer: &assetserver.Options{
	//		Assets: assets,
	//	},
	//	BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
	//	OnStartup:        app.startup,
	//	Frameless:        true,
	//	DisableResize:    true,
	//	Bind: []interface{}{
	//		app,
	//	},
	//	// Mac: &mac.Options{
	//	// 	TitleBar: mac.TitleBarDefault(),
	//	// },
	//})
	//
	//if err != nil {
	//	println("Error:", err.Error())
	//}
}

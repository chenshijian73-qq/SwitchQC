package main

import (
	"changeme/internal/config"
	"changeme/internal/db"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	dbPool *sync.Pool
)
var icon []byte
var assets embed.FS

func initData() {
	if err := config.Init(); err != nil { //初始化config
		panic(err)
	}

	if err := db.Init(); err != nil { //初始化db
		panic(err)
	}
	//if err := dataInit.Init(); err != nil { //初始化db
	//	panic(err)
	//}
}

func RunApp() {
	app := NewApp()

	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("File")
	// FileMenu.AddText("&Open", keys.CmdOrCtrl("o"), openFile)
	FileMenu.AddSeparator()
	FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(app.Ctx)
	})

	AppMenu.Append(menu.EditMenu()) // on macos platform, we should append EditMenu to enable Cmd+C,Cmd+V,Cmd+Z... shortcut

	// mac.TitleBarDefault().UseToolbar = true

	err := wails.Run(&options.App{
		Width:             1024,
		Height:            768,
		Title:             "QuickCMD",
		BackgroundColour:  &options.RGBA{R: 27, G: 38, B: 54, A: 0},
		DisableResize:     true,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       true,
		AlwaysOnTop:       true,
		HideWindowOnClose: true,
		Menu:              AppMenu,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:                        app.Startup,
		OnDomReady:                       app.DomReady,
		OnShutdown:                       app.Shutdown,
		CSSDragProperty:                  "--wails-draggable",
		CSSDragValue:                     "drag",
		EnableFraudulentWebsiteDetection: false,

		Bind: []interface{}{
			app,
		},
		// Windows平台特定选项
		Windows: &windows.Options{
			WebviewIsTransparent: false, // 背景是否半透明
			WindowIsTranslucent:  false, // 导航条是否半透明
			DisableWindowIcon:    false, // 是否关闭窗口上的图标
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			Appearance:           mac.NSAppearanceNameVibrantLight,
			About: &mac.AboutInfo{
				Title:   "Simple QuickCmd",
				Message: "For fun",
				Icon:    icon,
			},
		},
		Linux: &linux.Options{
			Icon:                icon,
			WindowIsTranslucent: false,
			WebviewGpuPolicy:    linux.WebviewGpuPolicyAlways,
		},
		Debug: options.Debug{
			OpenInspectorOnStartup: false,
		},
	})

	if err != nil {
		panic(err)
	}
}

func main() {
	initData()
	RunApp()
}

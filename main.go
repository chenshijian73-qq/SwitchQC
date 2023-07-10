package main

import (
	"changeme/internal/db"
	"changeme/internal/logic"
	"changeme/internal/logic/dataInit"
	"embed"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var icon []byte

//go:embed all:frontend/dist
var assets embed.FS

const QcPath = "~/.qc/"

func initData() {
	//if err := config.Init(); err != nil { //初始化config
	//	panic(err)
	//}

	if err := db.Init(); err != nil { //初始化db
		panic(err)
	}
	if err := logic.InitConfigQc(); err != nil {
		panic(err)
	}
	if err := dataInit.Init(); err != nil { //初始化db
		panic(err)
	}
}

func RunApp() {
	app := NewApp()
	recycle := NewRecycle()

	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("File")
	// FileMenu.AddText("&Open", keys.CmdOrCtrl("o"), openFile)
	FileMenu.AddSeparator()
	FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(app.Ctx)
	})

	AppMenu.Append(menu.EditMenu()) // on macos platform, we should append EditMenu to enable Cmd+C,Cmd+V,Cmd+Z... shortcut

	err := wails.Run(&options.App{
		Width:            1024,
		Height:           768,
		Title:            "QC",
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 0},
		//DisableResize:    true,
		Fullscreen: false,
		//Frameless:         false,
		//StartHidden:       true,
		AlwaysOnTop:       true,
		HideWindowOnClose: true,
		Menu:              AppMenu,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:                        app.startup,
		OnDomReady:                       app.domReady,
		OnShutdown:                       app.shutdown,
		CSSDragProperty:                  "--wails-draggable",
		CSSDragValue:                     "drag",
		EnableFraudulentWebsiteDetection: false,

		Bind: []interface{}{
			app,
			recycle,
		},
		// Windows平台特定选项
		Windows: &windows.Options{
			WebviewIsTransparent: false, // 背景是否半透明
			WindowIsTranslucent:  false, // 导航条是否半透明
			DisableWindowIcon:    false, // 是否关闭窗口上的图标
		},
		Mac: &mac.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			//Appearance:           mac.NSAppearanceNameAqua,
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
	})

	if err != nil {
		panic(err)
	}
}

func main() {
	initData()
	RunApp()
}

package main

import (
	"changeme/internal/logic"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
)

// App struct
type App struct {
	Ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.Ctx = ctx
}

func (a *App) DomReady(ctx context.Context) {

}

func (a *App) Shutdown(ctx context.Context) {

}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) AddFile() {
	fmt.Println("addhost3")
}

func (a *App) RemoveFile() {
	fmt.Println("deletehost")
}

func (a *App) EditFile() {
	fmt.Println("编辑 File")
}

func (a *App) ToggleFile() {
	fmt.Println("切换 File")
}

func (a *App) ViewFile() {
	fmt.Println("查看 File")
}

func (a *App) GetFiles() []logic.Qc {
	qcLogic := logic.NewQcLogic()
	qcs, err := qcLogic.GetQCList()
	if err != nil {
		log.Error(err)
	}
	return qcs
}

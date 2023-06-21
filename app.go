package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) domReady(ctx context.Context) {

}

func (a *App) shutdown(ctx context.Context) {

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

type QuickCmdFile struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (a *App) GetFiles() []QuickCmdFile {

	return nil
}

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

type Form struct {
	Name    string
	Content string
	Enabled bool
}

func (a *App) AddFile(form Form) bool {
	qcLogic := logic.NewQcLogic()
	err := qcLogic.CreateQC(logic.Qc{
		Name:     form.Name,
		Filepath: QcPath,
		Content:  form.Content,
		Enabled:  form.Enabled,
	})
	if err != nil {
		return false
	}
	return true
}

func (a *App) RemoveFile(data logic.Qc) bool {
	qcLogic := logic.NewQcLogic()
	err := qcLogic.DeleteQC(data)
	if err != nil {
		return false
	}
	return true
}

func (a *App) EditFile(data logic.Qc) bool {
	qcLogic := logic.NewQcLogic()
	err := qcLogic.UpdateQC(data)
	if err != nil {
		return false
	}
	return true
}

func (a *App) GetFiles() []logic.Qc {
	qcLogic := logic.NewQcLogic()
	qcs, err := qcLogic.GetQCList()
	if err != nil {
		log.Error(err)
	}
	return qcs
}

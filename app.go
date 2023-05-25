package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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

	dirPath := "files"

	var files []QuickCmdFile

	// 打开目录
	dir, err := os.Open(dirPath)
	if err != nil {
		println("Error:", err.Error())
		return nil
	}
	defer dir.Close()

	// 读取目录中的文件
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		println("Error:", err.Error())
		return nil
	}

	// 遍历文件，筛选出后缀名为 ".quickcmd" 的文件
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() && strings.HasSuffix(fileInfo.Name(), ".quickcmd") {
			fileContent, err := ioutil.ReadFile(filepath.Join(dirPath, fileInfo.Name()))
			if err != nil {
				println("Error:", err.Error())
				return nil
			}
			before := strings.TrimSuffix(fileInfo.Name(), ".quickcmd")
			files = append(files, QuickCmdFile{
				Name:    before,
				Content: string(fileContent),
			})
		}
	}

	return files
}

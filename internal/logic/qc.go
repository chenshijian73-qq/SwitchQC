package logic

import (
	"bufio"
	"changeme/internal/model"
	"changeme/internal/model/tables"
	"changeme/pkg/file"
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type Qc struct {
	ID       uint   `json:"ID"`
	Name     string `json:"name"`
	Filepath string `json:"filepath"`
	Content  string `json:"content"`
	Enabled  bool   `json:"enabled"`
}

type QcLogic struct {
}

func NewQcLogic() *QcLogic {
	return &QcLogic{}
}

func (qcLogic *QcLogic) CreateQC(qc Qc) (err error) {
	// 如果文件名末尾不带 .qc 扩展名，则添加它
	if filepath.Ext(qc.Name) != ".qc" {
		qc.Name = qc.Name + ".qc"
	}

	m := model.NewModels[tables.Qc]()
	m.Model = &tables.Qc{
		ID:       qc.ID,
		Filename: qc.Name,
		Content:  qc.Content,
		Path:     strings.Replace(qc.Filepath, "~", os.Getenv("HOME"), 1),
		Status:   qc.Enabled,
		CreateAt: gtime.New(time.Now()),
	}
	err = file.SaveFile(m.Model.Path+m.Model.Filename, qc.Content)
	if err != nil {
		return
	}
	err = m.Create()
	if err != nil {
		return
	}

	err = UpdateConfigQcFile()
	if err != nil {
		return
	}

	return
}

func (qcLogic *QcLogic) GetQCList() (qcs []Qc, err error) {
	m := model.NewModels[tables.Qc]()
	rows, err := m.GetsNoDeleted()
	if err != nil {
		log.Error("getQClist error")
		return nil, err
	}
	num := len(rows)
	if num > 0 {
		qcs = make([]Qc, num)
		for i, row := range rows {
			qcs[i].ID = row.ID
			qcs[i].Name = row.Filename
			qcs[i].Filepath = strings.Replace(row.Path, "~", os.Getenv("HOME"), 1)
			qcs[i].Enabled = row.Status
			//qcs[i].Content = row.Content
			fullPath := qcs[i].Filepath + qcs[i].Name
			if _, err := os.Stat(fullPath); os.IsNotExist(err) {
				log.Error(fmt.Sprintf("file %s dont exist\n", fullPath))
				continue
			}
			content, err := ioutil.ReadFile(fullPath)
			if err != nil {
				log.Error(fmt.Sprintf("read %s error: %s\n", fullPath, err))
				continue
			}
			qcs[i].Content = string(content)
		}
	}
	return
}

func (qcLogic *QcLogic) UpdateQC(qc Qc) (err error) {
	m := model.NewModels[tables.Qc]()
	m.Model = &tables.Qc{
		ID:       qc.ID,
		Filename: qc.Name,
		Content:  qc.Content,
		Path:     qc.Filepath,
		Status:   qc.Enabled,
		UpdateAt: gtime.New(time.Now()),
	}
	m.Update("filename", "path", "status", "content", "update_at")
	err = file.SaveFileIfModified(qc.Filepath+qc.Name, qc.Content)
	if err != nil {
		return
	}

	err = UpdateConfigQcFile()
	if err != nil {
		return
	}
	return
}

func (qcLogic *QcLogic) DeleteQC(qc Qc) (err error) {
	DeleteQcConfig(qc)

	m := model.NewModels[tables.Qc]()
	m.Model = &tables.Qc{
		ID:       qc.ID,
		Content:  qc.Content,
		DeleteAt: gtime.New(time.Now()),
	}
	m.Update("content", "delete_at")
	m.Get()
	err = file.DeleteFile(m.Model.Path + m.Model.Filename)

	return
}

func InitConfigQc() (err error) {
	homeDir, _ := os.UserHomeDir()
	qcConfigPath := homeDir + "/.qc/.qc"
	// 创建 ～/.qc/.qc文件
	qcConfigFile, err := os.OpenFile(qcConfigPath, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Error(err)
		return
	}
	qcConfigFile.Close()

	qcSource := fmt.Sprintf("\nsource %s", qcConfigPath)
	out, _ := exec.Command("bash", "-c", "echo $SHELL").CombinedOutput()
	shell := strings.Replace(string(out), "\n", "", -1)
	if shell == "/bin/bash" {
		bashPath := fmt.Sprintf("%s/.bashrc", homeDir)
		err = addConfigToShell(bashPath, qcSource)
	} else if shell == "/bin/zsh" {
		zshPath := fmt.Sprintf("%s/.zshrc", homeDir)
		err = addConfigToShell(zshPath, qcSource)
	} else {
		log.Error("(｡ì _ í｡) Sorry，cant not get which shell is using。")
		return fmt.Errorf("env SHELL=%s, shell bash/zsh dont find", shell)
	}

	err = UpdateConfigQcFile()
	if err != nil {
		return err
	}

	return
}

func UpdateConfigQcFile() (err error) {
	logic := NewQcLogic()
	qcs, err := logic.GetQCList()
	if len(qcs) > 1 {
		for _, qc := range qcs {
			if qc.Enabled {
				OpenQcConfig(qc)
			} else {
				DeleteQcConfig(qc)
			}
		}
		err = sourceConfig()
		if err != nil {
			return err
		}
	} else {
		log.Info("no qcFile using...")
	}
	return
}

func OpenQcConfig(qc Qc) (err error) {
	homeDir, _ := os.UserHomeDir()
	qcConfigPath := homeDir + "/.qc/.qc"

	sourceCmd := fmt.Sprintf("source %s%s\n", qc.Filepath, qc.Name)
	checkSourceCmd := fmt.Sprintf("cat %s |grep '%s'", qcConfigPath, strings.ReplaceAll(sourceCmd, "\n", ""))
	sourceIsExist, _ := exec.Command("bash", "-c", checkSourceCmd).CombinedOutput()
	if string(sourceIsExist) == "" {
		qcConfigFile, err := os.OpenFile(qcConfigPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
		if err != nil {
			log.Error(err)
			return err
		}
		defer qcConfigFile.Close()
		writer := bufio.NewWriter(qcConfigFile)
		writer.Write([]byte(sourceCmd))
		writer.Flush()
		log.Info(fmt.Sprintf("Add '%s' to %s success", strings.ReplaceAll(sourceCmd, "\n", ""), qcConfigPath))
	} else {
		log.Info(fmt.Sprintf("'%s' already in %s", strings.ReplaceAll(sourceCmd, "\n", ""), qcConfigPath))
	}
	return
}

func DeleteQcConfig(qc Qc) (err error) {
	homeDir, _ := os.UserHomeDir()
	qcConfigPath := homeDir + "/.qc/.qc"
	err = file.DeleteMatchRow(qcConfigPath, qc.Filepath+qc.Name)
	if err != nil {
		return
	}
	log.Info(fmt.Sprintf("delete %s in %s success", qc.Filepath+qc.Name, qcConfigPath))
	return
}

func addConfigToShell(shellConfigPath, qcSource string) error {
	checkSourceCmd := fmt.Sprintf(`cat %s |grep '%s'`, shellConfigPath, strings.ReplaceAll(qcSource, "\n", ""))
	sourceIsExist, _ := exec.Command("bash", "-c", checkSourceCmd).CombinedOutput()
	if string(sourceIsExist) == "" {
		configFile, err := os.OpenFile(shellConfigPath, os.O_WRONLY|os.O_APPEND, 0666)
		defer configFile.Close()
		if err != nil {
			return err
		}
		shellConfigWriter := bufio.NewWriter(configFile)
		shellConfigWriter.Write([]byte(qcSource))
		shellConfigWriter.Flush()
		log.Info(fmt.Sprintf("add '%s' to %s success", strings.ReplaceAll(qcSource, "\n", ""), shellConfigPath))
	} else {
		log.Info(fmt.Sprintf("'%s' already in %s", strings.ReplaceAll(qcSource, "\n", ""), shellConfigPath))
	}
	return nil
}

func sourceConfig() (err error) {
	homeDir, _ := os.UserHomeDir()

	out, _ := exec.Command("bash", "-c", "echo $SHELL").CombinedOutput()
	shell := strings.Replace(string(out), "\n", "", -1)
	var shellPath = ""
	if shell == "/bin/bash" {
		shellPath = fmt.Sprintf("%s/.bashrc", homeDir)
	} else if shell == "/bin/zsh" {
		shellPath = fmt.Sprintf("%s/.zshrc", homeDir)
	} else {
		log.Error("(｡ì _ í｡) Sorry，cant not get which shell is using。")
		return fmt.Errorf("env SHELL=%s, shell bash/zsh dont find", shell)
	}

	sourceCmd := fmt.Sprintf("source %s", shellPath)
	cmd := exec.Command("bash", "-c", sourceCmd)
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Error(fmt.Printf("cmd '%s' run error, combined out:\n%s\n", sourceCmd, string(out)))
		return err
	} else {
		log.Info(fmt.Sprintf("'%s' exec success", sourceCmd))
	}

	return
}

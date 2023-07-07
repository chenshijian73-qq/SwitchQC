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
	ID       uint   `json:"id"`
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

	err = file.SaveFile(strings.Replace(qc.Filepath, "~", os.Getenv("HOME"), 1)+qc.Name, qc.Content)
	if err != nil {
		return err
	}

	err = AddQcConfig(qc)
	if err != nil {
		return err
	}
	err = UpdateConfigQcFile(qc)
	if err != nil {
		return fmt.Errorf("添加文件失败，代码内容错误")
	}
	log.Info("testCreatQc")

	m := model.NewModels[tables.Qc]()
	m.Model = &tables.Qc{
		ID:       qc.ID,
		Filename: qc.Name,
		Content:  qc.Content,
		Path:     strings.Replace(qc.Filepath, "~", os.Getenv("HOME"), 1),
		Enabled:   qc.Enabled,
		CreateAt: gtime.New(time.Now()),
	}

	err = m.Create()
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
			qcs[i].Enabled = row.Enabled
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
		Enabled:   qc.Enabled,
		UpdateAt: gtime.New(time.Now()),
	}
	m.Update("filename", "path", "status", "content", "update_at")
	err = file.SaveFileIfModified(qc.Filepath+qc.Name, qc.Content)
	if err != nil {
		return
	}

	err = UpdateConfigQcFile(qc)
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
	m.Update("delete_at")
	m.Get()
	err = file.DeleteFile(m.Model.Path + m.Model.Filename)

	return
}

func InitConfigQc() (err error) {
	homeDir := os.Getenv("HOME")
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

	err = UpdateALlConfigQcFile()
	if err != nil {
		return err
	}

	return
}

// UpdateConfigQcFile 更新 ～/.qc/.qc 的内容
func UpdateConfigQcFile(qc Qc) (err error) {
	if qc.Enabled {
		CheckQcFileExist(qc)
		AddQcConfig(qc)
		err = sourceConfig()
		if err != nil {
			DeleteQcConfig(qc)
			return fmt.Errorf("代码有错，请检查")
		}
	} else {
		DeleteQcConfig(qc)
		sourceConfig()
	}
	return
}

// UpdateALlConfigQcFile 更新全部 ～/.qc/.qc 的内容
func UpdateALlConfigQcFile() (err error) {
	logic := NewQcLogic()
	qcs, err := logic.GetQCList()
	if len(qcs) > 1 {
		for _, qc := range qcs {
			if qc.Enabled {
				CheckQcFileExist(qc)
				err := AddQcConfig(qc)
				if err != nil {
					return err
				}
				err = sourceConfig()
				if err != nil {
					DeleteQcConfig(qc)
					return fmt.Errorf("代码有错，请检查")
				}
			} else {
				DeleteQcConfig(qc)
				sourceConfig()
			}
		}
	} else {
		log.Info("no qcFile using...")
	}
	return
}

// AddQcConfig 在 ～/.qc/.qc 添加 source $file
func AddQcConfig(qc Qc) (err error) {
	qcConfigPath := os.Getenv("HOME") + "/.qc/.qc"
	if qc.Name == "" {
		return fmt.Errorf("the file name cannot be empty")
	}
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

// DeleteQcConfig 在 ～/.qc/.qc 删除 source $file
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

// 添加 'source ～/.qc/.qc' 到 bashrc 或 zshrc
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

func CheckQcFileExist(qc Qc) (err error) {
	f := strings.Replace(qc.Filepath, "~", os.Getenv("HOME"), 1) + qc.Name
	if !file.IsExist(f) {
		err = file.SaveFile(f, qc.Content)
		if err != nil {
			return fmt.Errorf("file %s dont exist, create file err: %s", qc.Name, err)
		}
	}
	return nil
}

package logic

import (
	"bufio"
	"changeme/internal/model"
	"changeme/internal/model/tables"
	"changeme/pkg/file"
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type QcLogic struct {
}

func NewQcLogic() *QcLogic {
	return &QcLogic{}
}

func (qcLogic *QcLogic) CreateQC(qc tables.Qc) (err error) {
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

	m := model.NewModels[tables.Qc]()
	m.Model = &tables.Qc{
		ID:       qc.ID,
		Name:     qc.Name,
		Content:  qc.Content,
		Filepath: strings.Replace(qc.Filepath, "~", os.Getenv("HOME"), 1),
		Enabled:  qc.Enabled,
		CreateAt: gtime.New(time.Now()),
	}

	err = m.Create()
	if err != nil {
		return
	}

	return
}

func (qcLogic *QcLogic) GetQCList() (qcs []tables.Qc, err error) {
	m := model.NewModels[tables.Qc]()
	qcs, err = m.GetsNoDeleted()
	if err != nil {
		log.Error("get list error")
		return nil, err
	}
	return
}

func (qcLogic *QcLogic) GetRecycleList() (qcs []tables.Qc, err error) {
	m := model.NewModels[tables.Qc]()
	qcs, err = m.GetsDeleted()
	if err != nil {
		log.Error("GetRecycleList error")
		return nil, err
	}
	return
}

func (qcLogic *QcLogic) UpdateQC(qc tables.Qc) (err error) {
	m := model.NewModels[tables.Qc]()
	qc.UpdateAt = gtime.New(time.Now())
	m.Model = &qc
	m.Update("name", "path", "status", "content", "update_at")
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

func (qcLogic *QcLogic) DeleteQC(qc tables.Qc) (err error) {
	DeleteQcConfig(qc)

	m := model.NewModels[tables.Qc]()
	qc.DeleteAt = gtime.Now()
	m.Model = &qc
	m.Update("delete_at")

	m.Get()
	err = file.DeleteFile(m.Model.Filepath + m.Model.Name)

	return
}

func (qcLogic *QcLogic) RestoreQCFromBin(qc tables.Qc) (err error) {
	m := model.NewModels[tables.Qc]()
	qc.DeleteAt = nil
	m.Model = &qc
	err = m.Update("delete_at")

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

	return
}

func (qcLogic *QcLogic) DeleteQCFromBin(qc tables.Qc) (err error) {
	m := model.NewModels[tables.Qc]()
	m.Model = &qc
	err = m.Delete(qc.ID)
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
func UpdateConfigQcFile(qc tables.Qc) (err error) {
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

// AddQcConfig 添加 'source $file' 到 ～/.qc/.qc
func AddQcConfig(qc tables.Qc) (err error) {
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
func DeleteQcConfig(qc tables.Qc) (err error) {
	homeDir, _ := os.UserHomeDir()
	qcConfigPath := homeDir + "/.qc/.qc"
	err = file.DeleteMatchRow(qcConfigPath, qc.Name)
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

func CheckQcFileExist(qc tables.Qc) (err error) {
	f := strings.Replace(qc.Filepath, "~", os.Getenv("HOME"), 1) + qc.Name
	if !file.IsExist(f) {
		err = file.SaveFile(f, qc.Content)
		if err != nil {
			return fmt.Errorf("file %s dont exist, create file err: %s", qc.Name, err)
		}
	}
	return nil
}

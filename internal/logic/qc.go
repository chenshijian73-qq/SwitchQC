package logic

import (
	"changeme/internal/model"
	"changeme/internal/model/tables"
	"changeme/pkg/file"
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Qc struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Filepath string `json:"filepath"`
	Content  string `json:"content"`
	Status   bool   `json:"status"`
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
		Id:       qc.Id,
		Filename: qc.Name,
		Path:     strings.Replace(qc.Filepath, "~", os.Getenv("HOME"), 1),
		Status:   qc.Status,
		CreateAt: gtime.New(time.Now()),
	}
	err = m.Create()
	if err != nil {
		return
	}
	err = file.SaveFile(m.Model.Path+m.Model.Filename, qc.Content)
	return
}

func (qcLogic *QcLogic) GetQCList() (qcs []Qc, err error) {
	m := model.NewModels[tables.Qc]()
	rows, err := m.GetsNoDeleted()
	if err != nil {
		log.Error("getQClist error")
	}
	num := len(rows)
	if num > 0 {
		qcs = make([]Qc, num)
		for i, row := range rows {
			qcs[i].Id = row.Id
			qcs[i].Name = row.Filename
			qcs[i].Filepath = strings.Replace(row.Path, "~", os.Getenv("HOME"), 1)
			qcs[i].Status = row.Status
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
		Id:       qc.Id,
		Filename: qc.Name,
		Path:     qc.Filepath,
		Status:   qc.Status,
		UpdateAt: gtime.New(time.Now()),
	}
	m.Update("filename", "path", "status", "update_at")
	err = file.SaveFileIfModified(qc.Filepath+qc.Name, qc.Content)
	return
}

func (qcLogic *QcLogic) DeleteQC(qc Qc) (err error) {
	m := model.NewModels[tables.Qc]()
	m.Model = &tables.Qc{
		Id:       qc.Id,
		DeleteAt: gtime.New(time.Now()),
	}
	m.Update("delete_at")
	m.Get()
	err = file.DeleteFile(m.Model.Path + m.Model.Filename)
	return
}

package dataInit

import (
	"changeme/internal/model"
	"changeme/internal/model/tables"
	"github.com/gogf/gf/v2/os/gtime"
	"time"
)

var (
	datas = []tables.Qc{
		{
			Filename:  "open.qc",
			Path:      "~/.qc/",
			Status:    true,
			IsDeleted: false,
			CreateAt:  gtime.New(time.Now()),
		},
		{
			Filename:  "git.qc",
			Path:      "~/.qc/",
			Status:    true,
			IsDeleted: false,
			CreateAt:  gtime.New(time.Now()),
		},
		{
			Filename:  "test.qc",
			Path:      "~/.qc/",
			Status:    false,
			IsDeleted: false,
			CreateAt:  gtime.New(time.Now()),
		},
	}
)

func Init() error {
	err := qcInitData()
	return err
}

func qcInitData() error {
	qc := model.NewModels[tables.Qc]()
	_, err := qc.Creates(datas...)
	return err
}

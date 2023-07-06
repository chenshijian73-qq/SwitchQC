package dataInit

import (
	"changeme/internal/model"
	"changeme/internal/model/tables"
	"changeme/pkg/file"
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

var (
	home  = os.Getenv("HOME")
	datas = []tables.Qc{
		{
			Filename:  "open.qc",
			Path:      home + "/.qc/",
			Content:   "alias sublime='open -a /Applications/Sublime\\ Text.app'\nalias typora='open -a  /Applications/Typora.app'\nalias golang='open -a /Applications/GoLand.app'\nalias execl='open -a /Applications/Microsoft\\ Excel.app'\n",
			Status:    true,
			IsDeleted: false,
			CreateAt:  gtime.New(time.Now()),
		},
		{
			Filename:  "git.qc",
			Path:      home + "/.qc/",
			Content:   "# alias for git\nalias gcl=\"git clone\"\nalias gc=\"git checkout\"\nalias gcb=\"git checkout -b\"\nalias gb=\"git branch\"\nalias gm=\"git commit\"\nalias gp=\"git push\"\nalias gcp=\"git cherry-pick\"\n",
			Status:    true,
			IsDeleted: false,
			CreateAt:  gtime.New(time.Now()),
		},
		{
			Filename:  "test.qc",
			Path:      home + "/.qc/",
			Content:   "",
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

func qcInitData() (err error) {
	qc := model.NewModels[tables.Qc]()
	//_, err := qc.Creates(datas...)
	for _, data := range datas {
		row, err := qc.GetByName(data.Filename)
		if err != nil {
			return err
		}
		if row.Filename == "" {
			qc.Model = &data
			err = qc.Create()
			if err != nil {
				return err
			}
			file.CreateFile(data.Path+data.Filename, data.Content)
			log.Info(fmt.Sprintf("create init qc file %s", data.Filename))
		} else {
			log.Info(fmt.Sprintf("file %s is already existed", row.Filename))
		}
	}

	return err
}

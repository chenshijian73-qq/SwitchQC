package db

import (
	"changeme/internal/model/tables"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

const (
	dbFileName      = "qc.db"
	maxIdleCons     = 10
	maxOpenCons     = 100
	connMaxLifetime = 280
	dbFileDir       = ".qc"
)

var (
	Db *gorm.DB
)

func Init() (err error) {

	dbFilePath, err := CreateQcDB()
	if err != nil {
		return err
	}

	if Db, err = gorm.Open(sqlite.Open(dbFilePath), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)}); err != nil {
		return err
	}
	sqlDb, err := Db.DB()
	if err != nil {
		return
	}
	//Db.LogMode(true)
	sqlDb.SetMaxIdleConns(maxIdleCons)
	sqlDb.SetMaxOpenConns(maxOpenCons)
	sqlDb.SetConnMaxLifetime(time.Duration(10) * time.Second)

	err = autoMigrate() //自动迁移
	if err != nil {
		return
	}

	return err
}

func autoMigrate() (err error) {
	return Db.AutoMigrate(
		&tables.Event{},
		&tables.Qc{},
	)
}

func CreateQcDB() (qcDBFile string, err error) {
	currentUser, err := user.Current()
	if err != nil {
		log.Error(fmt.Sprintf("get current user error：%s", err))
		return
	}
	homeDir := currentUser.HomeDir
	qcDir := filepath.Join(homeDir, dbFileDir)
	qcDBFile = filepath.Join(qcDir, dbFileName)
	// 判断 .qc 文件夹是否存在，如果不存在，则创建
	if _, err := os.Stat(qcDir); os.IsNotExist(err) {
		if err := os.Mkdir(qcDir, 0755); err != nil {
			log.Error(fmt.Sprintf("create ~/.qc dir failed：%s", err))
			return "", err
		}
		log.Info("create ~/.qc dir success")
	}

	// 判断 qc.db 文件是否存在，如果不存在，则创建
	if _, err := os.Stat(qcDBFile); os.IsNotExist(err) {
		if _, err := os.Create(qcDBFile); err != nil {
			log.Error(fmt.Sprintf("create qc.db file failed：%s", err))
			return "", err
		}
		log.Info("create qc.db success")
	}

	return qcDBFile, err
}

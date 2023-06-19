package db

import (
	"changeme/internal/config"
	"changeme/internal/model/tables"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var Db *gorm.DB

func Init() (err error) {
	conf := config.Conf.Db
	dbFilePath := conf.FilePath
	if Db, err = gorm.Open(sqlite.Open(dbFilePath), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)}); err != nil {
		return err
	}
	sqlDb, err := Db.DB()
	if err != nil {
		return
	}
	//Db.LogMode(true)
	sqlDb.SetMaxIdleConns(conf.MaxIdleCons)
	sqlDb.SetMaxOpenConns(conf.MaxOpenCons)
	sqlDb.SetConnMaxLifetime(time.Duration(10) * time.Second)

	return autoMigrate() //自动迁移
}

func autoMigrate() (err error) {
	return Db.AutoMigrate(
		&tables.Event{},
		&tables.Qc{},
	)
}

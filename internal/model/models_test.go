package model

import (
	"changeme/internal/model/tables"
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	sqliteMock "github.com/lichmaker/sqlite-mock"
	"gorm.io/gorm"
	"testing"
)

var (
	mock sqlmock.Sqlmock
	DB   *gorm.DB
)

func TestMain(m *testing.M) {
	var db *sql.DB
	var err error
	db, mock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	myMock := sqliteMock.Open(db)

	dbSqlite, err := gorm.Open(myMock, &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("failed to connect mock database :" + err.Error())
	}
	DB = dbSqlite

	m.Run()
}

func TestQueryQc(t *testing.T) {

	qc := NewModels[tables.Qc]()
	qc.Model = &tables.Qc{
		ID:       798,
		Filename: "hello",
	}
	err := qc.Create()
	if err != nil {
		return
	}
	fmt.Println(qc.Model.ID)
}

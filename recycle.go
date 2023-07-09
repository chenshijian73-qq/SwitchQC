package main

import (
	"changeme/internal/logic"
	"changeme/internal/model/tables"
	"context"
	log "github.com/sirupsen/logrus"
)

// NewRecycle struct
type Recycle struct {
	Ctx context.Context
}

// NewRecycle creates a new NewRecycle application struct
func NewRecycle() *Recycle {
	return &Recycle{}
}

func (*Recycle) RestoreFileFromBin(data tables.Qc) (error string) {
	log.Info(data)
	qcLogic := logic.NewQcLogic()
	err := qcLogic.RestoreQCFromBin(data)
	if err != nil {
		log.Error(err)
		return err.Error()
	}
	return ""
}

func (*Recycle) CleanFileFromBin(data tables.Qc) (error string) {
	log.Info(data)
	qcLogic := logic.NewQcLogic()
	err := qcLogic.DeleteQCFromBin(data)
	if err != nil {
		log.Error(err)
		return err.Error()
	}
	return ""
}

func (*Recycle) GetRecycleList() []tables.Qc {
	qcLogic := logic.NewQcLogic()
	qcs, err := qcLogic.GetRecycleList()
	if err != nil {
		log.Error(err)
	}
	return qcs
}

func (*Recycle) CleanBin() (error string) {
	//qcLogic := logic.NewQcLogic()
	//err := qcLogic.CleanBin()
	//if err != nil {
	//	log.Error(err)
	//	return err.Error()
	//}
	return
}

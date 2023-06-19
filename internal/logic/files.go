package logic

import (
	"changeme/internal/dao"
	"changeme/internal/model/do"
	"context"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
)

type QcOperator struct {
}

var (
	Qc = QcOperator{}
)

func init() {
	ctx := gctx.New()
	Qc.Update(ctx, do.QcFile{
		Filename: "qc",
		Path:     "/user",
		Status:   false,
	})
}

func (o *QcOperator) Add(ctx context.Context, qc do.QcFile) (err error) {
	qc.CreateAt, qc.UpdateAt = gtime.New(), gtime.Now()
	_, err = dao.QcFile.Ctx(ctx).Insert()
	return err
}

func (o *QcOperator) Update(ctx context.Context, qc do.QcFile) (err error) {
	qc.UpdateAt = gtime.New()
	_, err = dao.QcFile.Ctx(ctx).Update(qc, "id=?", qc.Id)
	return err
}
func (o *QcOperator) Delete(ctx context.Context, qc do.QcFile) (err error) {
	qc.DeleteAt, qc.UpdateAt = gtime.New(), gtime.Now()
	_, err = dao.QcFile.Ctx(ctx).Update(qc, "id=?", qc.Id)
	return err
}

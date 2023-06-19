package tables

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type Qc struct {
	Id        int64       `json:"id"`        //
	Filename  string      `json:"filename"`  //
	Path      string      `json:"path"`      //
	Status    string      `json:"status"`    //
	IsDeleted bool        `json:"isDeleted"` //
	CreateAt  *gtime.Time `json:"createAt"`  //
	UpdateAt  *gtime.Time `json:"updateAt"`  //
	DeleteAt  *gtime.Time `json:"deleteAt"`  //
}

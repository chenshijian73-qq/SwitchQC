package tables

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type Qc struct {
	ID        uint        `gorm:"primary_key"`
	Filename  string      `json:"filename"` //
	Content   string      `gorm:"type:text;"`
	Path      string      `json:"path"`                       //
	Enabled    bool        `gorm:"type:boolean;default:true"`  //
	IsDeleted bool        `gorm:"type:boolean;default:false"` //
	CreateAt  *gtime.Time `json:"createAt"`                   //
	UpdateAt  *gtime.Time `json:"updateAt"`                   //
	DeleteAt  *gtime.Time `json:"deleteAt"`                   //
}

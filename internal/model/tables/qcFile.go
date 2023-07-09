package tables

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type Qc struct {
	ID       uint        `gorm:"primaryKey"`
	Name     string      `gorm:"column:name"`
	Content  string      `gorm:"type:text"`
	Filepath string      `gorm:"column:filepath"`
	Enabled  bool        `gorm:"type:boolean;default:true"`
	CreateAt *gtime.Time `json:"CreateAt"` //
	UpdateAt *gtime.Time `json:"UpdateAt"` //
	DeleteAt *gtime.Time `json:"DeleteAt"` //
}

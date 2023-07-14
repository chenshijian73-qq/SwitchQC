package tables

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type AiLink struct {
	ID       uint        `gorm:"primaryKey"`
	Name     string      `gorm:"column:name"`
	Link     string      `gorm:"column:link"`
	Enabled  bool        `gorm:"type:boolean;default:false"`
	CreateAt *gtime.Time `json:"CreateAt"` //
	UpdateAt *gtime.Time `json:"UpdateAt"` //
	DeleteAt *gtime.Time `json:"DeleteAt"` //
}

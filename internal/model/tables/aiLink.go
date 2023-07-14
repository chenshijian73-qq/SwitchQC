package tables

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type AiLink struct {
	ID       uint        `gorm:"primaryKey"`
	Link     string      `gorm:"column:link"`
	CreateAt *gtime.Time `json:"CreateAt"` //
	UpdateAt *gtime.Time `json:"UpdateAt"` //
	DeleteAt *gtime.Time `json:"DeleteAt"` //
}

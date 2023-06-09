// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Files is the golang structure for table files.
type Files struct {
	Id       int         `json:"id"       ` //
	Filename string      `json:"filename" ` //
	Status   bool        `json:"status"   ` //
	CreateAt *gtime.Time `json:"createAt" ` //
	UpdateAt *gtime.Time `json:"updateAt" ` //
	DeleteAt *gtime.Time `json:"deleteAt" ` //
}

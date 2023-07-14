package model

import (
	"changeme/internal/db"
	"changeme/internal/model/tables"
	"encoding/json"
	"gorm.io/gorm"
)

type Type interface {
	tables.Qc | tables.AiLink
}

type Models[T Type] struct {
	Model *T
	db    *gorm.DB
}

func NewModels[T Type]() *Models[T] {
	return &Models[T]{
		Model: new(T),
		db:    db.Db,
	}
}

func (t *Models[T]) Create() (err error) {
	return t.db.Create(t.Model).Error
}

func (t *Models[T]) Creates(ms ...T) (ids []T, err error) { //创建后id需要自己遍历出来
	err = t.db.Create(ms).Error
	ids = ms
	return
}

func (t *Models[T]) Begin() *gorm.DB {
	t.db = t.db.Begin()
	return t.db
}

func (t *Models[T]) SetTX(tx *gorm.DB) *Models[T] {
	t.db = tx
	return t
}

func (t *Models[T]) Save() (err error) {
	return t.db.Omit("created_by").Save(t.Model).Error
}

func (t *Models[T]) Delete(id uint) (err error) {
	return t.db.Where("id = ?", id).Delete(t.Model).Error
}

func (t *Models[T]) Update(columns ...interface{}) (err error) {
	dbs := t.db.Model(t.Model)
	if len(columns) > 0 {
		dbs = dbs.Select(columns[0], columns[1:]...)
	}
	return dbs.Updates(t.Model).Error
}

func (t *Models[T]) Get() (err error) {
	if err = t.db.First(t.Model, t.Model).Error; err != nil && err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}

func (t *Models[T]) GetNoDeleted() (err error) {
	if err = t.db.Where("delete_at IS NULL").First(t.Model, t.Model).Error; err != nil && err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}

func (t *Models[T]) Gets() (row []T, err error) {
	err = t.db.Find(&row, t.Model).Error
	return
}

func (t *Models[T]) GetsNoDeleted() (row []T, err error) {
	err = t.db.Where("delete_at IS NULL").Find(&row, t.Model).Error
	return
}

func (t *Models[T]) GetsDeleted() (row []T, err error) {
	err = t.db.Where("delete_at IS NOT NULL").Find(&row, t.Model).Error
	return
}

func (t *Models[T]) GetsByIds(ids []int64) (row []T, err error) {
	err = t.db.Where("id IN (?)", ids).Find(&row).Error
	return
}
func (t *Models[T]) GetByName(name string) (row T, err error) {
	err = t.db.Where("name = (?)", name).Find(&row).Error
	return
}

func (t *Models[T]) ToJson() (row []byte, err error) {
	return json.Marshal(t)
}

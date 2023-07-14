package logic

import (
	"changeme/internal/model"
	"changeme/internal/model/tables"
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	log "github.com/sirupsen/logrus"
	"time"
)

type AiLinkLogic struct {
	Ctx context.Context
}

func NewAiLinkLogic() *AiLinkLogic {
	return &AiLinkLogic{}
}

func (l *AiLinkLogic) SaveAiLink(data tables.AiLink) (err error) {
	// 如
	m := model.NewModels[tables.AiLink]()
	m.Model = &data
	m.Model.CreateAt = gtime.New(time.Now())

	err = m.Create()
	if err != nil {
		return
	}

	return
}

func (l *AiLinkLogic) GetAiLinkList() (links []tables.AiLink, err error) {
	m := model.NewModels[tables.AiLink]()
	links, err = m.GetsNoDeleted()
	if err != nil {
		log.Error("get list error")
		return nil, err
	}
	return
}

func (l *AiLinkLogic) UpdateAiLink(link tables.AiLink) (err error) {
	m := model.NewModels[tables.AiLink]()
	link.UpdateAt = gtime.New(time.Now())
	m.Model = &link
	m.Update("link", "update_at")

	return
}

func (l *AiLinkLogic) DeleteAiLink(link tables.AiLink) (err error) {

	m := model.NewModels[tables.AiLink]()
	link.DeleteAt = gtime.Now()
	m.Model = &link
	m.Update("delete_at")

	return
}

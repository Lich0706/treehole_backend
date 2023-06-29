package dao

import (
	"TreeHole/treehole_backend/internal/dao/model"
	"TreeHole/treehole_backend/utils"
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MessageDao struct {
	DB *gorm.DB
}

func NewMessageDao(DB *gorm.DB) *MessageDao {
	return &MessageDao{DB: DB}
}

func (m *MessageDao) ListAll(ctx context.Context) ([]*model.Message, error) {
	messages := []*model.Message{}
	res := m.DB.Find(&messages)
	if res.Error != nil {
		utils.Logger.Error("NewMessageDao.ListAll",
			zap.String("error", res.Error.Error()))
	}
	return messages, res.Error
}

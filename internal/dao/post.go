package dao

import (
	"TreeHole/treehole_backend/internal/dao/model"
	"TreeHole/treehole_backend/utils"
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PostDao struct {
	DB *gorm.DB
}

func NewPostDao(DB *gorm.DB) *PostDao {
	return &PostDao{DB: DB}
}

func (m *PostDao) ListAll(ctx context.Context) ([]*model.Post, error) {
	posts := []*model.Post{}
	res := m.DB.Find(&posts)
	if res.Error != nil {
		utils.Logger.Error("NewPostDao.ListAll",
			zap.String("error", res.Error.Error()))
	}
	return posts, res.Error
}

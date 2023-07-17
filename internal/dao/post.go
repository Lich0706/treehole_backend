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

func (p *PostDao) CreateNewPost(ctx context.Context, post model.Post) error {
	res := p.DB.Create(&post)
	if res.Error != nil {
		utils.Logger.Error("NewPostDao.CreateNewPost",
			zap.String("error", res.Error.Error()))
	}
	return res.Error
}

func (p *PostDao) ListAll(ctx context.Context) ([]*model.Post, error) {
	posts := []*model.Post{}
	res := p.DB.Find(&posts)
	if res.Error != nil {
		utils.Logger.Error("NewPostDao.ListAll",
			zap.String("error", res.Error.Error()))
	}
	return posts, res.Error
}

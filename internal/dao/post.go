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

func (p *PostDao) CreateNewPost(ctx context.Context, post model.Post) (int64, error) {
	res := p.DB.Create(&post)
	if res.Error != nil {
		utils.Logger.Error("NewPostDao.CreateNewPost",
			zap.String("error", res.Error.Error()))
	}
	return post.ID, res.Error
}

func (p *PostDao) ListAll(ctx context.Context) ([]*model.Post, error) {
	posts := []*model.Post{}
	res := p.DB.Order("created_at DESC").Find(&posts)
	if res.Error != nil {
		utils.Logger.Error("NewPostDao.ListAll",
			zap.String("error", res.Error.Error()))
	}
	return posts, res.Error
}

func (p *PostDao) GetPostById(ctx context.Context, pid int64) (*model.Post, error) {
	post := model.Post{}
	res := p.DB.Where("id = ?", pid).First(&post)
	if res.Error != nil {
		utils.Logger.Error("NewPostDao.GetPostById",
			zap.String("error", res.Error.Error()))
	}
	return &post, res.Error
}

func (p *PostDao) UpdateReplyCountById(ctx context.Context, pid int64) (*model.Post, error) {
	post := model.Post{}
	res := p.DB.Where("id = ?", pid).First(&post)
	if res.Error != nil {
		utils.Logger.Error("NewPostDao.GetPostById",
			zap.String("error", res.Error.Error()))
	}
	post.ReplyCount += 1
	p.DB.Save(&post)
	return &post, res.Error
}

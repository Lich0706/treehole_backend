package dao

import (
	"TreeHole/treehole_backend/internal/dao/model"
	"TreeHole/treehole_backend/utils"
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CommentDao struct {
	DB *gorm.DB
}

func NewCommentDao(DB *gorm.DB) *CommentDao {
	return &CommentDao{DB: DB}
}

func (m *CommentDao) GetCommentsByPid(ctx context.Context, pid int64) ([]*model.Comment, error) {
	comments := []*model.Comment{}
	res := m.DB.Where(&model.Comment{
		PostID: pid,
	}).Find(&comments)
	if res.Error != nil {
		utils.Logger.Error("NewCommentDao.GetCommentsByPid",
			zap.String("error", res.Error.Error()))
	}
	return comments, res.Error
}

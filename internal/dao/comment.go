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

func (m *CommentDao) CreateNewComment(ctx context.Context, comment model.Comment) (int64, error) {
	res := m.DB.Create(&comment)
	if res.Error != nil {
		utils.Logger.Error("NewCommentDao.CreateNewComment",
			zap.String("error", res.Error.Error()))
	}
	return comment.ID, res.Error
}

func (m *CommentDao) GetCommentsByPidUid(ctx context.Context, pid int64, uid int64) (*model.Comment, error) {
	comment := model.Comment{}
	res := m.DB.Where("user_id = ? AND post_id=?", uid, pid).First(&comment)
	if res.Error != nil {
		utils.Logger.Error("NewCommentDao.GetCommentsByPidUid",
			zap.String("error", res.Error.Error()))
	}
	return &comment, res.Error
}

func (m *CommentDao) GetCommentsByPid(ctx context.Context, pid int64) ([]*model.Comment, error) {
	comments := []*model.Comment{}
	res := m.DB.Where("post_id = ?", pid).Find(&comments)
	if res.Error != nil {
		utils.Logger.Error("NewCommentDao.GetCommentsByPid",
			zap.String("error", res.Error.Error()))
	}
	return comments, res.Error
}

func (m *CommentDao) CountExceptCommenter(ctx context.Context, pid int64, uid int64) (int64, error) {
	var count int64
	res := m.DB.Where("post_id = ? AND user_id != ?", pid, uid).Distinct("user_id").Count(&count)
	if res.Error != nil {
		utils.Logger.Error("NewCommentDao.GetCommentsByPid",
			zap.String("error", res.Error.Error()))
	}
	return count, res.Error
}

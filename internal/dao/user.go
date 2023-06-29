package dao

import (
	"TreeHole/treehole_backend/internal/dao/model"
	"TreeHole/treehole_backend/utils"
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserDao struct {
	DB *gorm.DB
}

func NewUserDao(DB *gorm.DB) *UserDao {
	return &UserDao{DB: DB}
}

func (u *UserDao) ListAll(ctx context.Context) ([]*model.User, error) {
	users := []*model.User{}
	res := u.DB.Find(&users)
	if res.Error != nil {
		utils.Logger.Error("NewUserDao.ListAll",
			zap.String("error", res.Error.Error()))
	}
	return users, res.Error
}

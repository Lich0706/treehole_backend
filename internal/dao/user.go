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

func (u *UserDao) CreateNormalUser(ctx context.Context, user model.User) error {
	res := u.DB.Create(&user)
	if res.Error != nil {
		utils.Logger.Error("NewUserDao.CreateOrUpdate",
			zap.String("error", res.Error.Error()))
	}
	return res.Error
}

func (u *UserDao) FindOne(ctx context.Context, user model.User) (*model.User, error) {
	var result model.User
	res := u.DB.Model(&user).First(&result)
	if res.Error != nil {
		utils.Logger.Error("NewUserDao.FindOne",
			zap.String("error", res.Error.Error()))
	}
	return &result, res.Error
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

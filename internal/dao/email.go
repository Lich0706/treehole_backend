package dao

import (
	"TreeHole/treehole_backend/internal/dao/model"
	"TreeHole/treehole_backend/utils"
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type EmailDao struct {
	DB *gorm.DB
}

func NewEmailDao(DB *gorm.DB) *EmailDao {
	return &EmailDao{DB: DB}
}

func (m *EmailDao) Create(ctx context.Context, email string) error {
	res := m.DB.Create(&model.Email{
		Email: email,
	})
	if res.Error != nil {
		utils.Logger.Error("NewHashedEmailDao.Create",
			zap.String("error", res.Error.Error()))
	}
	return res.Error
}

func (m *EmailDao) FindOne(ctx context.Context, email string) error {
	res := m.DB.First(&model.Email{
		Email: email,
	})
	if res.Error != nil {
		utils.Logger.Error("NewHashedEmailDao.Find",
			zap.String("error", res.Error.Error()))
	}
	return res.Error
}

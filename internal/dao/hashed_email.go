package dao

import (
	"TreeHole/treehole_backend/internal/dao/model"
	"TreeHole/treehole_backend/utils"
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type HashedEmailDao struct {
	DB *gorm.DB
}

func NewHashedEmailDao(DB *gorm.DB) *HashedEmailDao {
	return &HashedEmailDao{DB: DB}
}

func (m *HashedEmailDao) Create(ctx context.Context, hashedEmail string) error {
	res := m.DB.Create(&model.HashedEmail{
		HashedEmail: hashedEmail,
	})
	if res.Error != nil {
		utils.Logger.Error("NewHashedEmailDao.Create",
			zap.String("error", res.Error.Error()))
	}
	return res.Error
}

func (m *HashedEmailDao) FindOne(ctx context.Context, hashedEmail string) error {
	res := m.DB.First(&model.HashedEmail{
		HashedEmail: hashedEmail,
	})
	if res.Error != nil {
		utils.Logger.Error("NewHashedEmailDao.Find",
			zap.String("error", res.Error.Error()))
	}
	return res.Error
}

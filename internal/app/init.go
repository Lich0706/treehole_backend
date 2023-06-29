package app

import (
	"TreeHole/treehole_backend/internal/dao"
	"TreeHole/treehole_backend/internal/db"

	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	MessageDao *dao.MessageDao
)

func InitDao() {
	DB = db.DB
	MessageDao = dao.NewMessageDao(DB)
}

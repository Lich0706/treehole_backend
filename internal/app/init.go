package app

import (
	"TreeHole/treehole_backend/internal/dao"
	"TreeHole/treehole_backend/internal/db"

	"gorm.io/gorm"
)

var (
	DB       *gorm.DB
	UserDao  *dao.UserDao
	PostDao  *dao.PostDao
	EmailDao *dao.EmailDao
)

func InitDao() {
	DB = db.DB
	UserDao = dao.NewUserDao(DB)
	PostDao = dao.NewPostDao(DB)
	EmailDao = dao.NewEmailDao(DB)
}

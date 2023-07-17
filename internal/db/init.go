package db

import (
	"TreeHole/treehole_backend/internal/dao/model"
	"TreeHole/treehole_backend/utils"
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init(config *viper.Viper) {
	var err error

	dbName := config.GetString("db.name")
	user := config.GetString("db.user")
	pwd := config.GetString("db.pwd")
	host := config.GetString("db.host")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		pwd,
		host,
		dbName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	InitTables(DB)
	if err != nil {
		utils.Logger.Error("Init mysql db",
			zap.String("error", err.Error()))
	} else {
		utils.Logger.Info("Init mysql successfully")
	}
	// return DB, err
}

func InitTables(DB *gorm.DB) {
	if !(DB.Migrator().HasTable(&model.Post{})) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(&model.Post{})
	}
	if !(DB.Migrator().HasTable(&model.User{})) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(&model.User{})
	}
	if !(DB.Migrator().HasTable(&model.HashedEmail{})) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(&model.HashedEmail{})
	}
}

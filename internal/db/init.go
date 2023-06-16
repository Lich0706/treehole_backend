package db

import (
	utils "TreeHole/treehole_backend/util"
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
	if err != nil {
		utils.Logger.Error("Init mysql db",
			zap.String("error", err.Error()))
	} else {
		utils.Logger.Info("Init mysql successfully")
	}
}

package main

import (
	"TreeHole/treehole_backend/config"
	"TreeHole/treehole_backend/internal/db"
	utils "TreeHole/treehole_backend/util"

	"go.uber.org/zap"
)

var (
	Version = "0.1.0"
)

func main() {
	// Initialize logger
	utils.InitializeLogger()
	// Load Config
	config, err := config.Init()
	if err != nil {
		utils.Logger.Error(err.Error())
	}
	utils.Logger.Info("Starting application",
		zap.String("ver", Version))
	utils.Logger.Info("Showing Configs")
	utils.Logger.Info("Configs",
		zap.String("env", config.GetString("app.env")))

	// Init Server
	port := config.GetInt("server.port")
	utils.Logger.Info("Configs",
		zap.Int("port", port))

	// Init DB
	db.Init(config)

	defer func() {
		utils.Logger.Info("Closing mysql connection")
		// Return a db generic interface
		dbInstance, _ := db.DB.DB()
		err := dbInstance.Close()
		if err != nil {
			utils.Logger.Error("Closing mysql fail",
				zap.String("error", err.Error()))
		} else {
			utils.Logger.Info("Close successfully")
		}
	}()

}

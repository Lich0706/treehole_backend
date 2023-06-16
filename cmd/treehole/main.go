package main

import (
	"TreeHole/treehole_backend/config"
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
	utils.Logger.Info("Showing Configs")
	utils.Logger.Info("Configs",
		zap.Any("port", config.Get("port")),
		zap.String("env", config.GetString("app.env")))

	// Init DB

}

package main

import (
	"TreeHole/treehole_backend/config"
	"TreeHole/treehole_backend/internal/api/router"
	"TreeHole/treehole_backend/internal/db"
	utils "TreeHole/treehole_backend/util"
	"fmt"
	"net/http"

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
	router := router.InitRouter()
	port := config.GetInt("server.port")
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
	s.ListenAndServe()

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

package main

import (
	"github.com/RhoNit/trigger_mgmt_backend/common"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	logger, err := common.InitLogger()
	if err != nil {
		logger.Error("Failed to initialize logger", zap.Error(err))
	}

	err = godotenv.Load()
	if err != nil {
		logger.Error("Failed to load environment variables", zap.Error(err))
	}

	dbConn, err := common.InitDatabase()
	if err != nil {
		logger.Error("Failed to connect to database", zap.Error(err))
	}
}

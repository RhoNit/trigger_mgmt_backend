package common

import "go.uber.org/zap"

func InitLogger() (*zap.Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	
	return logger, nil
}

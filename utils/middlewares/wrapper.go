package middlewares

import (
	"go-web-sample/utils/config"
	"go-web-sample/utils/logger"
)

// Wrapper represents a interface for accessing the data which sharing in overall application.
type Wrapper interface {
	GetConfig() *config.EnvConfig
	GetLogger() *logger.Logger
}

// wrapper struct is for sharing data which such as database setting, the setting of application and logger in overall this application.
type wrapper struct {
	config *config.EnvConfig
	logger *logger.Logger
}

// NewWrapper is constructor.
func NewWrapper(config *config.EnvConfig, logger *logger.Logger) Wrapper {
	return &wrapper{config: config, logger: logger}
}

// GetConfig returns the object of configuration.
func (c *wrapper) GetConfig() *config.EnvConfig {
	return c.config
}

// GetLogger returns the object of logger.
func (c *wrapper) GetLogger() *logger.Logger {
	return c.logger
}

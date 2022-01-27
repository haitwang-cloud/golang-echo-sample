package config

import (
	"fmt"
	"github.com/jinzhu/configor"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type EnvConfig struct {
	Database struct {
		Dialect   string `json:"dialect"`
		Host      string `json:"host"`
		Port      int    `json:"port"`
		Dbname    string `json:"dbname"`
		Username  string `json:"username"`
		Password  string `json:"password"`
		Migration bool   `json:"migration"`
	} `json:"database"`
	Env struct {
		Environment string `json:"environment"`
	} `json:"env"`
	Log struct {
		RequestLogFormat string `json:"request_log_format"`
	} `json:"log"`
	ZapConfig zap.Config        `json:"zap_config" yaml:"zap_config"`
	LogRotate lumberjack.Logger `json:"log_rotate" yaml:"log_rotate"`
}

const (
	// DEV represents development environment
	DEV = "develop"
	// PRD represents production environment
	PRD = "production"
)

// Load reads the settings written to the yml file
func Load() *EnvConfig {

	config := &EnvConfig{}
	if err := configor.Load(config, "application.yml"); err != nil {
		fmt.Printf("Failed to read application.yml: %s", err)
		os.Exit(2)
	}
	return config
}

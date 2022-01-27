package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
	"go-web-sample/utils/config"
	"go-web-sample/utils/logger"
	"go-web-sample/utils/middlewares"
)

func main() {
	// Echo instance
	e := echo.New()

	//load env
	envConfig := config.Load()
	// Init logger
	zapLogger := logger.NewLogger(envConfig)
	zapLogger.GetZapLogger().Infof("Loaded this configuration : application.yml")

	// Middleware
	wrapper := middlewares.NewWrapper(envConfig, zapLogger)
	middlewares.InitMiddleware(e, wrapper)

	e.Use(middleware.Recover())
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	if err := e.Start(":8080"); err != nil {
		zapLogger.GetZapLogger().Errorf(err.Error())
	}
}

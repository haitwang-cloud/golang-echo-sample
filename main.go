package main

import (
	"github.com/haitwang-cloud/golang-echo-sample/utils/config"
	"github.com/haitwang-cloud/golang-echo-sample/utils/logger"
	"github.com/haitwang-cloud/golang-echo-sample/utils/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
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

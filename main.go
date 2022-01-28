package main

import (
	_ "github.com/haitwang-cloud/golang-echo-sample/docs"
	"github.com/haitwang-cloud/golang-echo-sample/router"
	"github.com/haitwang-cloud/golang-echo-sample/utils/config"
	"github.com/haitwang-cloud/golang-echo-sample/utils/logger"
	"github.com/haitwang-cloud/golang-echo-sample/utils/middlewares"
	"github.com/labstack/echo/v4"
)

// @title           Go-Echo-Sample
// @version         1.0
// @description     This is a sample golang-echo-web server.
// @termsOfService  http://swagger.io/terms/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080

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

	// Routes
	router.InitRoutes(e, wrapper)

	// Start server
	if err := e.Start(":8080"); err != nil {
		zapLogger.GetZapLogger().Errorf(err.Error())
	}
}

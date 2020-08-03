package main

import (
	db "go-logging-proposal/database"
	lg "go-logging-proposal/logger"
	rt "go-logging-proposal/rest"

	"os"
	"path"

	"github.com/jpfuentes2/go-env"
	"github.com/labstack/echo/v4"
)

func init() {
	// Pre load Stuff..
	pwd, _ := os.Getwd()
	env.ReadEnv(path.Join(pwd, ".env"))
}

func main() {
	logger := lg.InitLogger()
	database := db.InitDatabaseConnection(logger)
	echoInstance := echo.New()
	restAPI := &rt.RestApi{
		DB:     database,
		Logger: logger,
		API:    echoInstance,
	}
	restAPI.InitialiseSettings()
	restAPI.InitialiseRouters()
	restAPI.StartServer()
}

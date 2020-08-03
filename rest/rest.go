package rest

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	"go-logging-proposal/accounts"
)

//RestApi - Types for RestAPI
type RestApi struct {
	DB     *sqlx.DB
	Logger *zap.Logger
	API    *echo.Echo
}

//InitialiseSettings - Configures Settings
func (ra *RestApi) InitialiseSettings() {
	ra.API.HideBanner = true
	ra.API.Use(middleware.Gzip())      // Enable Compression
	ra.API.Use(middleware.Recover())   // Recover From Unknown Panics
	ra.API.Use(middleware.RequestID()) // Generate RequestIDs
}

//InitialiseStaticRoutes -
func (ra *RestApi) InitialiseStaticRoutes() {
	ra.Logger.Info("Initialise Static Routes Here")
}

//InitialiseRouters - Configure Routers
func (ra *RestApi) InitialiseRouters() {
	accounts.InitAccountRoutes(ra.API, ra.Logger, ra.DB)
}

//StartServer - BootPoint
func (ra *RestApi) StartServer() {
	// Defer Stack...
	defer ra.Logger.Sync()
	defer ra.DB.Close()
	ra.Logger.Sugar().Info(ra.API.Start(":3001")) // For Development GIN Default Port
}

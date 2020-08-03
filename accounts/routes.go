package accounts

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	accHandler "go-logging-proposal/accounts/handler"
	accRep "go-logging-proposal/accounts/repo"
	accSvc "go-logging-proposal/accounts/service"
)

//InitAccountRoutes - Setup AccountRoutes Handler
func InitAccountRoutes(api *echo.Echo, log *zap.Logger, db *sqlx.DB) {
	accountsRepo := accRep.New(db)
	accountsSvc := accSvc.New(accountsRepo, log)
	accountsHandler := accHandler.New(accountsSvc, log)
	apiGroup := api.Group("accounts")

	apiGroup.GET("/", accountsHandler.GetAccounts)
	apiGroup.GET("/error", accountsHandler.GetAccountsError)
}

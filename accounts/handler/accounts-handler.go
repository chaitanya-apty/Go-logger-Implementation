package handler

import (
	accSvc "go-logging-proposal/accounts/service"
	hp "go-logging-proposal/helpers"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

//AccountsHandler - Handler Layer
type AccountsHandler struct {
	svc *accSvc.AccountsService
	log *zap.Logger
}

//New - EmployeeHandlee
func New(svc *accSvc.AccountsService, log *zap.Logger) *AccountsHandler {
	return &AccountsHandler{
		svc: svc,
		log: log,
	}
}

//GetAccounts - Get All GetAccounts
func (ah *AccountsHandler) GetAccounts(c echo.Context) error {
	requestID := c.Response().Header().Get(echo.HeaderXRequestID)
	names, err := ah.svc.GetRegisteredAccountsService(requestID)
	code, response := hp.GetResponsePayload(names, err, requestID)
	return c.JSON(code, response)
}

//GetAccountsError - Get Error
func (ah *AccountsHandler) GetAccountsError(c echo.Context) error {
	requestID := c.Response().Header().Get(echo.HeaderXRequestID)
	names, err := ah.svc.GetRegisteredErrorService(requestID)
	code, response := hp.GetResponsePayload(names, err, requestID)
	return c.JSON(code, response)
}

package service

import (
	aTypes "go-logging-proposal/accounts/models"
	aRep "go-logging-proposal/accounts/repo"

	"go-logging-proposal/errors"

	"go.uber.org/zap"
)

//AccountsService - Repo Layer
type AccountsService struct {
	repo *aRep.AccountsRepo
	log  *zap.Logger
}

//New - EmployeRepo instance
func New(repo *aRep.AccountsRepo, log *zap.Logger) *AccountsService {
	return &AccountsService{
		repo: repo,
		log:  log,
	}
}

//GetRegisteredAccountsService - Get Registered Accounts Service Layer
func (accSvc *AccountsService) GetRegisteredAccountsService(reqID string) ([]aTypes.Accounts, *errors.Error) {
	const operation errors.Operation = "AccountsService.GetRegisteredAccountsService"

	var details []aTypes.Accounts
	details, err := accSvc.repo.GetRegisteredAccounts(reqID)
	if err != nil {
		accSvc.log.Error("Something went wrong:", zap.String("ReqID", err.RequestID()), zap.String("ERROR", err.Error())) //Maintain Strict types
		return nil, err.WithOperation(operation)                                                                          // Propogate with Current Operation
	}
	return details, nil
}

//GetRegisteredErrorService - Get Error
func (accSvc *AccountsService) GetRegisteredErrorService(reqID string) ([]aTypes.Accounts, *errors.Error) {
	const operation errors.Operation = "AccountsService.GetRegisteredAccountsService"

	var details []aTypes.Accounts
	details, err := accSvc.repo.GetErrorAccountsService(reqID)
	if err != nil {
		accSvc.log.Info("Something went wrong:", zap.String("ReqID", err.RequestID()), zap.String("ERROR", err.Error())) //Maintain Strict types
		return nil, err.WithOperation(operation)                                                                         // Propogate with Current Operation
	}
	return details, nil
}

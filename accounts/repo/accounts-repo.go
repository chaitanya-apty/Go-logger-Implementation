package repo

import (
	aTypes "go-logging-proposal/accounts/models"
	"go-logging-proposal/errors"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap/zapcore"
)

//AccountsRepo - Repo Layer
type AccountsRepo struct {
	Db *sqlx.DB
}

//New - EmployeRepo instance
func New(db *sqlx.DB) *AccountsRepo {
	return &AccountsRepo{
		Db: db,
	}
}

//GetRegisteredAccounts - Get Registered Accounts
func (accRep *AccountsRepo) GetRegisteredAccounts(reqID string) ([]aTypes.Accounts, *errors.Error) {
	const operation errors.Operation = "AccountsRepo.GetRegisteredAccounts"

	var details []aTypes.Accounts
	err := accRep.Db.Select(&details, "SELECT * FROM accounts")
	if err != nil {
		//Generate New Error & Propogate
		return nil, errors.NewError(operation, errors.Unexpected, err.Error(), zapcore.ErrorLevel, reqID)
	}
	return details, nil
}

//GetErrorAccountsService - Get Err Accounts
func (accRep *AccountsRepo) GetErrorAccountsService(reqID string) ([]aTypes.Accounts, *errors.Error) {
	const operation errors.Operation = "AccountsRepo.GetErrorAccountsService"

	var details []aTypes.Accounts
	err := accRep.Db.Select(&details, "SELECT names FROM accounts")
	if err != nil {
		//Generate New Error
		return nil, errors.NewError(operation, errors.Unexpected, err.Error(), zapcore.ErrorLevel, reqID)
	}
	return details, nil
}

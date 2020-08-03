package database

import (
	hp "go-logging-proposal/helpers"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

//InitDatabaseConnection - Initialises DB
func InitDatabaseConnection(zlog *zap.Logger) *sqlx.DB {
	db, err := sqlx.Connect("postgres", hp.GetDatabaseURL())
	if err != nil {
		zlog.Fatal("Cannot Initiate Database Connection %v\n", zap.Error(err), zap.Any("level", zap.ErrorLevel))
	}
	return db
}

package helpers

import "os"

//GetDatabaseURL - Gets Database URL
func GetDatabaseURL() string {
	return os.Getenv("DATABASE_URL")
}

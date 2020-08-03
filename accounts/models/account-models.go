package models

import "gopkg.in/guregu/null.v4"

//Accounts - Models for Accounts
type Accounts struct {
	UserID    string      `json:"id" db:"user_id"`
	Email     string      `json:"email" db:"email"`
	Name      string      `json:"name" db:"username"`
	Created   string      `json:"created" db:"created_on"`
	LastLogin null.String `json:"lastLogin,omitempty" db:"last_login"`
}

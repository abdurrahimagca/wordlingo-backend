// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
)

type Card struct {
	ID          int32
	Language1   string
	Language2   string
	Description string
	DeskID      int32
}

type Desk struct {
	ID          int32
	Title       string
	Description string
	UserID      int32
}

type User struct {
	ID               int32
	Email            string
	Password         string
	Name             string
	VerificationCode sql.NullString
	IsVerified       sql.NullBool
}
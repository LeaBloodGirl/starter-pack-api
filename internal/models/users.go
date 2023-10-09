package models

import "time"

// User model info
// @Description User account information
// @Description with user id, email, login, createdat, updatedat, token, validuntil, appname
type User struct {
	ID         uint      `gorm:"id, primaryKey"`
	Email      string    `gorm:"email"`
	Login      string    `gorm:"login"`
	CreatedAt  time.Time `gorm:"createdat"`
	UpdatedAt  time.Time `gorm:"updatedat"`
	Token      string    `gorm:"token"`
	ValidUntil time.Time `gorm:"validuntil"`
	AppName    string    `gorm:"appname"`
}

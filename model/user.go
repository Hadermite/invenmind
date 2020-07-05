package model

import "time"

// User - A user profile
type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Password     string    `json:"-"`
	RegisterDate time.Time `json:"registerDate"`
}

// AuthToken - A user auth token
type AuthToken struct {
	Token      string    `json:"token"`
	UserID     int       `json:"userId"`
	DeviceName string    `json:"deviceName"`
	CreateDate time.Time `json:"createDate"`
}

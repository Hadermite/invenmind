package model

import "time"

// Location - An inventory location
type Location struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	AddedDate time.Time `json:"addedDate"`
}

// LocationUser - A location member
type LocationUser struct {
	LocationID int       `json:"locationId"`
	UserID     int       `json:"userId"`
	AddedDate  time.Time `json:"addedDate"`
}

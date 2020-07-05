package database

import (
	"github.com/Hadermite/invenmind/model"
	"github.com/jinzhu/gorm"
)

// Connection - The database connection
var Connection *gorm.DB

// Initialize - Initialize the database
func Initialize() {
	var error error
	Connection, error = gorm.Open("sqlite3", "database.sqlite3")
	if error != nil {
		panic("Failed to open database connection!")
	}
	Connection.AutoMigrate(
		&model.User{},
		&model.AuthToken{},
	)
}

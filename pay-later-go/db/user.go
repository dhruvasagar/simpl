package db

import (
	"fmt"

	"github.com/dhruvasagar/pay-later-go/utils"
	"gorm.io/gorm"
)

// User is a gorm database model represent the users data
type User struct {
	gorm.Model
	Name        string
	Email       string
	CreditLimit float64
}

func (u User) String() string {
	return fmt.Sprintf("%s(%.4s)", u.Name, utils.FormatFloat(u.CreditLimit))
}

// CreateUser inserts a user record in the database
func (db *DB) CreateUser(user User) (User, error) {
	result := db.db.Create(&user)
	return user, result.Error
}

// UpdateUser updates a user record in the database
// NOTE: ID must be present on the model, otherwise this inserts a new record
//       instead of updating
func (db *DB) UpdateUser(user User) (User, error) {
	result := db.db.Save(&user)
	return user, result.Error
}

// FindUserByName looks up a user record by name from the database
func (db *DB) FindUserByName(name string) (User, error) {
	var user User
	result := db.db.Where("name = ?", name).First(&user)
	return user, result.Error
}

// Package user provides access to the user table in the MySQL database.
package user

import (
	database "github.com/pcieslar/goforge/core/storage/driver/gorm"

	"github.com/go-sql-driver/mysql"
	"github.com/pcieslar/goforge/model"
	"github.com/pcieslar/goforge/model/userstatus"
)

// User table.
type User struct {
	ID         uint32 `db:"id"`
	FirstName  string `db:"first_name"`
	LastName   string `db:"last_name"`
	Email      string `db:"email"`
	Password   string `db:"password"`
	UserStatus userstatus.UserStatus
	StatusID   uint8          `db:"status_id"`
	CreatedAt  mysql.NullTime `db:"created_at"`
	UpdatedAt  mysql.NullTime `db:"updated_at"`
	DeletedAt  mysql.NullTime `db:"deleted_at"`
}

// TableName for user table.
func (User) TableName() string {
	return "user"
}

// ByEmail gets user information from email.
func ByEmail(email string) (User, error) {
	result := User{}
	return result, model.StandardError(database.SQL.Where("email = ?", email).
		First(&result).Error)
}

// Create creates user.
func Create(firstName, lastName, email, password string) error {
	item := &User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		StatusID:  1,
	}
	return model.StandardError(database.SQL.Create(item).Error)
}

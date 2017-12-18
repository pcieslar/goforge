// Package userstatus provides access to the user_status table in the MySQL database.
package userstatus

import (
	"time"
)

// UserStatus table.
type UserStatus struct {
	ID        uint8     `db:"id"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}

// TableName for user_status table.
func (UserStatus) TableName() string {
	return "user_status"
}

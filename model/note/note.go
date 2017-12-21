// Package note provides access to the note table in the MySQL database.
package note

import (
	"fmt"

	database "github.com/pcieslar/goforge/core/storage/driver/gorm"
	"github.com/pcieslar/goforge/model"

	"github.com/go-sql-driver/mysql"
)

var (
	// table is the table name.
	table = "note"
)

// Note defines the model.
type Note struct {
	ID        uint32         `db:"id"`
	Name      string         `db:"name"`
	UserID    uint32         `db:"user_id"`
	CreatedAt mysql.NullTime `db:"created_at"`
	UpdatedAt mysql.NullTime `db:"updated_at"`
	DeletedAt mysql.NullTime `db:"deleted_at"`
}

// set Note's table name to be `note`
func (Note) TableName() string {
	return table
}

// ByID gets an item by ID.
func ByID(ID string, userID string) (Note, bool, error) {
	result := Note{}
	err := model.StandardError(database.SQL.Where("user_id = ?", userID).
		First(&result).Error)
	return result, err == model.ErrNoResult, err
}

// ByUserID gets all items for a user.
func ByUserID(userID string) ([]Note, bool, error) {
	var result []Note
	err := model.StandardError(database.SQL.Where("user_id = ?", userID).
		Find(&result).Error)
	return result, err == model.ErrNoResult, err
}

// ByUserIDPaginate gets items for a user based on page and max variables.
func ByUserIDPaginate(userID string, max int, page int) ([]Note, bool, error) {
	var result []Note
	err := model.StandardError(database.SQL.Limit(max).Offset(page).Where("user_id = ?", userID).
		Find(&result).Error)
	return result, err == model.ErrNoResult, err
}

// ByUserIDCount counts the number of items for a user.
func ByUserIDCount(userID string) (int, error) {
	var result int
	err := model.StandardError(database.SQL.Model(Note{}).Where("user_id = ?", userID).
		Count(&result).Error)
	return result, err
}

// Create adds an item.
func Create(name string, userID string) error {
	return model.StandardError(database.SQL.Exec(fmt.Sprintf(`
		INSERT INTO %v
		(name, user_id)
		VALUES
		(?,?)
		`, table),
		name, userID).Error)
}

// Update makes changes to an existing item.
func Update(name string, ID string, userID string) error {
	return model.StandardError(database.SQL.Model(Note{}).Where("user_id = ?", userID).
		Update("name", name).Error)
}

// DeleteHard removes an item.
func DeleteHard(ID string, userID string) error {
	return model.StandardError(database.SQL.Unscoped().Where("user_id = ?", userID).
		Where("id = ?", ID).Delete(Note{}).Error)
}

// DeleteSoft marks an item as removed.
func DeleteSoft(ID string, userID string) error {
	return model.StandardError(database.SQL.Where("user_id = ?", userID).
		Where("id = ?", ID).Delete(Note{}).Error)
}

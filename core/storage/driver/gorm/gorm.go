// Package gorm provides a wrapper around the jinzhu/gorm package.
package gorm

import (
	"fmt"
	"strings"
	"sync"

	gm "github.com/pcieslar/goforge/lib/gorm"

	_ "github.com/pcieslar/goforge/lib/gorm/dialects/mysql"
)

// *****************************************************************************
// Thread-Safe Configuration
// *****************************************************************************

var (
	info      Info
	infoMutex sync.RWMutex
)

// Info holds the details for the MySQL connection.
type Info struct {
	Username  string    `json:"Username"`
	Password  string    `json:"Password"`
	Database  string    `json:"Database"`
	Charset   string    `json:"Charset"`
	Collation string    `json:"Collation"`
	Hostname  string    `json:"Hostname"`
	Port      int       `json:"Port"`
	Parameter string    `json:"Parameter"`
	Migration Migration `json:"Migration"`
}

// Migration holds the MySQL migration information.
type Migration struct {
	Table     string
	Folder    string
	Extension string
}

// *****************************************************************************
// Database Handling
// *****************************************************************************

// Connect to the database.
func (c Info) Connect(specificDatabase bool) (*gm.DB, error) {
	var err error

	// Connect to MySQL and ping
	if SQL, err = gm.Open("mysql", c.dsn(specificDatabase)); err != nil {
		return nil, err
	}

	return SQL, err
}

// Disconnect the database connection.
func (c Info) Disconnect() error {
	return SQL.Close()
}

// Create a new database.
func (ci Info) Create() error {
	// Create the database
	err := SQL.Exec(fmt.Sprintf(`CREATE DATABASE %v
				DEFAULT CHARSET = %v
				COLLATE = %v
				;`, ci.Database,
		ci.Charset,
		ci.Collation)).Error
	return err
}

// Drop a database.
func (ci Info) Drop() error {
	// Drop the database
	err := SQL.Exec(fmt.Sprintf(`DROP DATABASE %v;`, ci.Database)).Error
	return err
}

// *****************************************************************************
// MySQL Specific
// *****************************************************************************

var (
	// SQL wrapper
	SQL *gm.DB
)

// DSN returns the Data Source Name.
func (ci Info) dsn(includeDatabase bool) string {
	// Build parameters
	param := ci.Parameter

	// If parameter is specified, add a question mark
	// Don't add one if a question mark is already there
	if len(ci.Parameter) > 0 && !strings.HasPrefix(ci.Parameter, "?") {
		param = "?" + ci.Parameter
	}

	// Add collation
	if !strings.Contains(param, "collation") {
		if len(param) > 0 {
			param += "&collation=" + ci.Collation
		} else {
			param = "?collation=" + ci.Collation
		}
	}

	// Add charset
	if !strings.Contains(param, "charset") {
		if len(param) > 0 {
			param += "&charset=" + ci.Charset
		} else {
			param = "?charset=" + ci.Charset
		}
	}

	// Example: root:password@tcp(localhost:3306)/test
	s := fmt.Sprintf("%v:%v@tcp(%v:%d)/%v", ci.Username, ci.Password, ci.Hostname, ci.Port, param)

	if includeDatabase {
		s = fmt.Sprintf("%v:%v@tcp(%v:%d)/%v%v", ci.Username, ci.Password, ci.Hostname, ci.Port, ci.Database, param)
	}

	return s
}

// setDefaults sets the charset and collation if they are not set.
func (ci Info) setDefaults() Info {
	if len(ci.Charset) == 0 {
		ci.Charset = "utf8"
	}
	if len(ci.Collation) == 0 {
		ci.Collation = "utf8_unicode_ci"
	}

	return ci
}

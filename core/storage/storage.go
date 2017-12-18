// Package storage loads the configuration file with only storage information.
package storage

import (
	"encoding/json"

	"github.com/pcieslar/goforge/core/jsonconfig"
	"github.com/pcieslar/goforge/core/storage/driver/gorm"
	"github.com/pcieslar/goforge/core/storage/driver/mysql"
	"github.com/pcieslar/goforge/core/storage/driver/postgresql"
)

// Info contains the database connection information for the different storage.
type Info struct {
	MySQL      mysql.Info      `json:"MySQL"`
	PostgreSQL postgresql.Info `json:"PostgreSQL"`
	GORM       gorm.Info       `json:"GORM"`
}

// ParseJSON unmarshals bytes to structs.
func (c *Info) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

// LoadConfig reads the configuration file.
func LoadConfig(configFile string) (*Info, error) {
	// Configuration
	config := &Info{}

	// Load the configuration file
	err := jsonconfig.Load(configFile, config)

	// Return the configuration
	return config, err
}

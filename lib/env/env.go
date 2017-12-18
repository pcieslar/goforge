// Package env reads the application settings.
package env

import (
	"encoding/json"

	"github.com/pcieslar/goforge/core/asset"
	"github.com/pcieslar/goforge/core/email"
	"github.com/pcieslar/goforge/core/form"
	"github.com/pcieslar/goforge/core/generate"
	"github.com/pcieslar/goforge/core/jsonconfig"
	"github.com/pcieslar/goforge/core/server"
	"github.com/pcieslar/goforge/core/session"
	"github.com/pcieslar/goforge/core/storage/driver/gorm"
	"github.com/pcieslar/goforge/core/storage/driver/mysql"
	"github.com/pcieslar/goforge/core/view"
)

// *****************************************************************************
// Application Settings
// *****************************************************************************

// Info structures the application settings.
type Info struct {
	Asset      asset.Info    `json:"Asset"`
	Email      email.Info    `json:"Email"`
	Form       form.Info     `json:"Form"`
	Generation generate.Info `json:"Generation"`
	MySQL      mysql.Info    `json:"MySQL"`
	GORM       gorm.Info     `json:"GORM"`
	Server     server.Info   `json:"Server"`
	Session    session.Info  `json:"Session"`
	Template   view.Template `json:"Template"`
	View       view.Info     `json:"View"`
	path       string
}

// Path returns the env.json path
func (c *Info) Path() string {
	return c.path
}

// ParseJSON unmarshals bytes to structs
func (c *Info) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

// New returns a instance of the application settings.
func New(path string) *Info {
	return &Info{
		path: path,
	}
}

// LoadConfig reads the configuration file.
func LoadConfig(configFile string) (*Info, error) {
	// Create a new configuration with the path to the file
	config := New(configFile)

	// Load the configuration file
	err := jsonconfig.Load(configFile, config)

	// Return the configuration
	return config, err
}

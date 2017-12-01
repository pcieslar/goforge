// Package boot handles the initialization of the web components.
package boot

import (
	"log"

	"github.com/pcieslar/goforge/controller"
	"github.com/pcieslar/goforge/lib/env"
	"github.com/pcieslar/goforge/lib/flight"
	"github.com/pcieslar/goforge/viewfunc/link"
	"github.com/pcieslar/goforge/viewfunc/noescape"
	"github.com/pcieslar/goforge/viewfunc/prettytime"
	"github.com/pcieslar/goforge/viewmodify/authlevel"
	"github.com/pcieslar/goforge/viewmodify/flash"
	"github.com/pcieslar/goforge/viewmodify/uri"

	"github.com/pcieslar/goforge/core/form"
	"github.com/pcieslar/goforge/core/pagination"
	"github.com/pcieslar/goforge/core/xsrf"
)

// RegisterServices sets up all the web components.
func RegisterServices(config *env.Info) {
	// Set up the session cookie store
	err := config.Session.SetupConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the MySQL database
	mysqlDB, _ := config.MySQL.Connect(true)

	// Load the controller routes
	controller.LoadRoutes()

	// Set up the views
	config.View.SetTemplates(config.Template.Root, config.Template.Children)

	// Set up the functions for the views
	config.View.SetFuncMaps(
		config.Asset.Map(config.View.BaseURI),
		link.Map(config.View.BaseURI),
		noescape.Map(),
		prettytime.Map(),
		form.Map(),
		pagination.Map(),
	)

	// Set up the variables and modifiers for the views
	config.View.SetModifiers(
		authlevel.Modify,
		uri.Modify,
		xsrf.Token,
		flash.Modify,
	)

	// Store the variables in flight
	flight.StoreConfig(*config)

	// Store the database connection in flight
	flight.StoreDB(mysqlDB)

	// Store the csrf information
	flight.StoreXsrf(xsrf.Info{
		AuthKey: config.Session.CSRFKey,
		Secure:  config.Session.Options.Secure,
	})
}

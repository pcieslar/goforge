// Package controller loads the routes for each of the controllers.
package controller

import (
	"github.com/pcieslar/goforge/controller/about"
	"github.com/pcieslar/goforge/controller/debug"
	"github.com/pcieslar/goforge/controller/home"
	"github.com/pcieslar/goforge/controller/login"
	"github.com/pcieslar/goforge/controller/notepad"
	"github.com/pcieslar/goforge/controller/register"
	"github.com/pcieslar/goforge/controller/static"
	"github.com/pcieslar/goforge/controller/status"
)

// LoadRoutes loads the routes for each of the controllers.
func LoadRoutes() {
	about.Load()
	debug.Load()
	register.Load()
	login.Load()
	home.Load()
	static.Load()
	status.Load()
	notepad.Load()
}

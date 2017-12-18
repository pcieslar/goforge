// Package register handles the user creation.
package register

import (
	"errors"
	"net/http"

	"github.com/pcieslar/goforge/lib/flight"
	"github.com/pcieslar/goforge/middleware/acl"
	"github.com/pcieslar/goforge/model"
	"github.com/pcieslar/goforge/model/user"

	"github.com/pcieslar/goforge/core/form"
	"github.com/pcieslar/goforge/core/passhash"
	"github.com/pcieslar/goforge/core/router"
)

// Load the routes.
func Load() {
	router.Get("/register", Index, acl.DisallowAuth)
	router.Post("/register", Store, acl.DisallowAuth)
}

// Index displays the register page.
func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	v := c.View.New("register/index")
	form.Repopulate(r.Form, v.Vars, "first_name", "last_name", "email")
	v.Render(w, r)
}

// Store handles the registration form submission.
func Store(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	// Validate with required fields
	if !c.FormValid("first_name", "last_name", "email", "password", "password_verify") {
		Index(w, r)
		return
	}

	// Get form values
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	email := r.FormValue("email")

	// Validate passwords
	if r.FormValue("password") != r.FormValue("password_verify") {
		c.FlashError(errors.New("Passwords do not match."))
		Index(w, r)
		return
	}

	// Hash password
	password, errp := passhash.HashString(r.FormValue("password"))

	// If password hashing failed
	if errp != nil {
		c.FlashErrorGeneric(errp)
		http.Redirect(w, r, "/register", http.StatusFound)
		return
	}

	// Get database result
	_, err := user.ByEmail(email)

	if err == model.ErrNoResult { // If success (no user exists with that email)
		err = user.Create(firstName, lastName, email, password)
		// Will only error if there is a problem with the query
		if err != nil {
			c.FlashErrorGeneric(err)
		} else {
			c.FlashSuccess("Account created successfully for: " + email)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
	} else if err != nil { // Catch all other errors
		c.FlashErrorGeneric(err)
	} else { // Else the user already exists
		c.FlashError(errors.New("Account already exists for: " + email))
	}

	// Display the page
	Index(w, r)
}

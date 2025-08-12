package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.IsProd,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// Adds auto-save-and-load of sessions on every request.
func SessionLoad(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}

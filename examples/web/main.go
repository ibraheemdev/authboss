package main

import (
	"encoding/base64"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ibraheemdev/authboss/pkg/authboss"
	"github.com/ibraheemdev/authboss/pkg/authboss/defaults"
	"github.com/ibraheemdev/authboss/pkg/rememberable"

	"github.com/ibraheemdev/authboss-examples/web/users"
)

func main() {
	users.Initialize()
	ab := setupAuthboss()

	r := chi.NewRouter()
	r.Use(ab.LoadClientStateMiddleware, rememberable.Middleware(ab))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	r.Mount("/auth", http.StripPrefix("/auth", ab.Config.Core.Router))
	http.ListenAndServe(":8080", r)
}

func setupAuthboss() *authboss.Authboss {
	ab := authboss.New()

	// Set default config
	defaults.SetCore(&ab.Config, false, false, "/auth", "./templates/authboss", "./templates/authboss/layout.html.tpl")

	// Setup Storage Server
	ab.Config.Storage.Server = users.DB

	if err := ab.Init(); err != nil {
		panic(err)
	}

	// Setup Cookie Store
	var cookieStore defaults.CookieStorer
	cookieStoreKey, _ := base64.StdEncoding.DecodeString(`NpEPi8pEjKVjLGJ6kYCS+VTCzi6BUuDzU0wrwXyf5uDPArtlofn2AG6aTMiPmN3C909rsEWMNqJqhIVPGP3Exg==`)
	cookieStore = defaults.NewCookieStorer(cookieStoreKey, nil)
	ab.Config.Storage.CookieState = cookieStore

	// Setup Session Store
	var sessionStore defaults.SessionStorer
	sessionStoreKey, _ := base64.StdEncoding.DecodeString(`AbfYwmmt8UCwUuhd9qvfNA9UCuN1cVcKJN1ofbiky6xCyyBj20whe40rJa3Su0WOWLWcPpO1taqJdsEI/65+JA==`)
	sessionStore = defaults.NewSessionStorer("authboss_session", sessionStoreKey, nil)
	ab.Config.Storage.SessionState = sessionStore

	return ab
}

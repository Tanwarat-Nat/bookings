package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Tanwarat-Nat/bookings/pkg/config"
	"github.com/Tanwarat-Nat/bookings/pkg/handlers"
	"github.com/Tanwarat-Nat/bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {
	log.Println("Starting the services.")

	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteDefaultMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("main: cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	log.Printf("The Services is start running on port %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: Routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal("main: ListenAndServe:", err)

}

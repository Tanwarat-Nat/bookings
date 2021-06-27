package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"time"

	"github.com/Tanwarat-Nat/bookings/internal/config"
	"github.com/Tanwarat-Nat/bookings/internal/handlers"
	"github.com/Tanwarat-Nat/bookings/internal/models"
	"github.com/Tanwarat-Nat/bookings/internal/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {
	//What am  I going to put in the session
	gob.Register(models.Reservation{})

	log.Println("Starting the services.")

	// change this to true when in production
	app.InProduction = false

	// set up the session
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

	log.Printf("The Services is ready to serve and listen on port %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: Routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("main: ListenAndServe:", err)
	}

}

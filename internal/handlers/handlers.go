package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Tanwarat-Nat/bookings/internal/config"
	"github.com/Tanwarat-Nat/bookings/internal/models"
	"github.com/Tanwarat-Nat/bookings/internal/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Respository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplete(w, r, "home.html", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send data to the template
	render.RenderTemplete(w, r, "about.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

// Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplete(w, r, "generals.html", &models.TemplateData{})
}

// Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplete(w, r, "majors.html", &models.TemplateData{})
}

// Availibility renders the search availibility page
func (m *Repository) Availibility(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplete(w, r, "search-availibility.html", &models.TemplateData{})
}

// PostAvailibility renders the search availibility page
func (m *Repository) PostAvailibility(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("Start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailibilityJSON handlers request for availibility and send JSON response
func (m *Repository) AvailibilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      false,
		Message: "Availibility!",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println("handlers: AvailibilityJSON: cannot send JSON response", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplete(w, r, "contact.html", &models.TemplateData{})
}

// Reservation renders the make a reservation page and display form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplete(w, r, "make-reservation.html", &models.TemplateData{})
}

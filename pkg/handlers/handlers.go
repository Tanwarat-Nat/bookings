package handlers

import (
	"net/http"

	"github.com/Tanwarat-Nat/bookings/pkg/config"
	"github.com/Tanwarat-Nat/bookings/pkg/models"
	"github.com/Tanwarat-Nat/bookings/pkg/render"
)

// Repo the repository used by the handler
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

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplete(w, "home.html", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	render.RenderTemplete(w, "about.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

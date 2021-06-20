package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Tanwarat-Nat/bookings/pkg/config"
	"github.com/Tanwarat-Nat/bookings/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefualtData add  data that available on every page.
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplete renders templates using html/template
func RenderTemplete(w http.ResponseWriter, html string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[html]
	if !ok {
		log.Fatal("render: RenderTemplete: Could not get template from template cache ", ok)
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("render: RenderTemplete: Error writing template to browser", w)
	}

	parsedTemplate, err := template.ParseFiles("./templates/" + html)
	if err != nil {
		log.Println("render: RenderTemplete: Error parsing template:", err)
		return
	}
	parsedTemplate.Execute(w, nil)
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts

	}
	return myCache, nil

}

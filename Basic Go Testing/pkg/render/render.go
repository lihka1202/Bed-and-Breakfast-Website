package render

import (
	"basicgotesting/pkg/config"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// ! Make a universal template cache to store parsed templates
var tc = make(map[string]*template.Template)

// ! Create a config for this
var app *config.AppConfig

func NewTemplateCache(a *config.AppConfig) {
	//! Set the global app pointer here
	app = a
}
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	//! make the variable
	var tc map[string]*template.Template

	//! Check the app config
	if app.UseCache {
		//! Reads from the template
		tc = app.TemplateCache
	} else {
		//! Reads from disk
		tc, _ = CreateTemplate()
	}

	//! Get the requested templates as needed
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(ok)
	}

	buf := new(bytes.Buffer)

	err := t.Execute(buf, nil)
	if err != nil {
		log.Fatal("Something went wrong while we dealt with the buffer")
	}

	//! Render the requested template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateTemplate
func CreateTemplate() (map[string]*template.Template, error) {

	var myCache = map[string]*template.Template{}

	//! I think as of now the file path might be wrong, let us try and see if this works

	pages, globErr := filepath.Glob("./../../templates/*.page.gohtml")
	if globErr != nil {
		return myCache, globErr
	}

	fmt.Println("Entering glob stage")

	//! Pages exists, so we can make use of it
	for _, page := range pages {
		name := filepath.Base(page)

		//! Parse this template and attach some name to it
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		//! Now check if some template files exist
		matches, layoutErr := filepath.Glob("./../../templates/*.layout.gohtml")
		if layoutErr != nil {
			return myCache, layoutErr
		}

		//! Check if there are any matches, ie more than 1
		if len(matches) > 0 {
			//! Need to assign each of these templates to a glob
			ts, err := ts.ParseGlob("./../../templates/*.layout.gohtml")

			if err != nil {
				return myCache, err
			}

			//? this is a little weird, why does this fail to exist outside the scope
			myCache[name] = ts
		}

	}
	return myCache, nil

}

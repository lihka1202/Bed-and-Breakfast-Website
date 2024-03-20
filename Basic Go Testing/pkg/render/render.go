package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	//! Create the template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	//? Get the requested template from the cache

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)
	if err != nil {
		log.Println("Error due to the map", err)
	}

	//! render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	//myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	//! Get everything with *.page.gohtml from ./templates
	pages, err := filepath.Glob("./templates/*.page.gohtml")

	//! Error check
	if err != nil {
		fmt.Println("Something's wrong", err)
		return myCache, err
	}

	//? look through all the pages in the slice
	for _, page := range pages {
		//! Get the name, filepath.Base() just gets the name
		name := filepath.Base(page)

		//! init the template string, and do this by just placing the normal page in , no templates
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		//? look through the layouts in the file system and get them out
		layouts, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return myCache, err
		}

		//! go through layout and see if you have found any matches
		if len(layouts) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts

	}

	return myCache, nil

}

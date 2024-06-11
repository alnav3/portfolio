package main

import (
	"fmt"
	"net/http"
	"path"
	"structure"
	"view"

	"github.com/BurntSushi/toml"
	"github.com/a-h/templ"
)

func main() {
    var webtext structure.WebText
    configpath := path.Join("resources", "webtext_es.toml")
    toml.DecodeFile(configpath, &webtext)
	// Manejador para la ruta "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := view.Base(view.HomePage(webtext))
		templ.Handler(c).ServeHTTP(w, r)
	})

    // Manejador para la ruta "/projects"
    http.HandleFunc("/language", func(w http.ResponseWriter, r *http.Request) {
        language := r.FormValue("language")
        fmt.Println(language)
        if (language == "EN") {
            c := view.ButtonES()
            templ.Handler(c).ServeHTTP(w, r)
        } else {
            c := view.ButtonEN()
            templ.Handler(c).ServeHTTP(w, r)
        }
    })

    http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
        c := view.Navbar(1, webtext.NavItems,view.Projects(webtext.Projects))
        templ.Handler(c).ServeHTTP(w, r)
    })

    // Manejador para la ruta "/experience"
	http.HandleFunc("/experience", func(w http.ResponseWriter, r *http.Request) {
		c := view.Navbar(0,  webtext.NavItems, view.DrawExperience(webtext.Experiences))
		templ.Handler(c).ServeHTTP(w, r)
	})

    // Manejador para la ruta "/experience"
	http.HandleFunc("/homelab", func(w http.ResponseWriter, r *http.Request) {
		c := view.Navbar(2, webtext.NavItems, view.Homelab(webtext.HomelabItems))
		templ.Handler(c).ServeHTTP(w, r)
	})

    handleDirectory("/style/", "/img/", "/js/")

    println("Servidor iniciado en el puerto 80")

	// Inicia el servidor HTTP en el puerto 80
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

func handleDirectory(dirs... string) {
    for _, dir := range dirs {
        http.Handle(dir, http.StripPrefix(dir, http.FileServer(http.Dir("."+dir))))
    }
}


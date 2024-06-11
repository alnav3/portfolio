package main

import (
	"net/http"
	"path"
	"strconv"
	"structure"
	"view"

	"github.com/BurntSushi/toml"
	"github.com/a-h/templ"
)

var webtext structure.WebText
var navbarComponents map[int]templ.Component

func main() {
	// Manejador para la ruta "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := view.Base()
		templ.Handler(c).ServeHTTP(w, r)
	})

    // Manejador para la ruta "/language"
    http.HandleFunc("/language", func(w http.ResponseWriter, r *http.Request) {
        getHomepage(w, r)
    })

    http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		c := view.Navbar(1, webtext.NavItems, navbarComponents[1])
        templ.Handler(c).ServeHTTP(w, r)
    })

    // Manejador para la ruta "/experience"
	http.HandleFunc("/experience", func(w http.ResponseWriter, r *http.Request) {
		c := view.Navbar(0, webtext.NavItems, navbarComponents[0])
		templ.Handler(c).ServeHTTP(w, r)
	})

    // Manejador para la ruta "/experience"
	http.HandleFunc("/homelab", func(w http.ResponseWriter, r *http.Request) {
		c := view.Navbar(2, webtext.NavItems, navbarComponents[2])
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

func getHomepage(w http.ResponseWriter, r *http.Request) {
    // select default language if not supported
    language := r.FormValue("language")
    if (language != "es" && language != "en") {
        language = "en"
    }

    // setup variables & get correct language file
    position, _ := strconv.Atoi(r.FormValue("position"))
    configName := "webtext_" + language + ".toml"
    configpath := path.Join("resources", configName)

    // decode & render
    toml.DecodeFile(configpath, &webtext)
    navbarComponents = mapNavbarComponents()
    c := view.HomePage(webtext, language, position, navbarComponents[position])
    templ.Handler(c).ServeHTTP(w, r)
}

func handleDirectory(dirs... string) {
    for _, dir := range dirs {
        http.Handle(dir, http.StripPrefix(dir, http.FileServer(http.Dir("."+dir))))
    }
}

func mapNavbarComponents() map[int]templ.Component {
    navbarComponents := make(map[int]templ.Component)
    navbarComponents[0] = view.DrawExperience(webtext.Experiences)
    navbarComponents[1] = view.Projects(webtext.Projects)
    navbarComponents[2] = view.Homelab(webtext.HomelabItems)
    return navbarComponents
}

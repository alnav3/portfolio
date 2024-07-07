package main

import (
	"net/http"
	"path"
	"structure"
	"view"

	"github.com/BurntSushi/toml"
	"github.com/a-h/templ"
	"github.com/gorilla/sessions"
)

var webtext structure.WebText
var navbarComponents map[int]templ.Component

// i don't care of this key as i'm not storing data it's not available using javascript
// it would be highly recommended to use a more secure key stored in an environment variable
// if you are storing sensitive data
var store = sessions.NewCookieStore([]byte("secret-key"))

func main() {
	// Manejador para la ruta "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        session, _ := store.Get(r, "preferences")
        session.Values["position"] = 0
        //get theme from cookies
        theme, ok := session.Values["theme"].(string)
        if !ok {
            theme = "dark"
        }

        session.Values["theme"] = theme

        session.Save(r, w)
		c := view.Base(theme)
		templ.Handler(c).ServeHTTP(w, r)
	})

	http.HandleFunc("/theme", func(w http.ResponseWriter, r *http.Request) {
        session, _ := store.Get(r, "preferences")
        //get theme from cookies
        theme, ok := session.Values["theme"].(string)
        if !ok {
            theme = "dark"
        }

        //toggle theme
        if theme == "dark" {
            theme = "light"
        } else {
            theme = "dark"
        }

        session.Values["theme"] = theme

        session.Save(r, w)

        // Respond with the new theme
        w.Header().Set("Content-Type", "text/plain")
        w.Write([]byte(theme))
	})

    // Manejador para la ruta "/language"
    http.HandleFunc("/language", func(w http.ResponseWriter, r *http.Request) {
        getHomepage(w, r)
    })

    // Manejador para la ruta "/experience"
    http.HandleFunc("/experience", func(w http.ResponseWriter, r *http.Request) {
        session, _ := store.Get(r, "preferences")
        session.Values["position"] = 1
        session.Save(r, w)
        c := view.Navbar(1, webtext.NavItems, navbarComponents[1])
        templ.Handler(c).ServeHTTP(w, r)
    })

    http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
        session, _ := store.Get(r, "preferences")
        session.Values["position"] = 0
        session.Save(r, w)
		c := view.Navbar(0, webtext.NavItems, navbarComponents[0])
        templ.Handler(c).ServeHTTP(w, r)
    })

    // Manejador para la ruta "/experience"
	http.HandleFunc("/homelab", func(w http.ResponseWriter, r *http.Request) {
        session, _ := store.Get(r, "preferences")
        session.Values["position"] = 2
        session.Save(r, w)
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
    session, _ := store.Get(r, "preferences")
    //get position from cookies
    position , ok := session.Values["position"].(int)
    if !ok {
        position = 0
        session.Values["position"] = position
    }

    //get language from cookies
    language := r.FormValue("language")
    if (language != "es" && language != "en") {
        language, ok = session.Values["language"].(string)
        if !ok {
            language = "en"
        }
    }

    //get theme from cookies
    theme, ok := session.Values["theme"].(string)
    if !ok {
        theme = "dark"
    }

    session.Values["theme"] = theme
    session.Values["language"] = language
    session.Save(r, w)

    // setup variables & get correct language file
    configName := "webtext_" + language + ".toml"
    configpath := path.Join("resources", configName)

    // decode & render
    toml.DecodeFile(configpath, &webtext)
    navbarComponents = mapNavbarComponents()
    c := view.HomePage(webtext, language, theme, position, navbarComponents[position])
    templ.Handler(c).ServeHTTP(w, r)
}

func handleDirectory(dirs... string) {
    for _, dir := range dirs {
        http.Handle(dir, http.StripPrefix(dir, http.FileServer(http.Dir("."+dir))))
    }
}

func mapNavbarComponents() map[int]templ.Component {
    navbarComponents := make(map[int]templ.Component)
    navbarComponents[0] = view.Projects(webtext.Projects)
    navbarComponents[1] = view.DrawExperience(webtext.Experiences)
    navbarComponents[2] = view.Homelab(webtext.HomelabItems)
    return navbarComponents
}

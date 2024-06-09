package main

import (
	"net/http"
	"view"

	"github.com/a-h/templ"
)

func main() {
	// Manejador para la ruta "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := view.Base(view.HomePage())
		templ.Handler(c).ServeHTTP(w, r)
	})

    // Manejador para la ruta "/projects"
    http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
        c := view.Navbar(1, view.Projects())
        templ.Handler(c).ServeHTTP(w, r)
    })

    // Manejador para la ruta "/experience"
	http.HandleFunc("/experience", func(w http.ResponseWriter, r *http.Request) {
		c := view.Navbar(0, view.Experience())
		templ.Handler(c).ServeHTTP(w, r)
	})

    // Manejador para la ruta "/experience"
	http.HandleFunc("/homelab", func(w http.ResponseWriter, r *http.Request) {
		c := view.Navbar(2, view.Homelab())
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


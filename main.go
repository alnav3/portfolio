package main

import (
	"view"
	"net/http"

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
        c := view.Navbar(0, view.Projects())
        templ.Handler(c).ServeHTTP(w, r)
    })

    // Manejador para la ruta "/experience"
	http.HandleFunc("/experience", func(w http.ResponseWriter, r *http.Request) {
		c := view.Navbar(1, view.Experience())
		templ.Handler(c).ServeHTTP(w, r)
	})

    handleDirectory("/style/", "/img/")

    println("Servidor iniciado en el puerto 42069")

	// Inicia el servidor HTTP en el puerto 42069
	err := http.ListenAndServe(":42069", nil)
	if err != nil {
		panic(err)
	}
}

func handleDirectory(dirs... string) {
    for _, dir := range dirs {
        http.Handle(dir, http.StripPrefix(dir, http.FileServer(http.Dir("."+dir))))
    }
}


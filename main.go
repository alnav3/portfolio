package main

import (
	"front"
	"net/http"
	"github.com/a-h/templ"
)

func main() {
	// Manejador para la ruta "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := front.Base(front.HomePage())
		templ.Handler(c).ServeHTTP(w, r)
	})


    // Manejador para la ruta "/projects"
    http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
        c := front.Navbar(0, front.Projects())
        templ.Handler(c).ServeHTTP(w, r)
    })

    // Manejador para la ruta "/experience"
	http.HandleFunc("/experience", func(w http.ResponseWriter, r *http.Request) {
		c := front.Navbar(1, front.Experience())
		templ.Handler(c).ServeHTTP(w, r)
	})

    fs := http.FileServer(http.Dir("./style"))
    http.Handle("/style/", http.StripPrefix("/style/", fs))

    fs = http.FileServer(http.Dir("./img"))
    http.Handle("/img/", http.StripPrefix("/img/", fs))

    println("Servidor iniciado en el puerto 8420")

	// Inicia el servidor HTTP en el puerto 8420
	err := http.ListenAndServe(":8420", nil)
	if err != nil {
		panic(err)
	}
}


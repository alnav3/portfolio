package main

import (
	"front"
	"io"
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

    // Manejador para la ruta "/about"
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		c := front.Navbar(1, front.About())
		templ.Handler(c).ServeHTTP(w, r)
	})

	// Manejador para la ruta "/style/tailwind.css"
	http.HandleFunc("/style/tailwind.css", func(w http.ResponseWriter, r *http.Request) {
		// Abre el archivo tailwind.css
		cssFile, err := http.Dir("./style").Open("tailwind.css")
		if err != nil {
			http.Error(w, "Archivo no encontrado", http.StatusNotFound)
			return
		}
		defer cssFile.Close()

		// Configura el encabezado de respuesta
		w.Header().Set("Content-Type", "text/css")

		// Copia el contenido del archivo al cuerpo de la respuesta
		_, err = io.Copy(w, cssFile)
		if err != nil {
			http.Error(w, "Error al copiar el archivo", http.StatusInternalServerError)
			return
		}
	})

    println("Servidor iniciado en el puerto 8420")

	// Inicia el servidor HTTP en el puerto 8420
	err := http.ListenAndServe(":8420", nil)
	if err != nil {
		panic(err)
	}
}


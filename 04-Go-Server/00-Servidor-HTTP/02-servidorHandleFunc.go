package servidorHTTP

import (
	"fmt"
	"net/http"
)

/* ServeMux: es una enrutador de peticiones HTTP, dependiendo de
la ruta llama al manejador o handler*/

func ServidorHandleFunc() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hola", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hola desde la ruta /hola!</h1>")
	})
	mux.HandleFunc("/nombre", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hola desde la ruta /nombre!</h1>")
	})

	mux.HandleFunc("/saludo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hola desde la ruta /saludo!</h1>")
	})
	http.ListenAndServe(":8080", mux)
}

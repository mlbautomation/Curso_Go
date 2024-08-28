package servidorHTTP

import (
	"fmt"
	"net/http"
)

/* Hagamos lo mismo que en Servidor2 pero definiendo
los handler */

func hola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola desde la ruta /hola!</h1>")
}

func nombre(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola desde la ruta /nombre!</h1>")
}

func saludo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola desde la ruta /saludo!</h1>")
}

/* Nos damos cuenta que ya mux.HandleFunc() toma esta
funcion que cuenta con la firma correcta y la conviente en un
handler, es decir le asigna un tipo que implementa ya el
metodo ServeHTTP() */

/* Como seria si ya tengo un objeto que implemente el
metodo ServeHTTP(), es decir, que ya sea un manejador en si mismo */

type mensaje struct {
	msg string
}

// Al implementar este método, ya estamos cumpliendo la firma de Handler.
func (m mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.msg)
}

var msg = mensaje{
	msg: "<h1>Hola desde la ruta /mensaje!</h1>",
}

func ServidorHandle() {

	mux := http.NewServeMux()

	// HandleFunc recibe una función y la trabaja como handler
	mux.HandleFunc("/hola", hola)
	mux.HandleFunc("/nombre", nombre)
	mux.HandleFunc("/saludo", saludo)

	// Handle recibe un handler directamente
	mux.Handle("/mensaje", msg)

	http.ListenAndServe(":8080", mux)
}

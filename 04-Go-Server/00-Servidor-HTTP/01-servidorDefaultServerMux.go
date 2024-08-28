package servidorHTTP

import (
	"fmt"
	"net/http"
)

/* Trabajaremos primero con el DefaultServerMux */
/* Escribiendo un texto plano */
/* Escribiendo en html */

/* "HandleFunc" registers "the handler function" for
"the given pattern" in the "DefaultServeMux". */
/* http.HandleFunc(the given pattern, the handler function) */

func ServidorDefaultServerMux() {
	http.HandleFunc("/hola", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hola Mundo!!!\n"))         //texto plano
		fmt.Fprintf(w, "<h1>Hola Mundo!!!</h1>\n") //texto html
		fmt.Fprintf(w, "Request is: %+v\n", r)
	})
	http.ListenAndServe(":8080", nil)
}

/* http.ListenAndServe(":8080", nil)  utiliza un puerto y un handler
colocamos nil porque estamos usando el DefaultServeMux, ¿quiere decir que es un handler?*/

/*
var DefaultServeMux = &defaultServeMux ...

var defaultServeMux ServeMux...

type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry
	es    []muxEntry // slice of entries sorted from longest to shortest.
	hosts bool       // whether any patterns contain hostnames
}...

func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(StatusBadRequest)
		return
	}
	h, _ := mux.Handler(r)
	h.ServeHTTP(w, r)
}...
*/

/* DefaultServeMux implementa el método ServeHTTP (w ResponseWriter, r *Request),
y parra ser un tipo Handler interace hay que cumplir con esa firma:

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}...
*/

/*
HandlerFunc es un handler porque implementa ServeHTTP(...)

type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
*/

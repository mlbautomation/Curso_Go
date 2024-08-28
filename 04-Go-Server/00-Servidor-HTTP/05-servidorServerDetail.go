package servidorHTTP

import (
	"log"
	"net/http"
	"time"
)

/* Creemo ahora nuestra propia estructura Server */
/* Estamos usando las funciones definidad en 04-servidorHandle
recuerda que estamos en el mismo package serrvidor */

func ServidorServerDetail() {

	mux := http.NewServeMux()

	mux.HandleFunc("/hola", hola)
	mux.HandleFunc("/nombre", nombre)
	mux.HandleFunc("/saludo", saludo)

	mux.Handle("/mensaje", msg)

	//http.ListenAndServe(":8080", mux)
	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}

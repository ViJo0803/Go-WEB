package main

import (
	"fmt"
	"log"
	"net/http"
)

//handler

func Hola(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("El metodo es +" + r.Method)
	fmt.Fprintln(rw, "Hola Mundo")
}

func PaginaNF(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)

}

func Error(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "Este es un error", http.StatusConflict)

}

func Saludar(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.URL.RawQuery)
	fmt.Println(r.URL.Query())

	name := r.URL.Query().Get("name")
	edad := r.URL.Query().Get("edad")
	fmt.Fprintf(rw, "Hola, %s tu edad es %s !!", name, edad)
}
func main() {

	//Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hola)
	mux.HandleFunc("/page", PaginaNF)
	mux.HandleFunc("/error", Error)
	mux.HandleFunc("/saludar", Saludar)
	/*
		//Ruta o Router
		http.HandleFunc("/", Hola)
		http.HandleFunc("/page", PaginaNF)
		http.HandleFunc("/error", Error)
		http.HandleFunc("/saludar", Saludar)
	*/
	//crear servidor
	server := &http.Server{
		Addr:    "Localhost:3000",
		Handler: mux,
	}
	fmt.Println("El servidor esta corriendo en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000/")
	//log.Fatal(http.ListenAndServe("localhost:3000", mux))
	log.Fatal(server.ListenAndServe())
}

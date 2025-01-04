package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main10() {
	http.HandleFunc("/", Inicio)
	fmt.Println("Servidor corriendo...")
	http.ListenAndServe(":8080", nil)
	//nil es valor cero para punteros
}

func Inicio(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola este es mi mensaje")
	plantillas.ExecuteTemplate(w, "inicio", nil)
}

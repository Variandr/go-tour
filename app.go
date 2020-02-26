package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", page)

	port := ":9999"
	println("Server port is=", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Listen and Serve", err)
	}

}

func page(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

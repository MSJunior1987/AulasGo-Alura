package main

import (
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)

	err := http.ListenAndServe(":8000", nil) //acessa com o endereco http://localhost/8000 no modo anonimo do chrome
	if err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
}

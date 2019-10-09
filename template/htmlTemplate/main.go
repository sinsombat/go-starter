package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type product struct {
	Name  string
	Price float64
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.ParseFiles("index.html"))
		myProduct := product{Name: "นม", Price: 12.50}
		tpl.ExecuteTemplate(w, "index.html", myProduct)
	})

	http.ListenAndServe(":8080", router)
}

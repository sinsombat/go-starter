package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	userDB := map[string]int{
		"sinsombat": 28,
		"kanjana":   26,
	}

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "sinsombat")
	})
	router.HandleFunc("/user/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		age := userDB[name]
		fmt.Fprintf(w, "Hi My name is %s I'm %d", name, age)
	}).Methods("GET")

	http.ListenAndServe(":8080", router)
}

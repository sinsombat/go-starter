package main

import (
	"fmt"
	"net/http"
)

func main() {
	userDB := map[string]int{
		"sinsombat": 28,
		"kanjana":   26,
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "sinsombat")
	})
	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[len("/user/"):]
		age := userDB[name]
		fmt.Fprintf(w, "Hi My name is %s I'm %d", name, age)
	})
	http.ListenAndServe(":8080", nil)
}

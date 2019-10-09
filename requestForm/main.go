package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type product struct {
	Name  string
	Price float64
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	router.HandleFunc("/login", login)
	http.ListenAndServe(":8080", router)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method", r.Method)
	r.ParseForm()
	fmt.Println("Username: ", r.Form["username"][0])
	fmt.Println("Password: ", r.Form["password"][0])
}

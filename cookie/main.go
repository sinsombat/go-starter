package main

import (
	"fmt"
	"net/http"
	"time"
)

type Cookie struct {
	Name    string
	Value   string
	Expires time.Time
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(time.Minute * 60 * 24 * 365)
	cookie := http.Cookie{Name: "user", Value: "Sinsombat", Expires: expiration}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "Create Cookie")
}

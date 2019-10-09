package main

import (
	"fmt"
)

type Books struct {
	title    string
	author   string
	subtitle string
	price    float64
}

func main() {
	var book1 Books
	book1.title = "Go Book"
	book1.author = "Sinsombat"
	book1.price = 400.25

	fmt.Println(book1)

	book2 := Books{title: "React", author: "Ton", price: 123.50}
	fmt.Println(book2)
}

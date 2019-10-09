package main

import "fmt"

type user struct {
	name string
	id   int
}

func main() {
	p1 := user{
		id:   123,
		name: "sinsombat",
	}
	test(&p1)
	fmt.Println(p1)
}

func test(person *user) {
	person.id = 234
}

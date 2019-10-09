package main

import "fmt"

type human struct {
	name string
	age  int
}

type parrot struct {
	name string
	age  int
}

type talker interface {
	talk()
}

func (h human) talk() {
	if h.age <= 40 {
		fmt.Println("Human - I'm talking.")
	} else {
		fmt.Println("Human - I'm many talking.")
	}
}

func (p parrot) talk() {
	fmt.Println("Parrot  - I'm talking.")
}

func main() {
	talkers := [3]talker{
		human{name: "Somchai", age: 23},
		human{name: "Somsri", age: 41},
		parrot{name: "Dum", age: 2},
	}

	for _, talker := range talkers {
		talker.talk()
	}

	// ผลลัพธ์
	// Human - I'm talking.
	// Parrot  - I'm talking.
}

package main

import (
	"fmt"
)

type human struct {
	speech string
	age    int
}

type chicken struct {
	speech string
	age    int
}

type Talker interface {
	talk()
	getSpeech() string
}

func (h *human) talk() {
	if h.age <= 50 {
		h.speech = "Hello world, i'm a human, i have a dream."
	} else {
		h.speech = "Hello every one i'm human, i happy for today."
	}
}

func (c *chicken) talk() {
	if c.age <= 20 {
		c.speech = "Jeap jeap"
	} else {
		c.speech = "kook kook"
	}
}

func (c *chicken) getSpeech() string {
	return c.speech
}

func (h *human) getSpeech() string {
	return h.speech
}

func main() {
	talkers := getTalkers()
	for _, t := range talkers {
		t.talk()
		fmt.Println(t.getSpeech())
	}
}

func getTalkers() []Talker {
	ton := &human{age: 28}
	iw := &human{age: 51}
	jenny := &chicken{age: 10}
	harry := &chicken{age: 30}

	list := []Talker{ton, iw, jenny, harry}
	return list
}

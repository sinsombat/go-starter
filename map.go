package main

import (
	"fmt"
)

func main() {
	x := make(map[string]string)
	x["TH"] = "Thailand"
	x["ENG"] = "England"
	x["JP"] = "Japan"
	fmt.Println(x)
	y := map[string]string{
		"H": "Hydrogen",
		"O": "Ogxygen",
		"C": "Carbon",
	}
	fmt.Println(y)
	z := map[int]string{
		2: "Hydrogen",
		4: "Ogxygen",
		7: "Carbon",
	}
	fmt.Println(z)
}

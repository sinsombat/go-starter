package main

import (
	"fmt"
)

func main() {
	sum := summation(1, 2, 3, 4, 5, 6)
	fmt.Println(sum)
}
func summation(args ...int) int {
	var total int
	for _, n := range args {
		total += n
	}
	return total
}

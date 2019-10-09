package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	slice2 := append(slice, 7, 8, 9, 10)
	fmt.Println(slice)
	fmt.Println(slice2)
	arr := []float64{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr[1:5])
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3[1:3])
	fmt.Println(slice4)
}

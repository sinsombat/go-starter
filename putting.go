package main

import "fmt"

func main() {
	var platCapacities []float64
	platCapacities = []float64{30, 30, 30, 60, 60, 100}

	var capacity float64 = platCapacities[0] + platCapacities[1] + platCapacities[2] + platCapacities[3] + platCapacities[4] + platCapacities[5]
	var gridLoad = 75.

	utilization := gridLoad / capacity

	fmt.Println("Capacity: ", capacity)
	fmt.Println("Load: ", gridLoad)
	fmt.Println("Utilization: ", utilization)

	fmt.Printf("%-20s%.0f\n", "Capacity: ", capacity)
	fmt.Printf("%-20s%.0f\n", "Load: ", gridLoad)
	fmt.Printf("%-20s%.1f%%\n", "Utilization: ", utilization)
}

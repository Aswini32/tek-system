package main

import "fmt"

func main() {
	numbers := []int{19, 37, 27, 35, 8}
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	fmt.Println("sum of the slices", sum)
	average := sum / len(numbers)

	fmt.Println("average of slice", average)
}

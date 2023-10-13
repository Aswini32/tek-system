package main

import "fmt"

func main() {
	numbers := []int{4, 5, 7, 8, 9, 3}
	fmt.Println(numbers)
	for i := 0; i < len(numbers); i++ {
		numbers[i] = numbers[i] * numbers[i]
	}
	fmt.Println(numbers)
}

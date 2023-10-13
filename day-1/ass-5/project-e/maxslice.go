package main

import "fmt"

func main() {

	numbers := []int{1, 2, 8, 4, 5}
	num := numbers[0]

	for i := 1; i < len(numbers); i++ {
		if numbers[i] > num {
			num = numbers[i]
		}
	}
	fmt.Println(num)
}

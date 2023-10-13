package main

import "fmt"

var numbers = []int{2, 3, 4, 5, 6, 7}
var x int

func searchElement(x int) (a bool, i int) {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == x {
			return true, i
		}
	}
	return false, 0
}

func main() {
	fmt.Scanln(&x)
	a, i := searchElement(x)
	if a {
		fmt.Println("number found")
		fmt.Println("index:", i, "number:", numbers[i])
		// fmt.Println(numbers[i])
	} else {
		fmt.Println("number not found")
	}

}

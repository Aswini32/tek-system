package main

import "fmt"

//var a int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	var (
		num   int
		count = 0
	)
	fmt.Scanln(&num)
	fmt.Println(num)
	for i := 0; i < len(numbers); i++ {
		//a = numbers[i]
		if numbers[i] == num {
			fmt.Println("we found the number")
			return
		}
		count++
	}

	if count != 0 {
		fmt.Println("value not found")
	}

}

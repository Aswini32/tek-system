package main

import "fmt"

func checkEvenOdd(x int) {
	if x%2 == 0 {
		fmt.Println(x, "is given number is even")
		return
	}
	fmt.Println(x, "is the odd number")
}
func main() {
	checkEvenOdd(23)
}

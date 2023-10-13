package main

import "fmt"

func main() {
	even := 0
	odd := 0
	x := []int{1, 2, 3, 4, 5, 6, 10, 8}
	for i := 0; i < len(x); i++ {
		if x[i]%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	fmt.Println("even numbers are", even)
	fmt.Println("odd numbers are", odd)
}

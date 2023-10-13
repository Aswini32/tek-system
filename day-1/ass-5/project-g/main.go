package main

import "fmt"

func main() {
	x := []int{2, 4, 6, 7, 8}
	fmt.Println(x)
	start := 0
	end := len(x) - 1

	for start <= end {
		x[start], x[end] = x[end], x[start]
		start++
		end--
	}

	fmt.Println(x)
}

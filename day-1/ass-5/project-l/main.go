package main

import "fmt"

func main() {
	sum := 0
	x := []int{2, 3, 4, 5, 6, 7, 8, 9, 1}
	for i := 0; i < len(x); i++ {
		if x[i]%2 == 0 {
			sum = sum + x[i]
		}
	}
	fmt.Println(sum)
}

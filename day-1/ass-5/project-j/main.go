package main

import "fmt"

func main() {
	slice1 := []int{2, 3, 4}
	slice2 := []int{5, 6, 7}
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)
}

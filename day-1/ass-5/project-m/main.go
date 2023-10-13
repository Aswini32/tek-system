package main

import "fmt"

var slice1 = []int{2, 4, 5, 6, 7}
var slice2 = []int{2, 4, 5, 6, 7}
var flag bool

func slice() bool {
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func main() {
	if len(slice1) == len(slice2) {
		flag = slice()
		if flag {
			fmt.Println("equal")
		} else {
			fmt.Println("not equal")
		}
	} else {
		fmt.Println("slices has different size")
	}

}

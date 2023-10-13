package main

import "fmt"

func main() {
	s := []string{"grapes", "apple", "pineapple", "mango", "dragon"}
	for _, v := range s {
		fmt.Println(v)
	}
}

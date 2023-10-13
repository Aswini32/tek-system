package main

import "fmt"

func main() {
	fruits := make(map[string]int)
	fruits["apple"] = 10
	fruits["orange"] = 22
	fruits["grapes"] = 15
	fruits["banana"] = 36
	for k, v := range fruits {
		fmt.Println(k, ":", v)
	}
}

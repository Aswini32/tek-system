package main

import (
	"fmt"
)

type values = map[string]any

func main() {
	studentData := make(map[string]values)

	studentData["teju"] = values{
		"age":   23,
		"grade": 4.8,
		"city":  "hyderabad",
	}
	studentData["surya"] = values{
		"age":   32,
		"grade": 8.2,
		"city":  "gadwal",
	}
	studentData["reddy"] = values{
		"age":   23,
		"grade": 4.8,
		"city":  "hyderabad",
	}
	for k, v := range studentData {
		fmt.Println(k, v)
	}

}

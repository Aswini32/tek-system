package main

import "fmt"

func main() {
	studentGrades := make(map[string]float64)
	studentGrades["teju"] = 8.9
	studentGrades["surya"] = 9.4
	studentGrades["harshath"] = 7.9
	for k, v := range studentGrades {
		fmt.Println("studentname:", k, "grades:", v)
	}
	delete(studentGrades, "teju")
	fmt.Println(studentGrades)
}

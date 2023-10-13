package main

import (
	"fmt"
	"project-b/temperature"
)

var (
	a   int
	err error
)

func main() {

	// a= fmt.Scanln()
	//fmt.Scanf("Enter the temperature in Fahrenheit", a)
	_, err = fmt.Scanln(&a)
	if err != nil {
		fmt.Println("give correct values")
		return
	}
	celcius := temperature.Temp(a)
	fmt.Println(celcius)
}

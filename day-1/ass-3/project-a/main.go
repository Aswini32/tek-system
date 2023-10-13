package main

import (
	"fmt"
	"project-a/calculator"
)

func main() {
	a := calculator.Square(21)
	b := calculator.Sum(10, 20)
	c := calculator.Subtraction(20, 13)
	d := calculator.Multiplication(14, 16)
	e, f := calculator.Division(9, 5)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

package main

import (
	"fmt"
	"math/rand"
)

var a int

func main() {
	for {
		a, _ := fmt.Scanln(&a)
		// if err != nil {
		// 	return
		// }
		b := rand.Intn(10)
		if a > b {
			fmt.Println("too high")
		} else if a < b {
			fmt.Println("too low")
		} else {
			fmt.Println("congratulations you guessed the correct number")
		}
	}

}

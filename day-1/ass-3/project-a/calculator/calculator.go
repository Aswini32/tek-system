package calculator

func Square(x int) int {
	return x * x
}
func Sum(x int, y int) int {
	return x + y
}
func Subtraction(x int, y int) int {
	return x - y
}
func Multiplication(x int, y int) int {
	return x * y
}
func Division(x int, y int) (float32, int) {
	var a float32
	a = float32(x / y)
	b := x % y
	// fmt.Println("quotient:", a)
	// fmt.Println("remainder:", b)
	return a, b
}

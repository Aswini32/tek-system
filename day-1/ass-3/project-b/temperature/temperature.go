package temperature

func Temp(x int) float64 {
	y := (x - 32) * 5 / 9
	return float64(y)
}

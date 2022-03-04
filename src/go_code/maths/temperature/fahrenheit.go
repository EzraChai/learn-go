package temperature

func FahrenheitToCelcius(fahrenheit float32) int {
	return int(5 / 9.0 * (fahrenheit - 32))
}

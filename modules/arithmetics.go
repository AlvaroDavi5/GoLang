package arithmetics

func Sum(a float64, b float64) float64 {
	return a + b
}

func Subtract(a float64, b float64) float64 {
	return a - b
}

func Multiply(a float64, b float64) float64 {
	return a * b
}

func Divide(a float64, b float64) float64 {
	checkDivisionByZero(b)

	return a / b
}

func checkDivisionByZero(value float64) {
	if value == 0.0 {
		panic("Division by zero is not allowed")
	}
}

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

func Pow(x float64, y float64) float64 {
	if y == 0 {
		return 1
	}

	if y < 0 {
		x = 1 / x
	}
	y = Abs(y)

	result := 1.0
	for i := 0; i < int(y); i++ {
		result *= x
	}

	return result
}

func Sqrt(value float64) float64 {
	if value < 0 {
		panic("Cannot calculate square root of a negative number")
	}

	if value == 0 {
		return 0
	}

	guess := value / 2.0
	for range 100 {
		guess = (guess + value/guess) / 2.0
	}

	return guess
}

func Abs(value float64) float64 {
	if value < 0 {
		return -value
	}

	return value
}

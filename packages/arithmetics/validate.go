package arithmetics

func checkDivisionByZero(value float64) {
	if value == 0.0 {
		panic("Division by zero is not allowed")
	}
}

package main

import (
	arithmetics "adtech.com/go_monorepo/modules"
)

func main() {
	calculate()
}

func calculate() {
	a := 10.0
	b := 5.0

	sumResult := arithmetics.Sum(a, b)
	println("The sum of", a, "and", b, "is", sumResult)

	subtractResult := arithmetics.Subtract(a, b)
	println("The subtraction of", a, "and", b, "is", subtractResult)

	multiplyResult := arithmetics.Multiply(a, b)
	println("The multiplication of", a, "and", b, "is", multiplyResult)

	divideResult := arithmetics.Divide(a, b)
	println("The division of", a, "and", b, "is", divideResult)

	zero := 0.0
	nearZero := 0.01
	for i := 0; i < 10; i++ {
		nearZero *= nearZero
	}
	println("Now let's try to divide by zero (", zero, ") and a very small number:", nearZero)

	defer func() {
		if r := recover(); r != nil {
			println("Division by zero caused a panic")
		}
	}()
	arithmetics.Divide(a, zero)

	defer func() {
		if r := recover(); r != nil {
			println("Division by a very small number caused a panic")
		}
	}()
	arithmetics.Divide(a, nearZero)
}

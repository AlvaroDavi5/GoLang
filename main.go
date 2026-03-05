package main

import (
	"adtech.com/go_monorepo/modules/arithmetics"
	"adtech.com/go_monorepo/modules/strings"
)

func main() {
	calculate()
}

func calculate() {
	a := 10.0
	b := 5.0

	sumResult := arithmetics.Sum(a, b)
	println("The sum of", strings.FormatFloat(a), "and", strings.FormatFloat(b), "is", strings.FormatFloat(sumResult))

	subtractResult := arithmetics.Subtract(a, b)
	println("The subtraction of", strings.FormatFloat(a), "and", strings.FormatFloat(b), "is", strings.FormatFloat(subtractResult))

	multiplyResult := arithmetics.Multiply(a, b)
	println("The multiplication of", strings.FormatFloat(a), "and", strings.FormatFloat(b), "is", strings.FormatFloat(multiplyResult))

	divideResult := arithmetics.Divide(a, b)
	println("The division of", strings.FormatFloat(a), "and", strings.FormatFloat(b), "is", strings.FormatFloat(divideResult))

	zero := 0.0
	nearZero := 0.01
	for i := 0; i < 10; i++ {
		nearZero *= nearZero
	}
	println("Now let's try to divide by zero (", strings.FormatFloat(zero), ") and a near-zero number:", strings.FormatFloat(nearZero))

	defer func() {
		if r := recover(); r != nil {
			println("Division by zero caused a panic")
		}
	}()
	arithmetics.Divide(a, zero)

	defer func() {
		if r := recover(); r != nil {
			println("Division by a near-zero number caused a panic")
		}
	}()
	arithmetics.Divide(a, nearZero)
}

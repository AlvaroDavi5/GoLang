package main

import (
	arith "adtech.com/go_monorepo/modules/arithmetics"
	str "adtech.com/go_monorepo/modules/strings"
	"rsc.io/quote"
)

func main() {
	helloWorld := quote.Hello()
	println(helloWorld, "\n")

	calculate(10.0, 5.0)
}

func calculate(a float64, b float64) {
	sumResult := arith.Sum(a, b)
	println("The sum of", str.FormatFloat(a), "and", str.FormatFloat(b), "is", str.FormatFloat(sumResult))

	subtractResult := arith.Subtract(a, b)
	println("The subtraction of", str.FormatFloat(a), "and", str.FormatFloat(b), "is", str.FormatFloat(subtractResult))

	multiplyResult := arith.Multiply(a, b)
	println("The multiplication of", str.FormatFloat(a), "and", str.FormatFloat(b), "is", str.FormatFloat(multiplyResult))

	divideResult := arith.Divide(a, b)
	println("The division of", str.FormatFloat(a), "and", str.FormatFloat(b), "is", str.FormatFloat(divideResult))

	zero := 0.0
	nearZero := 0.01
	for i := 0; i < 10; i++ {
		nearZero *= nearZero
	}
	println("Now let's try to divide by zero (", str.FormatFloat(zero), ") and a near-zero number:", str.FormatFloat(nearZero))

	defer func() {
		if r := recover(); r != nil {
			println("Division by zero caused a panic")
		}
	}()
	arith.Divide(a, zero)

	defer func() {
		if r := recover(); r != nil {
			println("Division by a near-zero number caused a panic")
		}
	}()
	arith.Divide(a, nearZero)
}

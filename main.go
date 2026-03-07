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

	println("Let's try to divide by a near-zero number")
	const nearZero = 0.000000000000000000000001
	divideResult = arith.Divide(a, nearZero)
	println("The division of", str.FormatFloat(a), "and ", str.FormatFloat(nearZero), " is", str.FormatFloat(divideResult))

	println("Let's try to divide by zero")
	const zero = 0.00
	defer func() {
		if r := recover(); r != nil {
			println("Division by", str.FormatFloat(zero), "caused a panic")
		}
	}()
	divideResult = arith.Divide(a, zero)
	println("The division of", str.FormatFloat(a), "and ", str.FormatFloat(zero), " is", str.FormatFloat(divideResult))
}

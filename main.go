package main

import (
	arith "adtech.com/go_monorepo/packages/arithmetics"
	str "adtech.com/go_monorepo/packages/strings"
)

func main() {
	helloWorld := str.GetHelloWorld()
	println(helloWorld, "\n")

	calculate(10.0, 5.0)

	println("\nDONE\n")
}

func calculate(a float64, b float64, c ...float64) {
	y := 2.0
	if len(c) > 0 {
		y = c[0]
	}

	powResult := arith.Pow(b, y)
	println(str.FormatFloat(b), "pow", str.FormatFloat(y), "is", str.FormatFloat(powResult))

	println("The square-root of", str.FormatFloat(powResult), "is", str.FormatFloat(arith.Sqrt(powResult)))

	println("The sum of", str.FormatFloat(a), "and", str.FormatFloat(b), "is", str.FormatFloat(arith.Sum(a, b)))

	println("The subtraction of", str.FormatFloat(a), "and", str.FormatFloat(b), "is", str.FormatFloat(arith.Subtract(a, b)))

	println("The multiplication of", str.FormatFloat(a), "and", str.FormatFloat(b), "is", str.FormatFloat(arith.Multiply(a, b)))

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

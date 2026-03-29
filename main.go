package main

import (
	log "log"

	concurrency "adtech.com/concurrency"
	arith "adtech.com/go_monorepo/packages/arithmetics"
	str "adtech.com/go_monorepo/packages/strings"
)

func main() {
	helloWorld := str.GetHelloWorld()
	println(helloWorld)

	startConcurrency()
	calculate(10.0, 5.0)

	println("\nDONE\n")
}

func startConcurrency() {
	println("\nStarting concurrency module...")
	concurrency.Start()
}

func calculate(a float64, b float64, c ...float64) {
	println("\nCalculating...")

	y := 2.0
	if len(c) > 0 {
		y = c[0]
	}

	powResult := arith.Pow(b, y)
	println(str.ParseFloat(b), "pow", str.ParseFloat(y), "is", str.ParseFloat(powResult))

	sqrtResult, err := arith.Sqrt(powResult)
	if err != nil {
		log.Fatalf("Error calculating square root: %s", err.Error())
	}
	println("The square-root of", str.ParseFloat(powResult), "is", str.ParseFloat(sqrtResult))

	println("The sum of", str.ParseFloat(a), "and", str.ParseFloat(b), "is", str.ParseFloat(arith.Sum(a, b)))

	println("The subtraction of", str.ParseFloat(a), "and", str.ParseFloat(b), "is", str.ParseFloat(arith.Subtract(a, b)))

	println("The multiplication of", str.ParseFloat(a), "and", str.ParseFloat(b), "is", str.ParseFloat(arith.Multiply(a, b)))

	divideResult := arith.Divide(a, b)
	println("The division of", str.ParseFloat(a), "by", str.ParseFloat(b), "is", str.ParseFloat(divideResult))

	println("Let's try to divide by a near-zero number")
	const nearZero = 0.000000000000000000000001
	divideResult = arith.Divide(a, nearZero)
	println("The division of", str.ParseFloat(a), "by", str.ParseFloat(nearZero), "is", str.ParseFloat(divideResult))

	println("Let's try to divide by zero")
	const zero = 0.00
	defer func() {
		if r := recover(); r != nil {
			println("Division by", str.ParseFloat(zero), "caused a panic:", r.(error).Error())
		}
	}()
	divideResult = arith.Divide(a, zero)
	println("The division of", str.ParseFloat(a), "by", str.ParseFloat(zero), "is", str.ParseFloat(divideResult))
}

package arithmetics

import (
	"errors"
)

func validateDivision(value float64) error {
	if value == 0.0 {
		return errors.New("Division by zero is not allowed")
	}
	return nil
}

func validateSqrt(value float64) error {
	if value < 0 {
		return errors.New("Cannot calculate square root of a negative number")
	}
	return nil
}

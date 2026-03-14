package strings

import (
	"fmt"
	"strconv"
	"strings"
)

func parseNumberToSring[T number](value T) string {
	return fmt.Sprintf("%v", value)
}

func parseIntegerToSring(value any) string {
	return fmt.Sprintf("%d", value)
}

func parseFloatToSring(value float64) string {
	strValue := strconv.FormatFloat(value, 'f', -1, 64)
	parts := strings.Split(strValue, ".")

	floatWithoutDecimal := strconv.FormatFloat(value, 'f', 1, 64)

	if len(parts) < 2 {
		return floatWithoutDecimal
	}

	floatPart := parts[1]
	allZeros := true
	for _, digit := range floatPart {
		if digit != '0' {
			allZeros = false
			break
		}
	}

	if allZeros {
		return floatWithoutDecimal
	}

	precision := len(floatPart)
	return strconv.FormatFloat(value, 'f', precision, 64)
}

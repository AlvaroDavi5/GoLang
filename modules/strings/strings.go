package strings

import (
	"strconv"
	"strings"
)

func FormatFloat(value float64) string {
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

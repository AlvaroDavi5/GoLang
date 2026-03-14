package strings

import (
	"rsc.io/quote"
)

func GetHelloWorld() string {
	helloWorld := quote.Hello()
	return helloWorld
}

func ParseNumber[T number](value T) string {
	switch v := any(value).(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return parseIntegerToSring(v)
	case float32:
		return parseFloatToSring(float64(v))
	case float64:
		return parseFloatToSring(v)
	default:
		return parseNumberToSring(value)
	}
}

func ParseInteger[T integer](value T) string {
	return parseIntegerToSring(value)
}

func ParseFloat[T float](value T) string {
	return parseFloatToSring(float64(value))
}

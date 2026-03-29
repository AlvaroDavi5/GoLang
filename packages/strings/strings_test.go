package strings

import (
	"testing"
)

func TestParseFloat(t *testing.T) {
	testParseFloatForZero(t)
	testParseFloatForNearZero(t)
}

func TestGetHelloWorld(t *testing.T) {
	const expected = "Hello, world."
	result := GetHelloWorld()

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func testParseFloatForZero(t *testing.T) {
	const zero = 0.0000000000
	const expected = "0.0"
	result := ParseFloat(zero)

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func testParseFloatForNearZero(t *testing.T) {
	const nearZero = 0.000000000000000000000001
	const expected = "0.000000000000000000000001"
	result := ParseFloat(nearZero)

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

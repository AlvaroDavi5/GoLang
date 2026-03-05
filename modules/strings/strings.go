package strings

import "strconv"

func FormatFloat(value float64) string {
	precision := 2
	if value < 1.0 {
		precision = 4
	}

	return strconv.FormatFloat(value, 'f', precision, 64)
}

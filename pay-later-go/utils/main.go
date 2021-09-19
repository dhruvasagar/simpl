package utils

import "strconv"

// FormatFloat formats a float value to a string representation showing smallest
// to show all necessary decimal digits
func FormatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

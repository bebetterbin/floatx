package floatx

import "math"

// InDelta returns ture if the abs difference value of expected and actual is in delta.
func InDelta(expected, actual, delta float64) bool {
	return math.Abs(expected-actual) <= delta
}

// CeilFloat returns the least value greater than or equal to x, keep decimal decimal digit.
func CeilFloat(x float64, decimal int) float64 {
	d := float64(1)
	if decimal > 0 {
		d = math.Pow10(decimal)
	}
	return math.Ceil(x*d) / d
}

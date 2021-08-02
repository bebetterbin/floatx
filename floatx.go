package floatx

import "math"

// InDelta return ture if the abs difference value of expected and actual is in delta.
func InDelta(expected, actual, delta float64) bool {
	return math.Abs(expected-actual) <= delta
}

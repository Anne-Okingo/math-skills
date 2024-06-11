package functions

import "math"

func StandardDev(number []float64) float64 {
	variance := Variance(number)

	return math.Sqrt(variance)
}

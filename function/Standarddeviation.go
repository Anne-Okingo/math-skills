package functions

import "math"

func StandardDev(nb []float64) float64 {
	variance := Variance(nb)

	return math.Sqrt(variance)
}

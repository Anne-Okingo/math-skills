package functions

func Variance(nb []float64) float64 {
	sum := 0.0

	mean := Average(nb)

	for _, char := range nb {
		sum += (char - mean) * (char - mean)
	}
	return sum / float64(len(nb))

}

package functions

func Variance(number []float64) float64 {
	sum := 0.0

	mean := Average(number)

	for _, char := range number {
		sum += (char - mean) * (char - mean)
	}
	return sum / float64(len(number))

}

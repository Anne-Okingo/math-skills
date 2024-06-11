package functions

func Average(num []float64) float64 {
	total := 0.0
	for _, ch := range num {
		total = total + float64(ch)
	}
	return (total / float64(len(num)))
}

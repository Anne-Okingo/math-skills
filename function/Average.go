package functions

func Average(nb []float64) float64 {
	total := 0.0
	for _, ch := range nb {
		total = total + float64(ch)
	}
	return (total / float64(len(nb)))
}

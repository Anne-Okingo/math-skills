package functions

import "sort"

func Median(number []float64) float64 {
	sort.Float64s(number)

	length := len(number)
	if length%2 == 0 {
		return (number[length/2-1] + number[length/2]) / 2
	} else {
		return number[length/2]
	}
}

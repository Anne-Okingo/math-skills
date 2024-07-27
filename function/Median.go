package functions

import "sort"

func Median(num []float64) float64 {
	sort.Float64s(num)

	length := len(num)
	if length%2 == 0 {
		return (num[length/2-1] + num[length/2]) / 2
	} else {
		return num[length/2]
	}
}

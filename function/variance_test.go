package functions

import (
	"math"
	"reflect"
	"testing"
)

func TestVariance(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := float64(7)

	results := Variance(input)

	if !reflect.DeepEqual(math.Round(results), expected) {
		t.Errorf("Test Variance failed Expected : %v Got : %v", expected, results)
	}
}

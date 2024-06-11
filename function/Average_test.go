package functions

import (
	"math"
	"reflect"
	"testing"
)

func TestAverage(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := float64(5)

	result := Average(input)

	if !reflect.DeepEqual(math.Round(result), expected) {
		t.Errorf("Test Average Failed Expected : %v, Got %v ", expected, result)
	}
}


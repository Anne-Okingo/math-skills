package functions

import (
	"math"
	"reflect"
	"testing"
)

func TestStandardDev(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := float64(3)

	results := StandardDev(input)

	if !reflect.DeepEqual(math.Round(results), expected) {
		t.Errorf("Test StandardDev Failed Expected : %v, Got : %v ", expected, results)
	}
}

package functions

import (
	"math"
	"reflect"
	"testing"
)

func TestMedian(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := float64(5)
	results := Median(input)

	if !reflect.DeepEqual((math.Round(results)), expected) {
		t.Errorf("Test Medial Failed Expected : %v. Got : %v .", expected, results)
	}
}
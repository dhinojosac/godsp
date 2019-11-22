package outliers

import (
	"errors"
	"sort"
)

func Median(input []float64) (float64, error) {

	if len(input) == 0 {
		return 0, errors.New("Input is empty")
	}

	sort.Float64s(input)

	half := len(input) / 2

	if len(input)%2 == 0 {
		return input[half], nil
	}

}

func Mean(input []float64) (float64, error) {
	sum := 0.0

	if len(input) == 0 {
		return 0, errors.New("Input is empty")
	}

	for _, val := range input {
		sum += val
	}

	return sum / float64(len(input)), nil

}

/*
func Filloutliers(input []float64) ([]float64, error) {

}
*/

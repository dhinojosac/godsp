package outliers

import (
	"errors"
	"math"
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
	return input[half], nil
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

func Filloutliers(input []float64) ([]float64, error) {
	if len(input) == 0 {
		return nil, errors.New("Input is empty")
	}

	out, err := LocateOutliers(input)
	if err != nil {
		return nil, err
	}

	return out, nil

}

func LocateOutliers(input []float64) ([]float64, error) {
	if len(input) == 0 {
		return nil, errors.New("Input is empty")
	}

	p := 3.0 //fixed!

	madfactor := -1 / (math.Sqrt(2) * (-0.4769)) //~1.4826

	center, err := Median(input)
	if err != nil {
		return nil, err
	}

	t1 := make([]float64, len(input))
	for i := range t1 {
		t1[i] = math.Abs(input[i] - center)
	}

	med, err := Median(t1)
	if err != nil {
		return nil, err
	}

	amad := madfactor * med

	lowerbound := center - p*amad
	upperbound := center + p*amad

	tf := make([]bool, len(input))
	for i := range tf {
		tf[i] = (input[i] > upperbound) || (input[i] < lowerbound)
	}

	b := make([]float64, len(input))
	for i := range b {
		if tf[i] == false {
			b[i] = input[i]
		} else {
			b[i] = center //Replace outlier by center
		}
	}

	return b, nil
}

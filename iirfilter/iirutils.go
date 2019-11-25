package iirfilter

import (
	"errors"
)

func CramerSolve(a, b, c, d, e, f float64) ([2]float64, error) {
	x := [2]float64{-1, -1}
	det := a*d - b*c
	if det == 0 {
		return x, errors.New("Determinant is zero")
	}
	x[0] = (e*d - b*f) / det
	x[1] = (e*d - b*f) / det

	//log.Printf("Solve\nx1:%v\nx2:%v\n", x[0], x[1])
	return x, nil

}

// Adds data at the beginning and at the end
func AddPadding(input []float64, N int) ([]float64, error) {
	if len(input) == 0 {
		return nil, errors.New("No input")
	}

	arr_start := make([]float64, N)
	arr_end := make([]float64, N)

	//Get first data to prepend to start of the array
	for i := N; i >= 1; i-- {
		arr_start[N-i] = input[0] - input[i] + input[0]
	}

	end := len(input) - 1 //end index

	//Get end data to append to end of the array
	for i := 0; i < N; i++ {
		arr_end[i] = input[end] - input[end-i-1] + input[end]
	}

	// Append original array to start array
	input = append(arr_start, input...)

	// Append end array to original array
	input = append(input, arr_end...)

	return input, nil

}

// Delete data added
func RemovePadding(input []float64, N int) ([]float64, error) {
	var output []float64
	for i, _ := range input {
		if i >= len(input)-N {
			break
		}
		if i >= N {
			output = append(output, input[i])
		}

	}
	return output, nil
}

func GetCol(a [][]float64, row int) []float64 {
	b := make([]float64, len(a))

	for i := range a {
		b[i] = a[i][row]
	}

	return b
}

func GetRow(a [][]float64, col int) []float64 {
	b := make([]float64, len(a[0]))
	for i := range a[0] {
		b[i] = a[col][i]
	}

	return b
}

func ScaleMatrix(a []float64, g float64) []float64 {
	for i := range a {
		a[i] *= g
	}
	return a
}

func ReverseMatrix(a []float64) []float64 {
	b := make([]float64, len(a))
	v := len(a) - 1
	for i := range a {
		b[i] = a[v]
		v--
	}
	return b
}

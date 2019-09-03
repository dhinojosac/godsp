package godsp

import (
	"errors"
	"fmt"
)

func FIR_filter(coef []float64, input []float64) ([]float64, error) {
	output := make([]float64, len(input)+len(coef))
	fmt.Printf("len:%v\n", len(output))

	for i, _ := range output {
		output[i] = 0
	}

	for i, ival := range input {
		for j, bval := range coef {
			output[i+j] = output[i+j] + ival*bval
		}

	}
	if len(coef) == 0 {
		return nil, errors.New("Test error")
	}
	return output, nil
}

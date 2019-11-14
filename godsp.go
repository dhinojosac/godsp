package godsp

import (
	"errors"
	"fmt"
	"log"
)

// Adds data at the beginning and at the end
func PreapareArray(input []float64, N int) ([]float64, error) {
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
func PostProcess(input []float64, N int) ([]float64, error) {
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

/*
 * IIR Filter
 * Using a and b coefficients, but this filter is nuemericaly very unstable
 */
func IIR_filter(numCoef, denCoef, input []float64) ([]float64, error) {

	//fmt.Printf("Len in:%v\nLen out:%v\n", len(input), len(output))
	fmt.Printf("Len B coef:%v\nLen A coef:%v\n", len(numCoef), len(denCoef))

	if len(numCoef) == 0 && len(denCoef) == 0 {
		return nil, errors.New("Test error")
	}

	fmt.Printf("Len input pre:%v\n", len(input))

	N := 50
	var err error
	input, err = PreapareArray(input, N) //add data to input
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Len input post:%v\n", len(input))

	output := make([]float64, len(input))

	n := len(denCoef)
	z := make([]float64, len(denCoef))

	// Fill z with zeros
	for i, _ := range z {
		z[i] = 0
	}

	// Fill output with zeros
	for i, _ := range output {
		output[i] = 0
	}

	for m, _ := range input {
		Xm := input[m]
		Ym := numCoef[0]*Xm + z[0]

		for i := 1; i < n; i++ {
			z[i-1] = numCoef[i]*Xm + z[i] - denCoef[i]*Ym
		}
		output[m] = Ym
	}

	return output, nil
}

/*
 * IIR FILTER
 * Using (ZPK) zero, poles and gain design.
 */
func IIR_filter_SOS(input []float64) ([]float64, error) {

	fmt.Printf("Len input pre:%v\n", len(input))

	N := 50
	var err error
	input, err = PreapareArray(input, N) //add data to input
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Len input post:%v\n", len(input))

	output := make([]float64, len(input))
	Ns := 8

	var state [8][2]float64

	/*
		for i, _ := range state {
			state[i][0] = 0
			state[i][1] = 0
		}
	*/

	// Fill output with zeros
	for i, _ := range output {
		output[i] = 0
	}
	/*
		//TEST
		fmt.Printf("B_Coef[0][0]: %v\n", B_Coef[0][0])
		fmt.Printf("B_Coef[0][1]: %v\n", B_Coef[0][1])

		fmt.Printf("B_Coef[1][0]: %v\n", B_Coef[1][0])
		fmt.Printf("B_Coef[1][1]: %v\n", B_Coef[1][1])

		fmt.Printf("B_Coef[1][0]: %v\n", B_Coef[2][0])
		fmt.Printf("B_Coef[1][1]: %v\n", B_Coef[2][1])
	*/

	for m, _ := range input {
		Xm := input[m] * G
		Ym := 0.0

		for i := 0; i < Ns; i++ {
			//fmt.Printf("i: %v\n", i)
			Ym = Xm + state[i][0]
			state[i][0] = B_Coef[i][0]*Xm - A_Coef[i][0]*Ym + state[i][1]
			state[i][1] = B_Coef[i][1]*Xm - A_Coef[i][1]*Ym
			Xm = Ym
		}

		output[m] = Ym
	}

	output, err = PostProcess(output, N)
	if err != nil {
		log.Fatal(err)
	}

	return output, nil
}

/*
func IIR_filter(numCoef, denCoef, input []float64) ([]float64, error) {
	output := make([]float64, len(input)+len(numCoef))
	fmt.Printf("Len in:%v\nLen out:%v\n", len(input), len(output))
	fmt.Printf("Len B coef:%v\nLen A coef:%v\n", len(numCoef), len(denCoef))

	for i, _ := range output {
		output[i] = 0
	}

	for i, _ := range input {

		if i >= len(numCoef) {
			bcc := 0.0
			acc := 0.0

			// sum of inputs and b coefficients
			for b := 0; b < len(numCoef); b++ {
				bcc += input[i-b] * numCoef[b]
				//fmt.Printf("X[%v]*b[%v]\n", i-b, b)
			}

			// sum of outputs and a coefficients
			for a := 1; a < len(denCoef); a++ {
				acc += output[i-a] * denCoef[a]
				//fmt.Printf("Y[%v]*a[%v]\n", i-a, a)
			}
			fmt.Printf("i:%v  bcc:%v   acc:%v\n", i, bcc, acc)

			// output y[n]
			output[i] = bcc - acc

		}

		//fmt.Printf("out[%d]: %v\n", i, output[i])
		//fmt.Printf("%v\n", output[i])
	}

	if len(numCoef) == 0 && len(denCoef) == 0 {
		return nil, errors.New("Test error")
	}

	return output, nil
}
*/

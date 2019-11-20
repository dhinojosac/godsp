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
func IIR_filter_AB(numCoef, denCoef, input []float64) ([]float64, error) {

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
func IIR_filter_SOS(f SosFilter, input []float64) ([]float64, error) {

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

	// Fill output with zeros
	for i, _ := range output {
		output[i] = 0
	}

	for m, _ := range input {
		Xm := input[m] * f.gain
		Ym := 0.0

		for i := 0; i < Ns; i++ {
			Ym = Xm + state[i][0]
			state[i][0] = f.bcoeff[i][0]*Xm - f.acoeff[i][0]*Ym + state[i][1]
			state[i][1] = f.bcoeff[i][1]*Xm - f.acoeff[i][1]*Ym
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
 * IIR FILTER
 * Using (ZPK) zero, poles and gain design.
 */
func IIR_filter_SOS_2(f SosFilter, input []float64) ([]float64, error) {

	fmt.Printf("Len input pre:%v\n", len(input))

	N := 50
	var err error
	input, err = PreapareArray(input, N) //add data to input
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Len input post:%v\n", len(input))

	output := make([]float64, len(input))
	Ns := len(f.acoeff)
	fmt.Printf("Ns:%v\n", len(f.acoeff))

	var state [8][2]float64

	// Fill output with zeros
	for i, _ := range output {
		output[i] = 0
	}

	for m, _ := range input {
		Xm := input[m] * f.gain
		Ym := 0.0

		for i := 0; i < Ns; i++ {
			Ym = Xm + state[i][0]
			state[i][0] = f.bcoeff[i][0]*Xm - f.acoeff[i][0]*Ym + state[i][1]
			state[i][1] = f.bcoeff[i][1]*Xm - f.acoeff[i][1]*Ym
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

// new implementation

//Not implemented
func MyFiltfilt(sos [][]float64, g float64, ord int, input []float64)(output []float64, err error){
	npts = len(input)
	b,a,zi,nfact,L, err := GetCoeffsAndInitialConditions(sos,g ,ord, npts)
	if err!=nil{
		return nil, err
	}


	output, err = FfOneChanCat(b, a, input, zi, nfact, L)
	if err!=nil{
		return nil, err
	}

	return output, nil

}

//Not implemented
func GetCoeffsAndInitialConditions(sos [][]float64, g float64, ord int, npts int) ([][]float64,[][]float64, [][]float64, int, int, error){
	L := len(sos)        //down
	ncols := len(sos[0]) //right

	for i := 0; i < 3; i++ {
		sos[0][i] *= g
	}

	var a [][]float64
	var b [][]float64

	for i := 0; i < 3; i++ {
		for j := 0; j < L; j++ {
			b[i][j] = sos[j][i]
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < L; j++ {
			a[i][j] = sos[j][i+3]
		}
	}

	nfact := 3 * ord

	if npts <= nfact {
		return "ERROR input too short"
	}

	var zi [2][L]float64
	var rhs [2][1]float64
	var aA [2][2]float64
	for i:=0; i<L; i++{
		rhs  = (b(2:3,i) - b(1,ii)*a(2:3,i)); //fila 2 y 3, columna i  (2x1)
    	zi(:,i) = ( eye(2) - [-a(2:3,i),[1;0]] ) \ rhs; // eye matriz identidad,  concatenar vectores 2x1 para armar una matriz 2x2
	}

	return b,a,zi,nfact,L, nil

}

func GetCol(a [][]float64, row int) []float64{
	b := make([]float64, len(a))
	
	for i:= range a{
		b[i] = a[i][row]
	}

	return b
}

func GetRow(a [][]float64, col int) []float64{
	b := make([]float64, len(a[0]))
	for i:= range a[0]{
		b[i] = a[col][i]
	}

	return b
}

func ScaleMatrix(a []float64, g float64)[]float64{
	for i:= range a{
		a[i]*=g
	}
	return a
}

func ReverseMatrix(a []float64) []float64{
	b := make([]float64, len(a))
	v := len(a)-1
	for i:= range a{
		b[i] = a[v]
		v--
	}
	return b
}

//Implemented, but not tested
func FfOneChanCat(b, a [][]float64, input []float64, zi [][]float64, nfact, l int) ([]float64, error) {
	if len(input) == 0 {
		return nil, errors.New("No input or zero")
	}
	for i := 0; i < l; i++ {
		//Padding
		var err error
		input, err = PreapareArray(input, nfact) //add data to input
		if err != nil {
			log.Fatal(err)
		}

		input = MyFilter(GetCol(b), GetCol(a), ScaleMatrix(GetCol(z), input[0]))
		input = ReverseMatrix(input)
		input = MyFilter(GetCol(b), GetCol(a), ScaleMatrix(GetCol(z), input[0]))
		input = ReverseMatrix(input)
		input = deletePadding(input, nfact)
	}

	return input, nil
}


//Inmplemented, not tested
func MyFilter(numCoef, denCoef, input []float64, zin [][]float64) ([]float64, error) {

	fmt.Printf("Len B coef:%v\nLen A coef:%v\n", len(numCoef), len(denCoef))
	fmt.Printf("Len input pre:%v\n", len(input))

	if len(numCoef) == 0 && len(denCoef) == 0 {
		return nil, errors.New("Test error")
	}

	if len(zin) == 0 {
		// Fill z with zeros
	}

	//Normalize
	for i := range numCoef {
		numCoef[i] = numCoef[i] / denCoef[0]
		denCoef[i] = denCoef[i] / denCoef[0]
	}

	// Fill output with zeros
	output := make([]float64, len(input))
	for i, _ := range output {
		output[i] = 0
	}

	n := len(denCoef)
	z := make([]float64, len(denCoef))

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

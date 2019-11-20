package iirfilter

import (
	"errors"
	"fmt"
	"log"
)

func Filfilt(sos [][]float64, g float64, ord int, input []float64) ([]float64, error) {
	fmt.Println("** Filfilt")
	npts := len(input)
	b, a, zi, nfact, L, err := GetCoeffsAndInitialConditions(sos, g, ord, npts)
	if err != nil {
		return nil, err
	}

	output, err := FfOneChanCat(b, a, input, zi, nfact, L)
	if err != nil {
		return nil, err
	}

	return output, nil
}

//Not implemented
func GetCoeffsAndInitialConditions(sos [][]float64, g float64, ord int, npts int) ([][]float64, [][]float64, [][]float64, int, int, error) {
	fmt.Println("** GetCoeffsAndInitialConditions")
	L := len(sos) //down

	fmt.Printf("L:%v\n", len(sos))

	//ncols := len(sos[0]) //right
	fmt.Println("pass1")
	for i := 0; i < 3; i++ {
		sos[0][i] *= g
	}
	fmt.Println("pass2")

	a := make([][]float64, 3)
	for i := 0; i < 3; i++ {
		a[i] = make([]float64, L)
	}
	b := make([][]float64, 3)
	for i := 0; i < 3; i++ {
		b[i] = make([]float64, L)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < L; j++ {
			b[i][j] = sos[j][i]
		}
	}
	fmt.Println("pass3")
	for i := 0; i < 3; i++ {
		for j := 0; j < L; j++ {
			a[i][j] = sos[j][i+3]
		}
	}
	fmt.Println("pass4")
	nfact := 3 * ord

	if npts <= nfact {
		return nil, nil, nil, -1, -1, errors.New("ERROR input too short")
	}
	fmt.Println("pass5")
	zi := make([][]float64, 2)
	for i := 0; i < 2; i++ {
		zi[i] = make([]float64, L)
	}
	var rhs [2]float64
	var tmp2 [2][2]float64
	//eye := [][]float64{{1, 0}, {0, 1}}

	for i := 0; i < L; i++ {
		rhs[0] = b[1][i] - b[0][i]*a[1][i] //fila 2 y 3, columna i  (2x1)
		rhs[1] = b[2][i] - b[0][i]*a[2][i] //fila 2 y 3, columna i  (2x1)
		fmt.Printf("pass5.1.%v\n", i)
		tmp2[0][0] = 1 - a[1][i]
		tmp2[1][0] = 0 - a[2][i]
		tmp2[0][1] = 0 - 1
		tmp2[1][1] = 1 - 0
		fmt.Printf("pass5.2.%v\n", i)
		sol, err := CramerSolve(tmp2[0][0], tmp2[0][1], rhs[0], tmp2[1][0], tmp2[1][1], rhs[1])
		if err != nil {

		}
		fmt.Printf("pass5.4.%v\n", i)
		zi[0][i] = sol[0]
		zi[1][i] = sol[1]
		fmt.Printf("pass5.5.%v\n", i)

		//zi(:,i) = ( eye(2) - [-a(2:3,i),[1;0]] ) \ rhs; // eye matriz identidad,  concatenar vectores 2x1 para armar una matriz 2x2
	}
	fmt.Println("pass6")
	return b, a, zi, nfact, L, nil

}

//Implemented, but not tested
func FfOneChanCat(b, a [][]float64, input []float64, zi [][]float64, nfact, l int) ([]float64, error) {
	fmt.Println("** FfOneChanCat")
	if len(input) == 0 {
		return nil, errors.New("No input or zero")
	}
	for i := 0; i < l; i++ {
		//Padding
		var err error
		input, err = AddPadding(input, nfact) //add data to input
		if err != nil {
			log.Fatal(err)
		}
		acol := GetCol(a, i)
		bcol := GetCol(b, i)
		scale_zcol := ScaleMatrix(GetCol(zi, i), input[0])
		input, err = MyFilter(bcol, acol, input, scale_zcol)
		input = ReverseMatrix(input)
		input, err = MyFilter(bcol, acol, input, scale_zcol)
		input = ReverseMatrix(input)
		input, err = RemovePadding(input, nfact)
	}

	return input, nil
}

//Inmplemented, not tested
func MyFilter(numCoef, denCoef, input []float64, zin []float64) ([]float64, error) {
	fmt.Println("** MyFilter")
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

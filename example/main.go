package main

import (
	"fmt"

	dsp "github.com/dhinojosac/godsp"
)

var a []float64 = []float64{1.0, 2.0}
var b []float64 = []float64{2.0, 2.0}

func main() {
	fmt.Println("godsp example")
	var output []float64

	fmt.Printf("SIGNAL+NOISE\n")
	for _, val := range dsp.InputSignal_32_1kHz_15kHz {
		fmt.Printf("%v\n", val)
	}
	fmt.Printf("FILTERED\n")
	output, err := dsp.FIR_filter(dsp.Impulse_response, dsp.InputSignal_32_1kHz_15kHz)
	if err == nil {
		for i := range output {
			fmt.Printf("%v\n", output[i])
		}
	}
}

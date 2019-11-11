package main

import (
	"fmt"
	"os"

	dsp "github.com/dhinojosac/godsp"
)

var a []float64 = []float64{1.0, 2.0}
var b []float64 = []float64{2.0, 2.0}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

func main() {
	fmt.Println("godsp example")
	fout, err := os.Create("out.dat")
	check(err)

	var output []float64

	fmt.Printf("[!]Reading input data\n")
	/*
		for _, val := range dsp.InputSignal_32_1kHz_15kHz {
			fmt.Fprintf(fout, "%v\n", val)
		}
	*/
	/*
		for _, val := range dsp.IIR_signal_test {
			fmt.Fprintf(fout, "%v\n", val)
		}
	*/

	fmt.Printf("[!] Filtering...\n")
	//output, err = dsp.FIR_filter(dsp.Impulse_response, dsp.InputSignal_32_1kHz_15kHz)
	output, err = dsp.IIR_filter(dsp.IIR_num, dsp.IIR_den, dsp.IIR_signal_test)
	check(err)

	for i := range output {
		fmt.Fprintf(fout, "%v\n", output[i])
	}
	fmt.Printf("[!] Filtered!\n")

}

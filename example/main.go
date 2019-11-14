package main

import (
	"fmt"
	"math"
	"os"

	dsp "github.com/dhinojosac/godsp"
	"github.com/mjibson/go-dsp/spectral"
)

const FREQ_SAMPLE = 75 // Sampling frequency in Hz

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
	psd_out, err := os.Create("psd.dat")
	check(err)

	var input []float64
	var output []float64

	fmt.Printf("[!]Reading input data\n")
	/*
		for _, val := range dsp.InputSignal_32_1kHz_15kHz {
			fmt.Fprintf(fout, "%v\n", val)
		}
	*/

	for _, val := range dsp.IIR_signal_test {
		//fmt.Fprintf(fout, "%v\n", val)
		input = append(input, val)
	}
	for _, val := range dsp.IIR_A_COEF {
		input = append(input, val)
	}

	fmt.Printf("[!] Filtering...\n")
	//output, err = dsp.FIR_filter(dsp.Impulse_response, dsp.InputSignal_32_1kHz_15kHz)
	//output, err = dsp.IIR_filter(dsp.IIR_B_COEF, dsp.IIR_A_COEF, dsp.IIR_signal_test)
	output, err = dsp.IIR_filter_SOS(dsp.IIR_signal_test)
	check(err)

	for i := range output {
		fmt.Fprintf(fout, "%v\n", output[i])
		//fmt.Fprintf(fout, "%v %v\n", input[i], output[i])
	}
	fmt.Printf("[!] Filtered!\n")

	pwelch_options := &spectral.PwelchOptions{}
	pwelch_options.NFFT = 1024
	Pxx, freqs := spectral.Pwelch(output, FREQ_SAMPLE, pwelch_options)
	iPxx, _ := spectral.Pwelch(input, FREQ_SAMPLE, pwelch_options)
	for i, freq := range freqs {
		//fmt.Printf("f:%f  p:%f \n", freq, Pxx[i])
		Pxx_scaled := 10 * math.Log10(Pxx[i])   //pass to dB
		iPxx_scaled := 10 * math.Log10(iPxx[i]) //pass to dB
		fmt.Fprintf(psd_out, "%f, %f, %f\n", freq, Pxx_scaled, iPxx_scaled)
	}

}

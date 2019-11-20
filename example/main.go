package main

import (
	"fmt"
	"math"
	"os"

	dsp "github.com/dhinojosac/godsp"
	"github.com/dhinojosac/godsp/spectral"
)

const FREQ_SAMPLE = 75 // Sampling frequency in Hz

//SOSO Filter
var G float64 = 0.0012662

var B_Coef = [][]float64{
	{-6.625455e-01, 1.000000e+00},
	{-1.705328e+00, 1},
	{-1.999908e+00, 1.000000e+00},
	{-1.819789e+00, 1},
	{-1.999416e+00, 1.000000e+00},
	{-1.845643e+00, 1.000000e+00},
	{-1.999016e+00, 1.000000e+00},
	{-1.998843e+00, 1},
}

var A_Coef = [][]float64{
	{-1.800924e+00, 8.311287e-01},
	{-1.800524e+00, 8.743608e-01},
	{-1.921320e+00, 9.266921e-01},
	{-1.835910e+00, 9.409341e-01},
	{-1.974803e+00, 9.770868e-01},
	{-1.864612e+00, 9.838206e-01},
	{-1.990704e+00, 9.923671e-01},
	{-1.996655e+00, 9.981525e-01},
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

var sosfilter dsp.SosFilter // Create a SOS filter

func main() {
	fmt.Println("godsp example")

	sosfilter.SetFilter(G, B_Coef, A_Coef) // Set SOS filter parameters

	psd_out, err := os.Create("psd.dat")
	check(err)

	var input []float64
	var output []float64

	for _, val := range dsp.IIR_signal_test {
		input = append(input, val)
	}
	for _, val := range dsp.IIR_A_COEF {
		input = append(input, val)
	}

	fmt.Printf("[!] Filtering...\n")
	//output2, err = dsp.IIR_filter_AB(dsp.IIR_B_COEF, dsp.IIR_A_COEF, dsp.IIR_signal_test)
	output, err = dsp.IIR_filter_SOS_2(sosfilter, dsp.IIR_signal_test)
	check(err)

	pwelch_options := &spectral.PwelchOptions{}
	pwelch_options.NFFT = 1024
	Pxx, freqs := spectral.Pwelch_stop(output, FREQ_SAMPLE, pwelch_options, 20.0)
	iPxx, _ := spectral.Pwelch_stop(input, FREQ_SAMPLE, pwelch_options, 20.0)
	//Pxx2, _ := spectral.Pwelch_stop(output2, FREQ_SAMPLE, pwelch_options, 20.0)
	for i, _ := range Pxx {
		//fmt.Printf("f:%f  p:%f \n", freq, Pxx[i])
		Pxx_scaled := 10 * math.Log10(Pxx[i])   //pass to dB
		iPxx_scaled := 10 * math.Log10(iPxx[i]) //pass to dB
		//Pxx_scaled2 := 10 * math.Log10(Pxx2[i]) //pass to dB
		fmt.Fprintf(psd_out, "%f, %f, %f\n", freqs[i], Pxx_scaled, iPxx_scaled)
	}

}

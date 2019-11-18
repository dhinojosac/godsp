package main

/* USAGE:
 * ./pwelch-script.exe < data_geenerator
 */

import (
	"fmt"
	"math"
	"os"

	"github.com/dhinojosac/godsp/spectral"
)

const START_COMMAND = 9.910
const FINISH_COMMAND = 9.911

const TIME_SAMPLE = 10 // Duration of each batch to process in seconds
const FREQ_SAMPLE = 75 // Sampling frequency in Hz

const data_arr_len = TIME_SAMPLE * FREQ_SAMPLE // Max of data in buffer

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	psd_out, err := os.Create("psd.dat")
	check(err)

	for {

		var data_arr []float64
		i := 0
		temp := 0.0

		//check start command
		for {
			fmt.Scanf("%f\n", &temp)
			if temp == START_COMMAND {
				fmt.Printf("Start PSD\n")
				break
			}
		}

		// Fill buffer with data from file
		for {
			fmt.Scanf("%f\n", &temp)
			if temp == FINISH_COMMAND {
				fmt.Printf("Finish PSD\n")
				break
			}
			data_arr = append(data_arr, temp)
			i++
		}

		/*
			// Check array
				for j := range data_arr {
					fmt.Printf("%f\n", data_arr[j])
				}
		*/

		// Filter

		// To save the output to a file
		pwelch_options := &spectral.PwelchOptions{}
		pwelch_options.NFFT = (1024)
		Pxx, freqs := spectral.Pwelch_stop(data_arr, FREQ_SAMPLE, pwelch_options, 30.0)

		//fo, err := os.Create("pwelch_output.dat")
		//check(err)
		//defer fo.Close()

		for i, _ := range Pxx {
			//fmt.Printf("f:%f  p:%f \n", freq, Pxx[i])
			Pxx_scaled := 10 * math.Log10(Pxx[i]) //pass to dB
			fmt.Fprintf(psd_out, "%f %f\n", freqs[i], Pxx_scaled)
		}
		return

	}

}

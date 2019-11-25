package main

/* USAGE:
 * ./pwelch-script.exe < data_geenerator
 */

import (
	"fmt"
	"log"
	"os"

	"github.com/dhinojosac/godsp/outliers"
)

const START_COMMAND = 9.910
const FINISH_COMMAND = 9.911

const FREQ_SAMPLE = 60 // Sampling frequency in Hz

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	fout, err := os.Create("out.dat")
	check(err)
	defer fout.Close()

	for {

		var data_arr []float64
		i := 0
		temp := 0.0

		//check start command
		for {
			fmt.Scanf("%f\n", &temp)
			if temp == START_COMMAND {
				log.Printf("Start Data input\n")
				break
			}
		}

		// Fill buffer with data from file
		for {
			fmt.Scanf("%f\n", &temp)
			if temp == FINISH_COMMAND {
				log.Printf("Finish Data input\n")
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

		// remove outlier
		o, err := outliers.Filloutliers(data_arr)
		if err != nil {
			log.Fatalln(err)
		}
		// To save the output to a file

		for i, _ := range o {
			fmt.Fprintf(fout, "%f %f\n", data_arr[i], o[i])
		}
		return

	}

}

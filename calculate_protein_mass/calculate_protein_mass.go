package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func doinit() []byte {

	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, fmt.Errorf("No file name given"))
		os.Exit(1)
	}

	f, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	return f
}

func main() {

	peptide := strings.Trim(string(doinit()), "\n")

	var monoiso = map[string]float64{
		"A": 71.03711,
		"C": 103.00919,
		"D": 115.02694,
		"E": 129.04259,
		"F": 147.06841,
		"G": 57.02146,
		"H": 137.05891,
		"I": 113.08406,
		"K": 128.09496,
		"L": 113.08406,
		"M": 131.04049,
		"N": 114.04293,
		"P": 97.05276,
		"Q": 128.05858,
		"R": 156.10111,
		"S": 87.03203,
		"T": 101.04768,
		"V": 99.06841,
		"W": 186.07931,
		"Y": 163.06333,
	}

	var mass float64
	for i := range peptide {
		aa := string(peptide[i])
		if curmass, ok := monoiso[aa]; ok {
			mass += curmass
		} else {
			fmt.Fprintln(os.Stderr,
				fmt.Errorf("The amino acid %s is not in the monoisotopic mass table",
					aa))
			os.Exit(3)
		}
	}
	fmt.Println(mass)
}

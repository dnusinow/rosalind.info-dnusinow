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

	f:= strings.Trim(string(doinit()), " \n")

	var codontbl = map[string]string {
		"UUU": "F",
		"CUU": "L",
		"AUU": "I",
		"GUU": "V",
		"UUC": "F",
		"CUC": "L",
		"AUC": "I",
		"GUC": "V",
		"UUA": "L",
		"CUA": "L",
		"AUA": "I",
		"GUA": "V",
		"UUG": "L",
		"CUG": "L",
		"AUG": "M",
		"GUG": "V",
		"UCU": "S",
		"CCU": "P",
		"ACU": "T",
		"GCU": "A",
		"UCC": "S",
		"CCC": "P",
		"ACC": "T",
		"GCC": "A",
		"UCA": "S",
		"CCA": "P",
		"ACA": "T",
		"GCA": "A",
		"UCG": "S",
		"CCG": "P",
		"ACG": "T",
		"GCG": "A",
		"UAU": "Y",
		"CAU": "H",
		"AAU": "N",
		"GAU": "D",
		"UAC": "Y",
		"CAC": "H",
		"AAC": "N",
		"GAC": "D",
		"UAA": "Stop",
		"CAA": "Q",
		"AAA": "K",
		"GAA": "E",
		"UAG": "Stop",
		"CAG": "Q",
		"AAG": "K",
		"GAG": "E",
		"UGU": "C",
		"CGU": "R",
		"AGU": "S",
		"GGU": "G",
		"UGC": "C",
		"CGC": "R",
		"AGC": "S",
		"GGC": "G",
		"UGA": "Stop",
		"CGA": "R",
		"AGA": "R",
		"GGA": "G",
		"UGG": "W",
		"CGG": "R",
		"AGG": "R",
		"GGG": "G",
	}

	for i := 0 ; i < len(f) ; i += 3 {
		codon := f[i:i+3]
		aa := codontbl[codon]
		if aa == "Stop" {
			fmt.Print("\n")
			os.Exit(0)
		} else {
			fmt.Print(aa)
		}
	}
}

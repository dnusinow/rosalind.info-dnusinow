package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

	f := doinit()
	
	for i := len(f) - 1; i >= 0 ; i-- {
		switch f[i] {
		case 'A':
			fmt.Print("T")
		case 'T':
			fmt.Print("A")
		case 'C':
			fmt.Print("G")
		case 'G':
			fmt.Print("C")
		}
	}
	fmt.Print("\n")
}

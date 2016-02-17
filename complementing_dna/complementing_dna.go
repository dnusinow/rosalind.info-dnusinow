package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Printf("No file name given\n")
		os.Exit(1)
	}
	var fpath = os.Args[1]
	var f, err = ioutil.ReadFile(fpath)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(2)
	}

	for i := len(f) - 1; i >= 0 ; i-- {
		switch f[i] {
		case 'A':
			fmt.Printf("%c", 'T')
		case 'T':
			fmt.Printf("%c", 'A')
		case 'C':
			fmt.Printf("%c", 'G')
		case 'G':
			fmt.Printf("%c", 'C')
		case '\n':
		default:
			// fmt.Printf("Error: %c", f[i])
			// os.Exit(3)
		}
	}
	fmt.Printf("\n")
}

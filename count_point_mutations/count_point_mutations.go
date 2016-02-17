package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Printf("No file name given\n")
		os.Exit(1)
	}
	var fpath = os.Args[1]
	var fb, err = ioutil.ReadFile(fpath)
	var f = string(fb)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(2)
	}

	var seqs = strings.Split(f, "\n")

	if len(seqs[0]) != len(seqs[1]) {
		fmt.Printf("Sequences were of different length!")
	}

	var n int
	for i := 0 ; i < len(seqs[0]) ; i++ {
		if seqs[0][i] != seqs[1][i] {
			n++
		}
	}

	fmt.Printf("%d\n", n)
}

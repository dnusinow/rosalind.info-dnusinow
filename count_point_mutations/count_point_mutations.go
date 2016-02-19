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

	f := string(doinit())
	
	var seqs = strings.Split(f, "\n")

	if len(seqs[0]) != len(seqs[1]) {
		errmsg := fmt.Errorf("Sequences were of different length: %d and %d",
			len(seqs[0]), len(seqs[1]))
		fmt.Fprintln(os.Stderr, errmsg)
	}

	var n int
	for i := range(seqs[0]) {
		if seqs[0][i] != seqs[1][i] {
			n++
		}
	}

	fmt.Println(n)
}

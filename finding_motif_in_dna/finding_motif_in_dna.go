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
	fs := strings.Split(string(doinit()), "\n")
	template := fs[0]
	motif := fs[1]

templatel:
	for i := range template {
		for j := range motif {
			if len(template) <= i + j {
				continue templatel
			}
			if motif[j] != template[i+j] {
				continue templatel
			}
		}
		fmt.Print(i+1, " ")
	}
	fmt.Println("")
}

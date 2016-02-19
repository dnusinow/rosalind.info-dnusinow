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
	
	var n map[byte]int
	n = make(map[byte]int)
	for i := range(f) {
		n[f[i]]++
	}
	fmt.Println(n['A'], n['C'], n['G'], n['T'])
}

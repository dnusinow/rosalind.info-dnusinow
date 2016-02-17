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

	var n map[byte]int
	n = make(map[byte]int)
	for i := 0 ; i < len(f) ; i++ {
		n[f[i]]++
	}
	fmt.Printf("%d %d %d %d\n", n['A'], n['C'], n['G'], n['T'])
}

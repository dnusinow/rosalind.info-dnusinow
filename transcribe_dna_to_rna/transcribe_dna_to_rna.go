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

	for i := 0 ; i < len(f) ; i++ {
		if f[i] == 'T' {
			fmt.Printf("%c", 'U')
		} else {
			fmt.Printf("%c", f[i])
		}
	}
}

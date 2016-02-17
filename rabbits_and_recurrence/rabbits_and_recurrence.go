package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
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

	var fields = strings.Fields(f)
	var n, nerr = strconv.Atoi(fields[0])
	var k, kerr = strconv.Atoi(fields[1])

	if nerr != nil || kerr != nil {
		fmt.Printf(nerr.Error())
		fmt.Printf(kerr.Error())
		os.Exit(3)
	}

	var rabbitPairs[]int
	rabbitPairs = make([]int, n)
	rabbitPairs[0] = 1
	rabbitPairs[1] = 1
	
	for i := 2 ; i < n ; i++ {
		rabbitPairs[i] = rabbitPairs[i-1] + (k * rabbitPairs[i-2])
	}
	fmt.Printf("%d\n", rabbitPairs[n-1])
}

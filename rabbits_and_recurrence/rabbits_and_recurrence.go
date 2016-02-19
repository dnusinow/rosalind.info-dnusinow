package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
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

	f  := string(doinit())
	
	var fields = strings.Fields(f)
	var n, nerr = strconv.Atoi(fields[0])
	var k, kerr = strconv.Atoi(fields[1])

	if nerr != nil || kerr != nil {
		fmt.Println(os.Stderr, nerr.Error())
		fmt.Println(os.Stderr, kerr.Error())
		os.Exit(3)
	}

	var rabbitPairs[]int
	rabbitPairs = make([]int, n)
	rabbitPairs[0] = 1
	rabbitPairs[1] = 1
	
	for i := 2 ; i < n ; i++ {
		rabbitPairs[i] = rabbitPairs[i-1] + (k * rabbitPairs[i-2])
	}
	fmt.Println(rabbitPairs[n-1])
}

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

	f := doinit()
	
	var dbEntries = strings.Split(string(f), ">")
	
	var gc map[string]float32
	gc = make(map[string]float32)

	var maxID string
	var maxVal float32

dbEntryLoop: for i := 0 ; i < len(dbEntries) ; i++ {
		if len(dbEntries[i]) == 0 {
			continue dbEntryLoop
		}

		var entry = strings.SplitN(dbEntries[i], "\n", 2)
		var ngc, ntot int = 0, 0
		var seq = entry[1]
		for j := range(seq) {
			switch seq[j] {
			case 'G', 'C':
				ngc++
				fallthrough
			case 'A', 'T':
				ntot++
			}
		}
		var curgc = 100.0 * (float32(ngc) / float32(ntot))
		gc[entry[0]] = curgc
		if curgc > maxVal {
			maxVal = curgc
			maxID = entry[0]
		}
		// fmt.Printf("%s: %d total and %d gc. %f\n", entry[0], ntot, ngc, curgc)
	}

	fmt.Printf("%s\n%f\n", maxID, maxVal)
}

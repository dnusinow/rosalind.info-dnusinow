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
	var f, err = ioutil.ReadFile(fpath)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(2)
	}

	var dbEntries = strings.Split(string(f), ">")
	
	var gc map[string]float32
	gc = make(map[string]float32)

	var maxID string
	var maxVal float32

dbEntryLoop:
	for i := 0 ; i < len(dbEntries) ; i++ {
		if len(dbEntries[i]) == 0 {
			continue dbEntryLoop
		}

		var entry = strings.SplitN(dbEntries[i], "\n", 2)
		var ngc, ntot int = 0, 0
		var seq = entry[1]
		for j := 0 ; j < len(seq) ; j++ {
			if seq[j] == 'A' || seq[j] == 'T' || seq[j] == 'G' || seq[j] == 'C' {
				ntot++
			}
			if seq[j] == 'G' || seq[j] == 'C' {
				ngc++
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

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

func buildAllele(c byte) ([]byte, error) {
	switch c {
	case 'D':
		return []byte{'X', 'X'}, nil
	case 'H':
		return []byte{'X', 'x'}, nil
	case 'R':
		return []byte{'x', 'x'}, nil
	}
	return nil, fmt.Errorf("Could not build an allele for %c", c)
}

func main() {

	f := doinit()

	strfields := strings.Fields(string(f))
	var norgs [3]int
	var err error

	for i := 0 ; i < 3 ; i++ {
		norgs[i], err = strconv.Atoi(strfields[i])
		if err != nil {
			fmt.Fprintln(os.Stderr, fmt.Errorf("Error converting field number %d (%s) to integer", i, strfields[i]))
			os.Exit(3)
		}
	}

	var ntot int = norgs[0] + norgs[1] + norgs[2]

	var orgs []byte
	for i := 0 ; i < norgs[0] ; i++ {
		orgs = append(orgs, 'D')
	}
	for i := 0 ; i < norgs[1] ; i++ {
		orgs = append(orgs, 'H')
	}
	for i := 0 ; i < norgs[2] ; i++ {
		orgs = append(orgs, 'R')
	}
	
	var matingsWithDom, totalMatings int
	// iterate through pairs of organisms
	for i := 0 ; i < ntot ; i++ {
		for j := 0 ; j < ntot ; j++ {
			if i == j {	continue }

			// Simulate all possible pairings and record if we got a dominant allele
			ial, err := buildAllele(orgs[i])
			if err != nil {
				fmt.Println(os.Stderr, err.Error())
				os.Exit(4)
			}
			jal, err := buildAllele(orgs[j])
			if err != nil {
				fmt.Println(os.Stderr, err.Error())
				os.Exit(5)
			}
			for iali := 0 ; iali < len(ial) ; iali++ {
				for jali := 0 ; jali < len(jal) ; jali++ {
					totalMatings++
					if ial[iali] == 'X' || jal[jali] == 'X' {
						matingsWithDom++
					}
				}
			}
		}
	}

	// fmt.Printf("%d matings with dominant\n%d total matings\n%f probability\n",
	// 	matingsWithDom, totalMatings, float32(matingsWithDom) / float32(totalMatings))
	fmt.Println(float32(matingsWithDom) / float32(totalMatings))
}

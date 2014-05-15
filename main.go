package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Example Usage: Clustering animal, e.g. Clustering Aardvark
// i = the unique clustering % id of each input file, ranging from 0.70 to 0.95 in 0.05 increments
func main() {

	animal := os.Args[1]

OuterLoop:
	for i := 70; i < 100; i += 5 {
		//define the input file
		inputFileName := animal + "_0." + strconv.Itoa(i) + "_sorted"
		var r *bufio.Reader

		// open input file in read-only mode
		inputFile, inputError := os.Open(inputFileName)
		if inputError != nil {
			fmt.Println("An error occurred opening the input file.")
			return
		}
		defer inputFile.Close()

		r = bufio.NewReader(inputFile)

		for a := 0; a < 1; a++ {
			mystring, err := r.ReadString('\t')

			// fmt.Println(mystring)
			mystring = strings.TrimSuffix(mystring, "\t")

			x, _ := strconv.Atoi(mystring)
			// 	id, _ := strconv.Atoi(i)

			fmt.Println("Clustering % id:", i)
			fmt.Println("Max size of any cluster:", x)

			if x <= 100 ||
				(i == 90 && x <= 500) ||
				(i > 90) {
				fmt.Println("This is the best clustering % id:", i)
				break OuterLoop
			} else {
				fmt.Println("Not alignable, too many seqs per cluster!")
				continue
			}

			if err != nil {
				fmt.Println(err)
				return

			}
		}
	}

}

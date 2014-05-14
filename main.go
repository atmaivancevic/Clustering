package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Example Usage: Clustering animal, e.g. Clustering Aardvark
// id = the unique clustering % id of each input file, ranging from 0.70 to 0.95 in 0.05 increments
func main() {

	animal := os.Args[1]

	for i := 70; i < 100; i += 5 {

		//define the input file
		inputFileName := animal + "_0." + strconv.Itoa(i) + "_sorted"

		// open input file in read-only mode
		inputFile, inputError := os.Open(inputFileName)
		if inputError != nil {
			fmt.Println("An error occurred opening the input file.")
			return
		}
		defer inputFile.Close()

		// read the input file
		inputScanner := bufio.NewScanner(inputFile)
		for inputScanner.Scan() {
		}
		if err := inputScanner.Err(); err != nil {
			fmt.Println("An error occurred while reading the file.")
			return
		}

		if  {
			break
		}
		fmt.Println("Best clustering % identity:", i)
	}

}

// fmt.Println(inputScanner.Text(), i)

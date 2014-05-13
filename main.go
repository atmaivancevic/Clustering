package main

import (
	"bufio"
	"fmt"
	"os"
)

// Example Usage: Clustering animal, e.g. Clustering Aardvark
// id = the unique clustering % id of each input file, ranging from 0.70 to 0.95 in 0.05 increments
func main() {

	animal := os.Args[1]
	id := os.Args[2]

	inputFileName := animal + "_" + id + "_sorted"

	// open inputFile in read-only mode
	inputFile, inputError := os.Open(inputFileName)
	if inputError != nil {
		fmt.Println("An error occurred opening the input file.")
		return
	}
	defer inputFile.Close()

	inputScanner := bufio.NewScanner(inputFile)
	for inputScanner.Scan() {
		fmt.Println(inputScanner.Text())
	}
	if err := inputScanner.Err(); err != nil {
		fmt.Println("An error occurred while reading the file.")
		return
	}
}

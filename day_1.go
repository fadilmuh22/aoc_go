package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// create function to read file from local with io and return the content with array of string
func readFile(filename string) []string {
	// read file from local
	file, err := os.Open(filename)
	// if error then print error
	if err != nil {
		log.Fatal(err)
	}
	// close file after read
	defer file.Close()
	// create array of string to store each line of file
	var lines []string
	// create scanner to read file line by line
	scanner := bufio.NewScanner(file)
	// loop until file end
	for scanner.Scan() {
		// add each line to array of string
		lines = append(lines, scanner.Text())
	}
	// return array of string
	return lines
}

func dayOne() {
	// read file from locala and store it in a variable line by line
	var lines []string = readFile("input_1")

	// sum the number grouped by empty line and find the biggest sum
	var biggestSum int = 0
	var sum int = 0
	// create array of number to store each group
	for i := 0; i < len(lines); i++ {
		// if line is empty then reset the sum
		if lines[i] == "" {
			// if sum is bigger than biggestSum then replace it
			if sum > biggestSum {
				biggestSum = sum
			}
			// reset the sum
			sum = 0
			continue
		}
		// convert string to int
		number, _ := strconv.Atoi(lines[i])
		// add number to sum
		sum += number
	}
	fmt.Printf("The biggest sum is %d\n", biggestSum)
}

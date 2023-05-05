package main

import (
	"fmt"
	"strconv"
)

func DayOne() {
	// read file from locala and store it in a variable line by line
	lines := ReadFile("input_1")

	// sum the number grouped by empty line and find the biggest sum
	var biggestSum int = 0
	var sum int = 0
	// create array of number to store each group
	for _, line := range lines {
		// if line is empty then reset the sum
		if line == "" {
			// if sum is bigger than biggestSum then replace it
			if sum > biggestSum {
				biggestSum = sum
			}
			// reset the sum
			sum = 0
			continue
		}
		// convert string to int
		number, _ := strconv.Atoi(line)
		// add number to sum
		sum += number
	}
	fmt.Printf("The biggest sum is %d\n", biggestSum)
}

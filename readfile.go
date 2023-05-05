package main

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(filename string) []string {
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

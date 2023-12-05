package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadLines(path string) []string {
	log.Println("Reading file:", path)

	// Open the file
	file, err := os.Open(path)

	if err != nil {
		log.Panic("Error opening file:", err)
	}

	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)

	// Init a lines variable
	var lines []string

	// Loop through the lines
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for errors
	err = scanner.Err()

	if err != nil {
		log.Panic("Error reading file:", err)
	}

	// Print the number of lines
	log.Println("Number of lines:", len(lines))

	return lines
}

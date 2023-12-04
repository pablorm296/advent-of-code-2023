package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Get the path of the target input file (expected to be the first argument)
	inputFile := os.Args[1]

	// Read the lines of the file
	lines := readLines(inputFile)

	// Get the result of part one
	partOneResult := partOne(lines)

	// Print the result of part one
	log.Println("Part one result:", partOneResult)

}

func readLines(path string) []string {
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

func partOne(lines []string) int {
	// For each lie, we assume there's a valid string
	// We have to get the first and last digit of the string
	// If there's only one digit, we assume it's the first and last
	// Once the digits are found, we have to combine them

	// Define a regular expression to match the string
	re := regexp.MustCompile(`[0-9]`)

	// Init a variable, to contain the digits
	var digits []int64

	for i := 0; i < len(lines); i++ {
		// Get the digits in the line
		digitsInLine := re.FindAllString(lines[i], -1)

		// If there are no digits, skip the line
		if len(digitsInLine) == 0 {
			continue
		}

		// Get the first and last digit
		firstDigit := digitsInLine[0]
		lastDigit := digitsInLine[len(digitsInLine)-1]

		// Combine the digits
		newDigit := firstDigit + lastDigit

		// Convert the new digit to an integer
		newDigitAsInt, err := strconv.ParseInt(newDigit, 10, 64)

		// Check for errors
		if err != nil {
			log.Panic("Error converting string to integer:", err)
		}

		digits = append(digits, newDigitAsInt)

		// For each line, print the line and the new digit
		log.Println("Line:", lines[i], "New digit:", newDigit)
	}

	// Sum the digits
	var sum int64

	for i := 0; i < len(digits); i++ {
		sum += digits[i]
	}

	return int(sum)
}

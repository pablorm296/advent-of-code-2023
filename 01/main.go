package main

import (
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/pablorm296/advent-of-code-2023/utils"
)

func main() {
	// Get the path of the target input file (expected to be the first argument)
	inputFile := os.Args[1]

	// Read the lines of the file
	lines := utils.ReadLines(inputFile)

	// Get the result of part one
	partOneResult := partOne(lines)

	// Print the result of part one
	log.Println("Part one result:", partOneResult)

	// Get the result of part two
	partTwoResult := PartTwo(lines)

	// Print the result of part two
	log.Println("Part two result:", partTwoResult)
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
		log.Println("Line ", i+1, ":", lines[i], "New digit:", newDigit)
	}

	// Sum the digits
	var sum int64

	for i := 0; i < len(digits); i++ {
		sum += digits[i]
	}

	return int(sum)
}

func PartTwo(lines []string) int {
	// Define an array of digits spelled out
	var spelledOutDigits = []string{
		"zero", "one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine",
	}

	// Define a map of spelled out digits
	var spelledOutDigitsMap = map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	// Init a variable, to contain the digits
	var digits []int64

	// For each line
	for i := 0; i < len(lines); i++ {
		// Init a variable, to contain the digitsInLine in the line
		var digitsInLine []string

		// Convert the line to a slice of runes
		lineAsRunes := []rune(lines[i])

		// Loop through the runes
		for j := 0; j < len(lineAsRunes); j++ {
			// Get the current rune
			currentRune := lineAsRunes[j]

			// If the rune is a number between 1 and 9,
			// convert it to an integer and append it to the digits
			if currentRune > '0' && currentRune <= '9' {
				digitsInLine = append(digitsInLine, string(currentRune))
				continue
			}

			// If the rune is not a number, loop through the digits
			for k := 0; k < len(spelledOutDigits); k++ {
				// Get the current digit
				currentDigit := spelledOutDigits[k]

				// If the current rune is the first letter of the current digit
				if currentRune == rune(currentDigit[0]) {
					// Get the length of the current digit
					currentDigitLength := len(currentDigit)

					// Check that the next letters of the current digit are in the line
					if (j + currentDigitLength) <= len(lineAsRunes) {
						// Get the content of the line, from the current rune to the end of the current digit
						content := string(lineAsRunes[j : j+currentDigitLength])

						// If the content is the current digit
						if content == currentDigit {
							// get the digit from the map
							digit := spelledOutDigitsMap[currentDigit]

							// Append the digit to the digitsInLine
							digitsInLine = append(digitsInLine, digit)
						}
					}
				}
			}
		}

		// If there are no digits in the line, skip the line
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

		// Append the new digit to the digits
		digits = append(digits, newDigitAsInt)

		// For each line, print the line and the new digit
		log.Println("Line ", i+1, ":", lines[i], "New digit:", newDigit)
	}

	// Sum the digits
	var sum int64

	for i := 0; i < len(digits); i++ {
		sum += digits[i]
	}

	return int(sum)
}

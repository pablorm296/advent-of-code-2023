package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	myutils "github.com/pablorm296/advent-of-code-2023/utils"
)

func main() {
	// Get the path of the target input file (expected to be the first argument)
	inputFile := os.Args[1]

	// Read the lines of the file
	lines := myutils.ReadLines(inputFile)

	// Get the result of part one
	partOneResult := partOne(lines)

	fmt.Println("Part one result:", partOneResult)

	// Get the result of part two
	partTwoResult := partTwo(lines)

	fmt.Println("Part two result:", partTwoResult)
}

// Parse line
// Given a line in the format "Game [number]: [n] [color]; [n] [color]; [n] [color]"
// Return a map with the keys: game, results
func getGameNumberAndResult(line string) map[string]string {
	numberAndResultRegex := regexp.MustCompile(`^Game (?P<number>\d+): (?P<result>.+)$`)

	match := numberAndResultRegex.FindStringSubmatch(line)

	result := make(map[string]string)

	for i, name := range numberAndResultRegex.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	return result
}

func getDrawsByGame(result string) []string {
	// Split the result by ";"
	draws := strings.Split(result, ";")

	// Remove any leading or trailing spaces
	for i := 0; i < len(draws); i++ {
		draws[i] = strings.TrimSpace(draws[i])
	}

	return draws
}

func getCubesByDraw(draw string) map[string]int {
	// Init a map to contain the cubes
	result := make(map[string]int)

	// Split the draw by " "
	drawElements := strings.Split(draw, ",")

	// Loop through the cubes
	for i := 0; i < len(drawElements); i++ {
		// Remove any leading or trailing spaces
		drawElements[i] = strings.TrimSpace(drawElements[i])

		// Split by " "
		cubes := strings.Split(drawElements[i], " ")

		// Loop through the cubes
		color := cubes[1]
		number := cubes[0]

		// Convert the number to an integer
		numberInt, err := strconv.Atoi(number)

		// Check for errors
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
		}

		// Add the cube to the map
		result[color] = numberInt
	}

	return result
}

func partOne(lines []string) int {
	// Init a slice to contain the valid games
	var validGames []int

	// Iterate over the lines
	for i := 0; i < len(lines); i++ {
		lineNumberAndResult := getGameNumberAndResult(lines[i])
		lineNumber := lineNumberAndResult["number"]
		lineResult := lineNumberAndResult["result"]
		lineDraws := getDrawsByGame(lineResult)

		lineIsValid := true

		// Iterate over the draws
		for j := 0; j < len(lineDraws); j++ {
			lineDraw := lineDraws[j]
			lineCubes := getCubesByDraw(lineDraw)

			// Check if the draw is possible given the restriction 12 red cubes, 13 green cubes, and 14 blue cubes
			if lineCubes["red"] > 12 || lineCubes["green"] > 13 || lineCubes["blue"] > 14 {
				lineIsValid = false
			}
		}

		// If the line is valid, add the game number to the validGames slice
		if lineIsValid {
			lineNumberInt, err := strconv.Atoi(lineNumber)

			// Check for errors
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
			}

			validGames = append(validGames, lineNumberInt)
		}

		fmt.Println("Line", i+1, ":", lineResult, "Valid:", lineIsValid)
	}

	// Sum the valid games
	var sum int

	for i := 0; i < len(validGames); i++ {
		sum += validGames[i]
	}

	return sum
}

func partTwo(lines []string) int {
	// Init a slice to contain the power
	var powers []int

	// Iterate over the lines
	for i := 0; i < len(lines); i++ {
		lineNumberAndResult := getGameNumberAndResult(lines[i])
		lineNumber := lineNumberAndResult["number"]
		lineResult := lineNumberAndResult["result"]
		lineDraws := getDrawsByGame(lineResult)

		// Get the minimum number of cubes needed to replicate the draw
		minimumRedCubes := 0
		minimumGreenCubes := 0
		minimumBlueCubes := 0

		// Iterate over the draws
		for j := 0; j < len(lineDraws); j++ {
			lineDraw := lineDraws[j]
			lineCubes := getCubesByDraw(lineDraw)

			// Get the minimum number of cubes needed to replicate the draw
			if lineCubes["red"] > minimumRedCubes {
				minimumRedCubes = lineCubes["red"]
			}

			if lineCubes["green"] > minimumGreenCubes {
				minimumGreenCubes = lineCubes["green"]
			}

			if lineCubes["blue"] > minimumBlueCubes {
				minimumBlueCubes = lineCubes["blue"]
			}
		}

		// Get the power
		power := minimumRedCubes * minimumGreenCubes * minimumBlueCubes

		// Add the power to the powers slice
		powers = append(powers, power)

		fmt.Println("Line", lineNumber, ":", lineResult, "|", "minRed:", minimumRedCubes, "minGreen:", minimumGreenCubes, "minBlue:", minimumBlueCubes, "Power:", power)
	}

	// Sum the powers
	var sum int

	for i := 0; i < len(powers); i++ {
		sum += powers[i]
	}

	return sum
}

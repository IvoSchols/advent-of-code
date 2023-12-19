package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Solution for Day 6 of Advent of Code 2023

func main() {
	fmt.Printf("Part One: %d\n", part_one())
	fmt.Printf("Part Two: %d\n", part_two())
}

// We try are trying to find the multiplicative of the number of ways each race can be won
// A race is won if the distance traveled is larger than the recorded distance in the time
// For each second we wait, we increase the speed by 1 so that when we race we have a higher speed in the remaining time
func part_one() int {
	inputFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	scanner.Scan()
	times := getIntegers(scanner.Text())
	scanner.Scan()
	distances := getIntegers(scanner.Text())

	waysToWin := 1 // Magic number so that we can multiply

	// Iterate over the races
	for i := 0; i < len(times); i++ {
		winCount := calculateWinCount(times[i], distances[i])
		waysToWin *= winCount

	}

	return waysToWin
}

func part_two() int {
	inputFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	scanner.Scan()
	time := getInteger(scanner.Text())
	scanner.Scan()
	distance := getInteger(scanner.Text())

	winCount := calculateWinCount(time, distance)
	return winCount
}

func getIntegers(line string) []int {
	split := strings.Fields(line)
	split = split[1:] // Skip the first element

	integers := make([]int, len(split))

	for i, s := range split {
		integer, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
		}
		integers[i] = integer
	}
	return integers
}

func getInteger(line string) int {
	split := strings.Fields(line)
	split = split[1:] // Skip the first element

	stringInteger := ""

	for _, s := range split {
		stringInteger += s
	}

	integer, err := strconv.Atoi(stringInteger)
	if err != nil {
		fmt.Println(err)
	}
	return integer
}

func calculateWinCount(time int, recordDistance int) int {
	winCount := 0

	// Iterate over the seconds in the race (skip second 0, as it has speed 0)
	for i := 1; i < time; i++ {
		// Speed is the amount of seconds we have waited
		speed := i
		remainingTime := time - i

		// Calculate the distance traveled
		distance := speed * remainingTime

		// If the distance traveled is larger than the recorded distance, we have won
		if distance > recordDistance {
			winCount++
		} else if winCount > 0 {
			// If we have already won, we can stop iterating as we get diminishing returns
			break
		}

	}
	return winCount
}

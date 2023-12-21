package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Solution for Day 9 of Advent of Code 2023

func main() {
	fmt.Printf("Part One: %d\n", part_one())
	fmt.Printf("Part Two: %d\n", part_two())
}

func part_one() int {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	sum := 0

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		historyString := strings.Fields(line)

		history := make([]int, 0, len(historyString))
		sequences := make([][]int, 0)

		// Create history as a slice of ints
		for _, value := range historyString {
			number, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			history = append(history, number)
		}

		// Create sequences
		sequence := history

		for !isAllZeroes(sequence) {
			sequences = append(sequences, sequence)
			sequence = buildDifference(sequence)
		}
		sequences = append(sequences, sequence)

		// Calculate diagonal
		diagonal := calculateDiagonal(sequences)
		sum += diagonal[0]
	}

	return sum
}

// Start at the bottom of the triangle and work up, we add zero for the bottom row
// And for each consecutive row we add the last value of the previous row and the last of the current row
func calculateDiagonal(sequences [][]int) []int {
	diagonal := make([]int, len(sequences))
	diagonal[len(diagonal)-1] = 0

	for i := len(sequences) - 2; i >= 0; i-- {
		sequence := sequences[i]
		lastSequenceElement := len(sequences[i]) - 1

		diagonal[i] = diagonal[i+1] + sequence[lastSequenceElement]
	}

	return diagonal

}

func buildDifference(slice []int) []int {
	difference := make([]int, len(slice)-1)
	for i := 0; i < len(difference); i++ {
		difference[i] = slice[i+1] - slice[i]
	}
	return difference
}

func isAllZeroes(slice []int) bool {
	for _, value := range slice {
		if value != 0 {
			return false
		}
	}
	return true
}

func part_two() int {
	return 0
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Solution for Day 4 of Advent of Code 2023
func main() {
	fmt.Printf("Part One: %d\n", part_one())
	fmt.Printf("Part Two: %d\n", part_two())
}

func part_one() int {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	sum := 0

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		_, winningNumbers, myNumbers := parseCard(line)

		score := countEqualNumbers(winningNumbers, myNumbers)
		sum += calculatePoints(score)
	}
	return sum
}

func part_two() int {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	// Use dynamic programming to build up the total score
	var scratchCardCount [188]int // max amount of scratch cards is 188 (input.txt)
	for i := 0; i < 188; i++ {
		scratchCardCount[i] = 1
	}
	totalScratchCards := 0

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		cardId, winningNumbers, myNumbers := parseCard(line)
		cardId-- // CardId starts at 1, but we want to start at 0

		// Sum up the total amount of scratch cards
		totalScratchCards += scratchCardCount[cardId]

		winningNumbersCount := countEqualNumbers(winningNumbers, myNumbers)

		// For each winning number, we get one more scratch card
		for j := 0; j < winningNumbersCount; j++ {
			// The amount of scratch cards we have determines how many scratch cards we get for each winning number
			scratchCardCount[cardId+j+1] += scratchCardCount[cardId]
		}
	}
	return totalScratchCards
}

func calculatePoints(score int) int {
	if score == 0 {
		return 0
	}
	score--
	result := 1

	for i := 0; i < score; i++ {
		result *= 2
	}
	return result
}

func countEqualNumbers(numbers []int, myNumbers []int) int {
	count := 0
	for _, number := range numbers {
		for _, myNumber := range myNumbers {
			if number == myNumber {
				count++
			}
		}
	}
	return count
}

func parseCard(line string) (int, []int, []int) {
	// Split the line into the card, winning numbers and my numbers
	split := strings.Split(line, ": ")

	stringCardId := strings.Fields(split[0][5:]) // Remove "Card " from the beginning of the string
	cardId, err := strconv.Atoi(stringCardId[0])

	if err != nil {
		fmt.Println(err)
	}

	split = strings.Split(split[1], " | ")

	// Split winning numbers by space, and convert to int
	winningNumbersString := strings.Fields(split[0])
	winningNumbers := make([]int, len(winningNumbersString))
	for i, v := range winningNumbersString {
		winningNumbers[i], _ = strconv.Atoi(v)
	}

	// Split my numbers by space, and convert to int
	myNumbersString := strings.Fields(split[1])
	myNumbers := make([]int, len(myNumbersString))
	for i, v := range myNumbersString {
		myNumbers[i], _ = strconv.Atoi(v)
	}

	return cardId, winningNumbers, myNumbers
}

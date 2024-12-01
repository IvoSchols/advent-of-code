package main

// Solution for Day 1 of Advent of Code 2023
// Sloppy solution, but it works

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	fmt.Printf("Part One:%d\n", part_one())
	fmt.Printf("Part Two:%d\n", part_two())
}
func part_one() int {
	// Open input.txt for reading
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	sum := 0

	// Create scanner to read input.txt
	scanner := bufio.NewScanner(inputFile)
	// Iterate through each line of input.txt
	for scanner.Scan() {

		line := scanner.Text()
		runes := []rune(line)

		// Create a slice to hold the two digits
		var string_digit [2]string

		// Scan the line forwards for the first digit
		for _, r := range runes {
			if unicode.IsDigit(r) {
				string_digit[0] = string(r)
				break
			}
		}
		// Scane the line backwards for the second digit
		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				string_digit[1] = string(runes[i])
				break
			}
		}

		// Concatenate the two digits and convert to int
		digit, err := strconv.Atoi(string_digit[0] + string_digit[1])

		if err != nil {
			fmt.Println(err)
		}

		// Add the digit to the sum
		sum += digit

	}

	return sum
}

func part_two() int {
	// Open input.txt for reading
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	// Define map to map valid substrings to digits
	valid_digit_map := make(map[string]string)

	// Define valid substrings to search for
	valid_digit_substrings := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	valid_digits := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	// Populate map
	for i := 0; i < len(valid_digit_substrings); i++ {
		valid_digit_map[valid_digit_substrings[i]] = valid_digits[i]
	}

	sum := 0

	// Create scanner to read input.txt
	scanner := bufio.NewScanner(inputFile)
	// Iterate through each line of input.txt
	for scanner.Scan() {

		line := scanner.Text()

		// Create a slice to hold the two digits
		var string_digit [2]string

		// Scan the line forwards for the first digit
		for i, c := range line {
			if unicode.IsDigit(c) {
				string_digit[0] = string(c)
				break
			}
			found_digit := false
			// Check if the current character is a valid substring
			for j := 0; j < len(valid_digit_substrings); j++ {
				last_letter_index := i + len(valid_digit_substrings[j])
				if last_letter_index > len(line) {
					continue
				}
				if line[i:last_letter_index] == valid_digit_substrings[j] {
					string_digit[0] = valid_digit_map[valid_digit_substrings[j]]
					found_digit = true
					break
				}

			}
			if found_digit {
				break
			}
		}
		// Scane the line backwards for the second digit
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				string_digit[1] = string(line[i])
				break
			}
			found_digit := false
			// Check if the current character is a valid substring in reverse
			for j := 0; j < len(valid_digit_substrings); j++ {
				// Iterate through each valid substring whilst reading the line backwards
				// The string is still read forwards, but the index is read backwards
				first_letter_index := i - len(valid_digit_substrings[j]) + 1
				if first_letter_index < 0 {
					continue
				}

				if line[first_letter_index:i+1] == valid_digit_substrings[j] {
					string_digit[1] = valid_digit_map[valid_digit_substrings[j]]
					found_digit = true
					break
				}
			}
			if found_digit {
				break
			}

		}

		// Concatenate the two digits and convert to int
		digit, err := strconv.Atoi(string_digit[0] + string_digit[1])

		if err != nil {
			fmt.Println(err)
		}

		// Add the digit to the sum
		sum += digit

	}
	return sum
}

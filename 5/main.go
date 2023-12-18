package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Solution for Day 5 of Advent of Code 2023

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

	scanner := bufio.NewScanner(inputFile)
	scanner.Scan()
	line := scanner.Text()
	seeds := getSeeds(line)
	scanner.Scan() // Skip the empty line

	// If there is an empty line, there is a map
	for scanner.Scan() {
		seeds = parseSeedMap(seeds, scanner)
	}

	// Return minimum location
	return slices.Min(seeds)
}

func getSeeds(line string) []int {
	seeds := strings.Split(line, "seeds: ")
	seeds = strings.Split(seeds[1], " ")

	seedsInt := make([]int, len(seeds))

	for i, seed := range seeds {
		seedInt, err := strconv.Atoi(seed)
		if err != nil {
			fmt.Println(err)
		}
		seedsInt[i] = seedInt
	}

	return seedsInt
}

func parseSeedMap(seeds []int, scanner *bufio.Scanner) []int {
	// Make a slice of the seeds
	mappedSeeds := make([]int, 0)

	scanner.Text() // map title
	// fmt.Println(mapTitle)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		split := strings.Split(line, " ")

		// Parse the mapping
		dest, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Println(err)
		}
		src, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Println(err)
		}
		m_len, err := strconv.Atoi(split[2])
		if err != nil {
			fmt.Println(err)
		}

		// Apply the mapping
		for j := 0; j < len(seeds); j++ {
			seed := seeds[j]

			if seed >= src && seed < src+m_len {
				// Store the mapped seed
				mappedSeed := dest + (seed - src)
				mappedSeeds = append(mappedSeeds, mappedSeed)
				// Remove the seed from the list
				seeds = append(seeds[:j], seeds[j+1:]...)
				j-- // Account for the removed seed
			}
		}
	}
	mappedSeeds = append(mappedSeeds, seeds...)

	return mappedSeeds
}

func part_two() int {
	inputFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Scan()
	line := scanner.Text()

	seeds := getSeedsWithRange(line)

	scanner.Scan() // Skip the empty line

	seedMaps := getSeedMaps(scanner)

	// Parse the seed maps
	for _, seedMap := range seedMaps {
		seeds = parseSeedMapWithRange(seeds, seedMap)
	}

	// Return minimum location
	min := seeds[0][0]
	for _, seed := range seeds {
		if seed[0] < min {
			min = seed[0]
		}
	}
	return min
}

func getSeedsWithRange(line string) [][2]int {
	split := strings.Split(line, "seeds: ")
	split = strings.Split(split[1], " ")

	seeds := make([][2]int, len(split)/2)

	for i := 0; i < len(split)-1; i += 2 {
		start, err := strconv.Atoi(split[i])
		if err != nil {
			fmt.Println(err)
		}
		length, err := strconv.Atoi(split[i+1])
		if err != nil {
			fmt.Println(err)
		}
		// Create a seed array of length 2, with start and length
		seed := [2]int{start, length}
		seeds[i/2] = seed
	}
	return seeds
}

func getSeedMaps(scanner *bufio.Scanner) [][][3]int {
	seedMaps := make([][][3]int, 0)

	// If there is an empty line, there is a map
	for scanner.Scan() {
		// scanner.Text() // map title

		seedMap := getSeedMap(scanner)
		seedMaps = append(seedMaps, seedMap)
	}
	return seedMaps
}

func getSeedMap(scanner *bufio.Scanner) [][3]int {
	seedMap := make([][3]int, 0)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if line == "" {
			break
		}
		split := strings.Split(line, " ")

		// Parse the mapping
		dest, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Println(err)
		}
		src, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Println(err)
		}
		m_len, err := strconv.Atoi(split[2])
		if err != nil {
			fmt.Println(err)
		}
		seedMap = append(seedMap, [3]int{dest, src, m_len})
	}
	return seedMap
}

func parseSeedMapWithRange(seeds [][2]int, seedMap [][3]int) [][2]int {
	// Eventually atleast all seeds will be mapped
	mappedSeeds := make([][2]int, 0, len(seeds))

	// Seeds are not mapped yet and may be added during the loop, therefore we need to loop over the seeds first
	for i := 0; i < len(seeds); i++ {
		seed := seeds[i]
		start := seed[0]
		length := seed[1]

		for j, m := range seedMap {
			dest := m[0]
			src := m[1]
			m_len := m[2]

			// Check if the seed is mapped and in range
			if isInMapRange(start, length, src, m_len) {
				// There are 4 mapping cases
				if shouldMapCompletely(start, length, src, m_len) { // 1. The seed is completely mapped
					mappedSeed := mapSeedCompletely(start, length, src, m_len, dest)
					mappedSeeds = append(mappedSeeds, mappedSeed)

				} else if shouldMapPartiallyLeft(start, length, src, m_len) { // 2. The seed is partially mapped on the left
					mappedSeed, unmappedSeed := mapSeedPartiallyLeft(start, length, src, m_len, dest)
					mappedSeeds = append(mappedSeeds, mappedSeed)

					seeds = append(seeds, unmappedSeed)

				} else if shouldMapPartiallyRight(start, length, src, m_len) { // 3. The seed is partially mapped on the right
					mappedSeed, unmappedSeed := mapSeedPartiallyRight(start, length, src, m_len, dest)
					mappedSeeds = append(mappedSeeds, mappedSeed)

					seeds = append(seeds, unmappedSeed)

				} else if shouldMapPartiallyMiddle(start, length, src, m_len) { // 4. The seed is partially mapped in the middle
					mappedSeed, unmappedSeedLeft, unmappedSeedRight := mapSeedPartiallyMiddle(start, length, src, m_len, dest)
					mappedSeeds = append(mappedSeeds, mappedSeed)

					seeds = append(seeds, unmappedSeedLeft)
					seeds = append(seeds, unmappedSeedRight)

				} else {
					message := fmt.Sprintf("Seed %d with start %d and length %d is not mapped, with src %d, m_len %d and dest %d \n", i, start, length, src, m_len, dest)
					panic(message)
				}
				break // Seed is mapped, break out of the loop
			}

			// If last j, seed is not mapped, add it to the mapped seeds
			if j == len(seedMap)-1 {
				mappedSeeds = append(mappedSeeds, seed)
			}

		}
	}
	return mappedSeeds
}

// There are 4 mapping cases and 2 non-mapping cases, invert non-mapping cases
func isInMapRange(start int, length int, src int, m_len int) bool {
	return !((src+m_len <= start) || (src >= start+length))
}

func shouldMapCompletely(start int, length int, src int, m_len int) bool {
	seed_end := start + length
	map_end := src + m_len
	return src <= start && map_end >= seed_end
}
func shouldMapPartiallyLeft(start int, length int, src int, m_len int) bool {
	seed_end := start + length
	map_end := src + m_len
	return src <= start && map_end < seed_end
}
func shouldMapPartiallyRight(start int, length int, src int, m_len int) bool {
	return src > start && src+m_len >= start+length
}
func shouldMapPartiallyMiddle(start int, length int, src int, m_len int) bool {
	return src > start && src+m_len < start+length
}

// Map the seed completely
func mapSeedCompletely(start int, length int, src int, m_len int, dest int) [2]int {
	mapped := [2]int{start - src + dest, length}
	return mapped
}

// Map the beginning of the seed and return unmapped, i.e.: mapped, unmapped
func mapSeedPartiallyLeft(start int, length int, src int, m_len int, dest int) ([2]int, [2]int) {
	partial_map_length := src + m_len - start

	mapped := [2]int{start - src + dest, partial_map_length}
	unmapped := [2]int{start + partial_map_length, length - partial_map_length}

	return mapped, unmapped
}

// Map the end of the seed and return unmapped, i.e.: mapped, unmapped
func mapSeedPartiallyRight(start int, length int, src int, m_len int, dest int) ([2]int, [2]int) {
	partial_map_length := start + length - src

	mapped := [2]int{dest, partial_map_length}
	unmapped := [2]int{start, length - partial_map_length}

	return mapped, unmapped
}

// Map the middle of the seed and return unmapped, i.e.: mapped, unmapped_left, unmapped_right
func mapSeedPartiallyMiddle(start int, length int, src int, m_len int, dest int) ([2]int, [2]int, [2]int) {
	left_map_length := src - start
	right_map_length := start + length - (src + m_len)

	mapped := [2]int{dest, m_len}
	left_unmapped := [2]int{start, left_map_length}
	right_unmapped := [2]int{src + m_len, right_map_length}

	return mapped, left_unmapped, right_unmapped
}

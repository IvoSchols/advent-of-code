package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Solution for Day 10 of Advent of Code 2023

func main() {
	fmt.Printf("Part One: %d\n", part_one())
	fmt.Printf("Part Two: %d\n", part_two())
}

// The pipes are arranged in a two-dimensional grid of tiles:
//
//	| is a vertical pipe connecting north and south.
//	- is a horizontal pipe connecting east and west.
//	L is a 90-degree bend connecting north and east.
//	J is a 90-degree bend connecting north and west.
//	7 is a 90-degree bend connecting south and west.
//	F is a 90-degree bend connecting south and east.
//	. is ground; there is no pipe in this tile.
//	S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
//
// Directions: pipe -> [north, east, south, west]
var directions = map[rune][4]bool{
	'|': {true, false, true, false},
	'-': {false, true, false, true},
	'L': {true, true, false, false},
	'J': {true, false, false, true},
	'7': {false, false, true, true},
	'F': {false, true, true, false},
	'.': {false, false, false, false},
}

func part_one() int {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	diagram, width, start := readDiagram(scanner)
	component := buildConnectedComponent(diagram, width, start)

	diameter := getDiameter(component)

	return diameter
}

// Construct network of size |V| x |D| as adjacency matrix, where |V| is the number of vertices and |D| is the number of directions
func readDiagram(scanner *bufio.Scanner) ([][4]bool, int, int) {
	// Store the network of pipes as a kind of adjacency matrix, very sparse but easy to work with and fast to access
	network := make([][4]bool, 0)

	// Populate network
	start := -1
	current := 0

	width := 0

	// Build network
	for scanner.Scan() {
		line := scanner.Text()
		width = len(line)

		for _, value := range line {
			network = append(network, directions[value])

			if value == 'S' {
				start = current
			}
			current++
		}
	}

	// Replace start node directions with connecting neighbors
	neighbour_north := start - width
	neighbour_east := start + 1
	neighbour_south := start + width
	neighbour_west := start - 1

	// In range, and north points to south, east points to west, etc.
	if neighbour_north >= 0 && network[neighbour_north][2] {
		network[start][0] = true
	}
	if neighbour_east < width*len(network) && network[neighbour_east][3] {
		network[start][1] = true
	}
	if neighbour_south < width*len(network) && network[neighbour_south][0] {
		network[start][2] = true
	}
	if neighbour_west >= 0 && network[neighbour_west][1] {
		network[start][3] = true
	}

	return network, width, start
}

// Build the connected component of the directions network (|V|x|D|) using width and starting at the given node
func buildConnectedComponent(diagram [][4]bool, width int, start int) [][2]int {
	// We know each node of the component has two neighbors and that the component is cyclic
	// Store node i with left, right neighbor as [i][2]int
	component := make([][2]int, 0)

	current := start

	// Build component using do-while, go right until we reach the start node
	visited := make([]bool, len(diagram))
	visited[current] = true

	for {
		neighbors := neighbors(current, diagram, width)
		component = append(component, [2]int{neighbors[0], neighbors[1]})

		// Go right if possible, otherwise go left
		if !visited[neighbors[0]] {
			current = neighbors[0] // Go right
			visited[current] = true
		} else if !visited[neighbors[1]] {
			current = neighbors[1] // Go left
			visited[current] = true
		} else {
			break
		}
	}

	return component
}

// Get the diameter of the cyclic network starting from the first node
func getDiameter(network [][2]int) int {
	diameter := 0

	// Diameter is the count of nodes divided by two, rounded up
	diameter = len(network) / 2

	return diameter
}

// Get neighbors of node i from network
func neighbors(current int, diagram [][4]bool, width int) [2]int {
	n := make([]int, 0, 2)
	for i, direction := range diagram[current] {
		if direction {
			// North, East, South, West
			switch {
			case i == 0:
				n = append(n, current-width)
			case i == 1:
				n = append(n, current+1)
			case i == 2:
				n = append(n, current+width)
			case i == 3:
				n = append(n, current-1)
			}
		}
	}
	return [2]int{n[0], n[1]}
}

func part_two() int {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	diagram, width, start := readDiagram(scanner)

	component := buildConnectedComponent(diagram, width, start)

	componentMap := make(map[int]bool)
	for _, node := range component {
		componentMap[node[0]] = true
		componentMap[node[1]] = true
	}

	sum := 0

	// Use point in polygon algorithm to determine if a point is inside the component
	// https://en.wikipedia.org/wiki/Point_in_polygon
	for i := 0; i < len(diagram); i++ {
		inPolygon := false

		for j := 0; j < width; j++ {

			// Flip inPolygon if we encounter a vertical border of the component
			if componentMap[i*width+j] && isVerticalBorder(diagram[i*width+j]) {
				inPolygon = !inPolygon
			} else if inPolygon && !componentMap[i*width+j] {
				sum++
			}

		}
	}

	return sum
}

// A vertical border is a pipe that connects north and south
// Or ann L or J pipe (north and east or north and west)
func isVerticalBorder(directions [4]bool) bool {
	return directions[0] && directions[2] || directions[0] && directions[1] || directions[0] && directions[3]
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Solution for Day 8 of Advent of Code 2023

type Network struct {
	instructions []rune
	nodes        map[string]*Node
	ghostStart   []*Node
}

type Node struct {
	left  *Node
	right *Node
	value string
}

func main() {
	fmt.Printf("Part One: %d\n", part_one())
	fmt.Printf("Part Two: %d\n", part_two())
}

func part_one() int {
	// Read the input file
	inputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	network := readMap(scanner)

	// Follow the instructions
	currentNode := network.nodes["AAA"]
	endNode := network.nodes["ZZZ"]

	instructionsFollowed := 0

	for {
		for _, instruction := range network.instructions {
			switch instruction {
			case 'L':
				currentNode = currentNode.left
			case 'R':
				currentNode = currentNode.right
			}
			instructionsFollowed++

			if currentNode == endNode {
				return instructionsFollowed
			}
		}
	}
}

func part_two() int {
	// Read the input file
	inputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	network := readMapGhosts(scanner)

	// Follow the instructions
	ghostStartNodes := network.ghostStart
	leastCommonInstructions := make([]int, len(ghostStartNodes))

	for i := range ghostStartNodes {
		currentNode := ghostStartNodes[i]
		instructionsFollowed := 0
		solved := false

		for !solved {
			for _, instruction := range network.instructions {
				switch instruction {
				case 'L':
					currentNode = currentNode.left
				case 'R':
					currentNode = currentNode.right
				}
				instructionsFollowed++

				if currentNode.value[2] == 'Z' {
					solved = true
					break
				}
			}
		}
		leastCommonInstructions[i] = instructionsFollowed
	}

	// Find the least common multiple of the instructions
	lcm := lcm(leastCommonInstructions...)
	return lcm
}

// Read the input file and return the instructions and the nodes with their edges
func readMap(scanner *bufio.Scanner) Network {
	// Read the input file, first line are the instructions
	scanner.Scan()
	instructions := []rune(scanner.Text())

	scanner.Scan() // skip empty line

	// Read the input file, following lines are the nodes
	nodes := make(map[string]*Node)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " = (")

		node := split[0]

		// Check if node is already in the map, if not add it
		if _, ok := nodes[node]; !ok {
			nodes[node] = &Node{value: node}
		}

		split = strings.Split(split[1], ", ")

		left := split[0]

		right := split[1]
		right = right[:len(right)-1] // Remove the trailing ")"

		// Check if left node is already in the map, if not add it
		if _, ok := nodes[left]; !ok {
			nodes[left] = &Node{value: left}
		}
		// Check if right node is already in the map, if not add it
		if _, ok := nodes[right]; !ok {
			nodes[right] = &Node{value: right}
		}

		// Add the nodes to the map
		nodeLeft := nodes[left]
		nodeRight := nodes[right]

		mNode := nodes[node]
		mNode.left = nodeLeft
		mNode.right = nodeRight

	}

	network := Network{instructions: instructions, nodes: nodes}
	return network
}

// Read the input file and return the instructions and the nodes with their edges
func readMapGhosts(scanner *bufio.Scanner) Network {
	// Read the input file, first line are the instructions
	scanner.Scan()
	instructions := []rune(scanner.Text())

	scanner.Scan() // skip empty line

	// Read the input file, following lines are the nodes
	nodes := make(map[string]*Node)
	startGhost := make([]*Node, 0)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " = (")

		node := split[0]

		// Check if node is already in the map, if not add it
		if _, ok := nodes[node]; !ok {
			nodes[node] = &Node{value: node}
		}

		split = strings.Split(split[1], ", ")

		left := split[0]

		right := split[1]
		right = right[:len(right)-1] // Remove the trailing ")"

		// Check if left node is already in the map, if not add it
		if _, ok := nodes[left]; !ok {
			nodes[left] = &Node{value: left}
		}
		// Check if right node is already in the map, if not add it
		if _, ok := nodes[right]; !ok {
			nodes[right] = &Node{value: right}
		}

		// Add the nodes to the map
		nodeLeft := nodes[left]
		nodeRight := nodes[right]

		mNode := nodes[node]
		mNode.left = nodeLeft
		mNode.right = nodeRight

		// Check if the node is a ghost node
		if node[2] == 'A' {
			startGhost = append(startGhost, mNode)
		}
	}

	network := Network{instructions: instructions, nodes: nodes, ghostStart: startGhost}
	return network
}

// Find the least common multiple of a slice of numbers
func lcm(list ...int) int {
	// Not sure if this is allowed in the general case
	if len(list) == 1 {
		return list[0]
	}

	a := list[0]
	b := lcm(list[1:]...)

	// Find the least common multiple of two numbers
	return a * b / gcd(a, b)
}

// Find the greatest common divisor of two numbers
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

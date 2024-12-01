package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

// Solution for Day 7 of Advent of Code 2023

type Hand struct {
	sortedCards []rune
	cards       []rune
	bid         int
}

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

	// High to low
	var cardOrder = []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

	hands := getHands(scanner, cardOrder)
	hands = orderHands(hands, cardOrder)

	// Calculate the winnings by multiplying the rank with the bid and summing the winnings
	winnings := calculateWinnings(hands)

	return winnings
}

func part_two() int {
	inputFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	// High to low
	var cardOrderJoker = []rune{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}

	hands := getHands(scanner, cardOrderJoker)
	hands = orderHandsJoker(hands, cardOrderJoker)

	// Calculate the winnings by multiplying the rank with the bid and summing the winnings
	winnings := calculateWinnings(hands)

	return winnings
}

func getHands(scanner *bufio.Scanner, cardOrder []rune) []Hand {
	hands := make([]Hand, 0)

	// Read the file line by line and order the cards in each hand by the card order
	for scanner.Scan() {
		line := scanner.Text()
		hand := getHand(line, cardOrder)
		hands = append(hands, hand)
	}

	return hands
}

// Returns the cards, ordered camelCards (low to high) and the bid
func getHand(line string, cardOrder []rune) Hand {
	split := strings.Fields(line)

	camelCards := []rune(split[0])

	sortedCards := make([]rune, len(camelCards))
	copy(sortedCards, camelCards)

	bid, _ := strconv.Atoi(split[1])

	// Order the cards in the hand by the card order
	sort.Slice(sortedCards, func(i, j int) bool {
		cardOrderIndexI := slices.Index(cardOrder, sortedCards[i])
		cardOrderIndexJ := slices.Index(cardOrder, sortedCards[j])

		return cardOrderIndexI < cardOrderIndexJ
	})

	hand := Hand{sortedCards, camelCards, bid}

	return hand
}

func isHandOrderLessThan(hand1 Hand, hand2 Hand, cardOrder []rune) bool {
	for i, card := range hand1.cards {
		cardOrderIndexI := slices.Index(cardOrder, card)
		cardOrderIndexJ := slices.Index(cardOrder, hand2.cards[i])

		// If index is higher, the card is lower in the card order
		// Remember: Is i less than j?
		if cardOrderIndexI > cardOrderIndexJ {
			return true
		} else if cardOrderIndexI < cardOrderIndexJ {
			return false
		}
	}
	return false
}

func orderHands(hands []Hand, cardOrder []rune) []Hand {

	// Rank the hands according to: five of a kind, four of a kind, full house, three of a kind, two pair, one pair, high card
	// If two hands are of same type, hand with higher card wins
	sort.SliceStable(hands, func(i, j int) bool {
		cardsI := hands[i].sortedCards
		cardsJ := hands[j].sortedCards

		// Five of a kind
		iHasFiveOfAKind := cardsI[0] == cardsI[1] && cardsI[1] == cardsI[2] && cardsI[2] == cardsI[3] && cardsI[3] == cardsI[4]
		jHasFiveOfAKind := cardsJ[0] == cardsJ[1] && cardsJ[1] == cardsJ[2] && cardsJ[2] == cardsJ[3] && cardsJ[3] == cardsJ[4]

		// Is i less than j?
		if iHasFiveOfAKind && !jHasFiveOfAKind {
			return false
		} else if !iHasFiveOfAKind && jHasFiveOfAKind {
			return true
		} else if iHasFiveOfAKind && jHasFiveOfAKind {
			return isHandOrderLessThan(hands[i], hands[j], cardOrder)
		}

		// Four of a kind
		iHasFourOfAKind := cardsI[0] == cardsI[1] && cardsI[1] == cardsI[2] && cardsI[2] == cardsI[3]
		iHasFourOfAKind = iHasFourOfAKind || cardsI[1] == cardsI[2] && cardsI[2] == cardsI[3] && cardsI[3] == cardsI[4]

		jHasFourOfAKind := cardsJ[0] == cardsJ[1] && cardsJ[1] == cardsJ[2] && cardsJ[2] == cardsJ[3]
		jHasFourOfAKind = jHasFourOfAKind || cardsJ[1] == cardsJ[2] && cardsJ[2] == cardsJ[3] && cardsJ[3] == cardsJ[4]

		if iHasFourOfAKind && !jHasFourOfAKind {
			return false
		} else if !iHasFourOfAKind && jHasFourOfAKind {
			return true
		} else if iHasFourOfAKind && jHasFourOfAKind {
			return isHandOrderLessThan(hands[i], hands[j], cardOrder)
		}

		// Full house
		iHasFullHouse := cardsI[0] == cardsI[1] && cardsI[1] == cardsI[2] && cardsI[3] == cardsI[4]
		iHasFullHouse = iHasFullHouse || cardsI[0] == cardsI[1] && cardsI[2] == cardsI[3] && cardsI[3] == cardsI[4]

		jHasFullHouse := cardsJ[0] == cardsJ[1] && cardsJ[1] == cardsJ[2] && cardsJ[3] == cardsJ[4]
		jHasFullHouse = jHasFullHouse || cardsJ[0] == cardsJ[1] && cardsJ[2] == cardsJ[3] && cardsJ[3] == cardsJ[4]

		if iHasFullHouse && !jHasFullHouse {
			return false
		} else if !iHasFullHouse && jHasFullHouse {
			return true
		} else if iHasFullHouse && jHasFullHouse {
			return isHandOrderLessThan(hands[i], hands[j], cardOrder)
		}

		// Three of a kind
		iHasThreeOfAKind := cardsI[0] == cardsI[1] && cardsI[1] == cardsI[2]
		iHasThreeOfAKind = iHasThreeOfAKind || cardsI[1] == cardsI[2] && cardsI[2] == cardsI[3]
		iHasThreeOfAKind = iHasThreeOfAKind || cardsI[2] == cardsI[3] && cardsI[3] == cardsI[4]

		jHasThreeOfAKind := cardsJ[0] == cardsJ[1] && cardsJ[1] == cardsJ[2]
		jHasThreeOfAKind = jHasThreeOfAKind || cardsJ[1] == cardsJ[2] && cardsJ[2] == cardsJ[3]
		jHasThreeOfAKind = jHasThreeOfAKind || cardsJ[2] == cardsJ[3] && cardsJ[3] == cardsJ[4]

		if iHasThreeOfAKind && !jHasThreeOfAKind {
			return false
		} else if !iHasThreeOfAKind && jHasThreeOfAKind {
			return true
		} else if iHasThreeOfAKind && jHasThreeOfAKind {
			return isHandOrderLessThan(hands[i], hands[j], cardOrder)
		}

		// Two pair
		iHasTwoPair := cardsI[0] == cardsI[1] && cardsI[2] == cardsI[3]
		iHasTwoPair = iHasTwoPair || cardsI[0] == cardsI[1] && cardsI[3] == cardsI[4]
		iHasTwoPair = iHasTwoPair || cardsI[1] == cardsI[2] && cardsI[3] == cardsI[4]

		jHasTwoPair := cardsJ[0] == cardsJ[1] && cardsJ[2] == cardsJ[3]
		jHasTwoPair = jHasTwoPair || cardsJ[0] == cardsJ[1] && cardsJ[3] == cardsJ[4]
		jHasTwoPair = jHasTwoPair || cardsJ[1] == cardsJ[2] && cardsJ[3] == cardsJ[4]

		if iHasTwoPair && !jHasTwoPair {
			return false
		} else if !iHasTwoPair && jHasTwoPair {
			return true
		} else if iHasTwoPair && jHasTwoPair {
			return isHandOrderLessThan(hands[i], hands[j], cardOrder)
		}

		// One pair (if any of the subsequent are equal)
		iHasOnePair := cardsI[0] == cardsI[1]
		iHasOnePair = iHasOnePair || cardsI[1] == cardsI[2]
		iHasOnePair = iHasOnePair || cardsI[2] == cardsI[3]
		iHasOnePair = iHasOnePair || cardsI[3] == cardsI[4]

		jHasOnePair := cardsJ[0] == cardsJ[1]
		jHasOnePair = jHasOnePair || cardsJ[1] == cardsJ[2]
		jHasOnePair = jHasOnePair || cardsJ[2] == cardsJ[3]
		jHasOnePair = jHasOnePair || cardsJ[3] == cardsJ[4]

		if iHasOnePair && !jHasOnePair {
			return false
		} else if !iHasOnePair && jHasOnePair {
			return true
		} else if iHasOnePair && jHasOnePair {
			return isHandOrderLessThan(hands[i], hands[j], cardOrder)
		}

		// High card
		return isHandOrderLessThan(hands[i], hands[j], cardOrder)
	})

	return hands
}

func calculateWinnings(hands []Hand) int {
	winnings := 0
	for i, hand := range hands {
		rank := i + 1
		winnings += hand.bid * rank
	}
	return winnings
}

// J is a joker and can be used as any card for determining the type of hand
func orderHandsJoker(hands []Hand, cardOrder []rune) []Hand {

	// Rank the hands according to: five of a kind, four of a kind, full house, three of a kind, two pair, one pair, high card
	// If two hands are of same type, hand with higher card wins
	sort.SliceStable(hands, func(i, j int) bool {
		cardsI := hands[i].sortedCards
		cardsJ := hands[j].sortedCards

		// Collapse the hands to count the number of cards of each type
		jokerCountI := 0
		jokerCountJ := 0

		collapsedHandI := make([]int, 0)
		collapsedHandJ := make([]int, 0)

		collapsedHandI = append(collapsedHandI, 1)
		collapsedHandJ = append(collapsedHandJ, 1)

		// Count the number of consecutive cards of the same type, when a distinct card is found, expand the slice
		// Hand I
		for k := 1; k < len(cardsI); k++ {
			if cardsI[k] == 'J' {
				jokerCountI++
			} else if cardsI[k] == cardsI[k-1] {
				collapsedHandI[len(collapsedHandI)-1]++
			} else {
				collapsedHandI = append(collapsedHandI, 1)
			}
		}
		// Hand J
		for k := 1; k < len(cardsJ); k++ {
			if cardsJ[k] == 'J' {
				jokerCountJ++
			} else if cardsJ[k] == cardsJ[k-1] {
				collapsedHandJ[len(collapsedHandJ)-1]++
			} else {
				collapsedHandJ = append(collapsedHandJ, 1)
			}
		}

		// Sort the collapsed hands in descending order
		sort.Slice(collapsedHandI, func(i, j int) bool {
			return collapsedHandI[i] > collapsedHandI[j]
		})
		sort.Slice(collapsedHandJ, func(i, j int) bool {
			return collapsedHandJ[i] > collapsedHandJ[j]
		})

		// Five of a kind
		iHasFiveOfAKind := collapsedHandI[0]+jokerCountI == 5
		jHasFiveOfAKind := collapsedHandJ[0]+jokerCountJ == 5

		// Is i less than j?
		if iHasFiveOfAKind && !jHasFiveOfAKind {
			return false
		} else if !iHasFiveOfAKind && jHasFiveOfAKind {
			return true
		} else if iHasFiveOfAKind && jHasFiveOfAKind {
			return isHandOrderLessThan(hands[i], hands[j], cardOrder)
		}

		// Four of a kind
		iHasFourOfAKind := collapsedHandI[0]+jokerCountI == 4
		jHasFourOfAKind := collapsedHandJ[0]+jokerCountJ == 4

		if iHasFourOfAKind && !jHasFourOfAKind {
			return false
		} else if !iHasFourOfAKind && jHasFourOfAKind {
			return true
		} else if iHasFourOfAKind && jHasFourOfAKind {
			return isHandOrderLessThan(hands[i], hands[j], cardOrder)
		}

		// Full house
		fullHouseJokerCountI := jokerCountI
		fullHouseJokerCountJ := jokerCountJ

		iHasFullHouse := false
		jHasFullHouse := false

		// Each borrowed can only be used once
		if collapsedHandI[0]+jokerCountI >= 3 {
			fullHouseJokerCountI -= 3 - collapsedHandI[0] // Subtract the borrowed jokers
			iHasFullHouse = collapsedHandI[1]+fullHouseJokerCountI == 2
		}
		if collapsedHandJ[0]+jokerCountJ >= 3 {
			fullHouseJokerCountJ -= 3 - collapsedHandJ[0] // Subtract the borrowed jokers
			jHasFullHouse = collapsedHandJ[1]+fullHouseJokerCountJ == 2
		}

		if iHasFullHouse && !jHasFullHouse {
			return false
		} else if !iHasFullHouse && jHasFullHouse {
			return true
		} else if iHasFullHouse && jHasFullHouse {
			return isHandOrderLessThan(hands[i], hands[j], cardOrder)
		}

		// Three of a kind
		iHasThreeOfAKind := collapsedHandI[0]+jokerCountI == 3
		jHasThreeOfAKind := collapsedHandJ[0]+jokerCountJ == 3

		if iHasThreeOfAKind && !jHasThreeOfAKind {
			return false
		} else if !iHasThreeOfAKind && jHasThreeOfAKind {
			return true
		} else if iHasThreeOfAKind && jHasThreeOfAKind {
			return isHandOrderLessThan(hands[i], hands[j], cardOrder)
		}

		// Two pair
		twoPairJokerCountI := jokerCountI
		twoPairJokerCountJ := jokerCountJ

		iHasTwoPair := false
		jHasTwoPair := false

		// Each borrowed can only be used once
		if collapsedHandI[0]+jokerCountI >= 2 {
			twoPairJokerCountI -= 2 - collapsedHandI[0] // Subtract the borrowed jokers
			iHasTwoPair = collapsedHandI[1]+twoPairJokerCountI == 2
		}
		if collapsedHandJ[0]+jokerCountJ >= 2 {
			twoPairJokerCountJ -= 2 - collapsedHandJ[0] // Subtract the borrowed jokers
			jHasTwoPair = collapsedHandJ[1]+twoPairJokerCountJ == 2
		}

		if iHasTwoPair && !jHasTwoPair {
			return false
		} else if !iHasTwoPair && jHasTwoPair {
			return true
		} else if iHasTwoPair && jHasTwoPair {
			return isHandOrderLessThan(hands[i], hands[j], cardOrder)
		}

		// One pair (if any of the subsequent are equal)
		iHasOnePair := collapsedHandI[0]+jokerCountI >= 2
		jHasOnePair := collapsedHandJ[0]+jokerCountJ >= 2

		if iHasOnePair && !jHasOnePair {
			return false
		} else if !iHasOnePair && jHasOnePair {
			return true
		} else if iHasOnePair && jHasOnePair {
			return isHandOrderLessThan(hands[i], hands[j], cardOrder)
		}

		// Equal type
		return isHandOrderLessThan(hands[i], hands[j], cardOrder)
	})

	return hands
}

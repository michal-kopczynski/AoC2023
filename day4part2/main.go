package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Failed to open faile input.txt")
	}
	defer file.Close()

	matches := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches = append(matches, checkPoints(scanner.Text()))
	}

	// set initial number of cards
	cards := []int{}
	for _, _ = range matches {
		cards = append(cards, 1)
	}

	// go through all cards
	for cardIdx, cardQuantity := range cards {
		// for each of the instance of the card (orignal and won)
		for numOfCards := 0; numOfCards < cardQuantity; numOfCards++ {
			// add (next) won cards according to the matches of current card
			for wonCardIdx := cardIdx + 1; wonCardIdx < cardIdx+1+matches[cardIdx]; wonCardIdx++ {
				cards[wonCardIdx]++
			}
		}
	}

	sum := 0

	// sum number of cards
	for _, card := range cards {
		sum += card
	}
	fmt.Println(cards)
	fmt.Println(sum)
}

func checkPoints(line string) int {
	matches := 0

	winningNumbers, yourNumbers := parseLine(line)

	for _, yourNumber := range yourNumbers {
		for _, winningNumber := range winningNumbers {
			if yourNumber == winningNumber {
				matches++
				break
			}
		}
	}

	return matches
}

func parseLine(line string) ([]int, []int) {
	var winning []int
	var my []int
	// fmt.Println(line)

	_, numbersText, found := strings.Cut(line, ": ")
	if !found {
		panic(": seperator not found")
	}

	winningNumbersText, myNumbersText, found := strings.Cut(numbersText, " | ")
	if !found {
		panic("| seperator not found")
	}

	winningNumbers := strings.Split(winningNumbersText, " ")
	for _, winningNumber := range winningNumbers {
		if winningNumber != "" {
			n, err := strconv.Atoi(winningNumber)
			if err != nil {
				panic(fmt.Sprintf("Failed to convert winnning number: %s to int: %s", winningNumber, err))
			}
			winning = append(winning, n)
		}
	}

	myNumbers := strings.Split(myNumbersText, " ")
	for _, myNumber := range myNumbers {
		if myNumber != "" {
			n, err := strconv.Atoi(myNumber)
			if err != nil {
				panic(fmt.Sprintf("Failed to convert my number: %s to int: %s", myNumber, err))
			}
			my = append(my, n)
		}
	}

	return winning, my
}

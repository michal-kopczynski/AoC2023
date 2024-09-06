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

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += checkPoints(scanner.Text())
	}

	fmt.Println(sum)
}

func checkPoints(line string) int {
	matches := 0

	winningNumbers, yourNumbers := parseLine(line)
	// fmt.Println(winningNumbers, yourNumbers)

	for _, yourNumber := range yourNumbers {
		for _, winningNumber := range winningNumbers {
			if yourNumber == winningNumber {
				matches++
				break
			}
		}
	}

	if matches == 0 {
		return 0
	}
	sum := 1
	for ; matches > 1; matches-- {
		sum *= 2
	}

	return sum
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

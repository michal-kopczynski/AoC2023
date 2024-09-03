package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Cannot open file input.txt")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		sum += parseLine(scanner.Text())
	}

	fmt.Println(sum)
}

func parseLine(line string) int {
	var firstDigit, lastDigit int

	runes := []rune(line)
	for i := 0; i < len(runes); i++ {
		num := int(runes[i] - '0')
		if num >= 0 && num <= 9 {
			firstDigit = num
			fmt.Println(firstDigit)
			break
		}
	}

	for i := len(runes) - 1; i >= 0; i-- {
		num := int(runes[i] - '0')
		if num >= 0 && num <= 9 {
			lastDigit = num
			break
		}
	}

	return 10*firstDigit + lastDigit
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var digitsMap = map[string]string{
	"zero":  "0o",
	"one":   "1e",
	"two":   "2o",
	"three": "3e",
	"four":  "4",
	"five":  "5e",
	"six":   "6",
	"seven": "7n",
	"eight": "8t",
	"nine":  "9e",
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Cannot open file input.txt")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		sum += parseLine(decodeDigits(scanner.Text()))
	}

	fmt.Println(sum)
}

func decodeDigits(input string) string {
	if len(input) == 0 {
		return ""
	}

	for k, v := range digitsMap {
		if strings.HasPrefix(input, k) {
			input = strings.Replace(input, k, v, 1)
			break
		}
	}

	return input[0:1] + decodeDigits(input[1:])
}

func parseLine(line string) int {
	var firstDigit, lastDigit int

	runes := []rune(line)
	for i := 0; i < len(runes); i++ {
		num := int(runes[i] - '0')
		if num >= 0 && num <= 9 {
			firstDigit = num
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

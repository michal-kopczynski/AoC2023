package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Number struct {
	value    int
	startPos int
	endPos   int
}

type Symbol struct {
	pos int
}

type Line struct {
	numbers []Number
	symbols []Symbol
}

type Schematic []Line

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Failed to open faile input.txt")
	}
	defer file.Close()

	schematic := parseInput(file)

	fmt.Println(schematic)

	// fmt.Println(findSumOfPartNumbers(schematic))
	fmt.Println(findSumOfGearRatios(schematic))
}

func parseInput(file *os.File) Schematic {
	var schematic Schematic
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		schematic = append(schematic, parseLine(line))
	}

	return schematic
}

func findSumOfGearRatios(s Schematic) int {
	sum := 0
	for lineIdx, _ := range s {
		fmt.Println("------Line idx:", lineIdx)
		for _, symbol := range s[lineIdx].symbols {
			fmt.Println("---Symbol:", symbol)
			adjacentNumbers := []int{}
			// check numbers in line above, unless first line
			if lineIdx != 0 {
				for _, number := range s[lineIdx-1].numbers {
					if symbol.pos >= number.startPos-1 && symbol.pos <= number.endPos+1 {
						adjacentNumbers = append(adjacentNumbers, number.value)
					}
				}
			}

			// check numbers in current line
			for _, number := range s[lineIdx].numbers {
				if symbol.pos >= number.startPos-1 && symbol.pos <= number.endPos+1 {
					adjacentNumbers = append(adjacentNumbers, number.value)
				}
			}

			if lineIdx != len(s)-1 {
				for _, number := range s[lineIdx+1].numbers {
					if symbol.pos >= number.startPos-1 && symbol.pos <= number.endPos+1 {
						adjacentNumbers = append(adjacentNumbers, number.value)
					}
				}
			}

			if len(adjacentNumbers) > 1 {
				sum += adjacentNumbers[0] * adjacentNumbers[1]
			}
		}
	}
	return sum
}

func findSumOfPartNumbers(s Schematic) int {
	sum := 0
	for lineIdx, _ := range s {
		fmt.Println("------Line idx:", lineIdx)
		for _, number := range s[lineIdx].numbers {

			isPart := false

			fmt.Println("---Number:", number)
			// check symbols in line above, unless first line
			if lineIdx != 0 {
				for _, symbol := range s[lineIdx-1].symbols {
					if symbol.pos >= number.startPos-1 && symbol.pos <= number.endPos+1 {
						sum += number.value
						isPart = true
						break
					}
				}
			}

			// check symbols in current line
			if isPart == false {
				for _, symbol := range s[lineIdx].symbols {
					if symbol.pos >= number.startPos-1 && symbol.pos <= number.endPos+1 {
						sum += number.value
						isPart = true
						break
					}
				}
			}

			// check symbol in line below, unless last line
			if isPart == false && lineIdx != len(s)-1 {
				for _, symbol := range s[lineIdx+1].symbols {
					if symbol.pos >= number.startPos-1 && symbol.pos <= number.endPos+1 {
						sum += number.value
						isPart = true
						break
					}
				}
			}

			fmt.Println("Is part:", isPart)
		}
	}
	return sum
}

func parseLine(l string) Line {
	var line Line

	startPos := 0
	pos := 0
	foundNumber := ""
	for len(l) > 0 {
		if foundNumber == "" {
			startPos = pos
		}

		if unicode.IsNumber(rune(l[0])) {
			foundNumber += string(l[0])
		} else {
			if foundNumber != "" {
				// we have found a number -> store it
				number, err := strconv.Atoi(foundNumber)
				if err != nil {
					panic(fmt.Sprintf("Failed to parse number: %s", err))
				}
				line.numbers = append(line.numbers, Number{value: number, startPos: startPos, endPos: pos - 1})
				foundNumber = "" // reset stored number
			}

			// found symbol
			if l[0] != '.' {
				line.symbols = append(line.symbols, Symbol{pos: pos})
			}
		}

		l = l[1:]
		pos++
	}

	// handle the case when line ends with number
	if foundNumber != "" {
		// we have finished a number -> store it
		number, err := strconv.Atoi(foundNumber)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse number: %s", err))
		}
		line.numbers = append(line.numbers, Number{value: number, startPos: startPos, endPos: pos - 1})
	}
	return line
}

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
		panic("Cannot open file input.txt")
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += verifyLine(scanner.Text())
	}

	fmt.Println(sum)
}

type game struct {
	id    int
	red   int
	green int
	blue  int
}

func verifyLine(line string) int {
	var game game

	idString, gamesString, found := strings.Cut(line, ": ")
	if !found {
		panic(": separator not found")
	}

	_, idString, found = strings.Cut(idString, " ")
	if !found {
		panic("space seperator not found")
	}

	var err error
	game.id, err = strconv.Atoi(idString)
	if err != nil {
		panic(fmt.Sprintf("Failed to convert game ID to int: %s", err))
	}

	gameStrings := strings.Split(gamesString, ";")
	for _, gameString := range gameStrings {
		cubeStrings := strings.Split(gameString, ", ")
		for _, cube := range cubeStrings {
			cube, _ = strings.CutPrefix(cube, " ")
			cubeQuantity, cubeColor, _ := strings.Cut(cube, " ")
			quantity, err := strconv.Atoi(cubeQuantity)
			if err != nil {
				panic(fmt.Sprintf("Failed to convert cube quantity to int: %s", err))
			}
			switch cubeColor {
			case "red":
				if quantity > game.red {
					game.red = quantity
				}
			case "green":
				if quantity > game.green {
					game.green = quantity
				}
			case "blue":
				if quantity > game.blue {
					game.blue = quantity
				}
			default:
				panic(fmt.Sprintf("Invalid color: %s", cubeColor))
			}
		}
	}

	if game.red <= 12 && game.green <= 13 && game.blue <= 14 {
		return game.id
	}

	return 0
}

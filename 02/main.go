package main

import (
	"strings"
	"os"
	"log"
)

func SumPowers(input string, maxRed, maxGreen, maxBlue int) int {
	sum := 0
	gameInputs := strings.Split(input, "\n")
	for _, gameInput := range gameInputs {
		if gameInput == "" {
			continue
		}
		game := Game{}
		game.Load(gameInput)
		r,g,b := game.GetFewestNumberCubes(maxRed, maxGreen, maxBlue)
		log.Printf("Game %d: %d red, %d green, %d blue", game.ID, r, g, b)
			sum += r * g * b
	}

	return sum
}
func SumPossibleGameIds(input string, maxRed, maxGreen, maxBlue int) int {
	sum := 0
	gameInputs := strings.Split(input, "\n")
	for _, gameInput := range gameInputs {
		if gameInput == "" {
			continue
		}
		game := Game{}
		game.Load(gameInput)
		if game.IsPossible(maxRed, maxGreen, maxBlue) {
			sum += game.ID
		}
	}

	return sum
}

func main() {
	input, _ := os.ReadFile("input.txt")
	log.Printf("Task1: %+v", SumPossibleGameIds(string(input), 12, 13, 14))
	log.Printf("Task2: %+v", SumPowers(string(input), 12, 13, 14))
}

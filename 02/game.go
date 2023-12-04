package main

import (
	// "os"
	// "slices"
	"strconv"
	"strings"
)

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
type Game struct {
	ID   int
	Sets []Set
}

type Set struct {
	Red   int
	Green int
	Blue  int
}

func (g *Game) Load(imput string) {
	gameId := strings.Split(strings.Split(imput, ":")[0], " ")[1]
	g.ID, _ = strconv.Atoi(strings.TrimSpace(gameId))
	for _, setInput := range strings.Split(strings.Split(imput, ":")[1], ";") {
		set := Set{}
		set.Load(setInput)
		g.Sets = append(g.Sets, set)
	}
}

func (g *Game) GetFewestNumberCubes(maxRed, maxGreen, maxBlue int) (red, green, blue int) {
	for _, set := range g.Sets {
		if set.Red > red {
			red = set.Red
		}
		if set.Blue > blue {
			blue = set.Blue
		}
		if set.Green > green {
			green = set.Green
		}
	}
	return red, green, blue
}
func (g *Game) IsPossible(maxRed, maxGreen, maxBlue int) bool {
	for _, set := range g.Sets {
		if set.Red > maxRed || set.Green > maxGreen || set.Blue > maxBlue {
			return false
		}
	}
	return true
}
func (s *Set) Load(input string) {
	//1 green, 3 red, 6 blue
	colorSets := strings.Split(input, ",")
	for _, colorSet := range colorSets {
		colorSet = strings.TrimSpace(colorSet)
		color := strings.Split(colorSet, " ")
		switch color[1] {
		case "green":
			s.Green, _ = strconv.Atoi(color[0])
		case "red":
			s.Red, _ = strconv.Atoi(color[0])
		case "blue":
			s.Blue, _ = strconv.Atoi(color[0])
		}
	}
}

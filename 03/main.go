package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"strconv"
	"strings"
)

func isSymbol(line string, position int) bool {
	notSymbols := "0123456789."
	return !strings.Contains(notSymbols, string(line[position]))
}

func isNumber(line string, position int) bool {
	notSymbols := "0123456789"
	return strings.Contains(notSymbols, string(line[position]))
}

func getFullNumber(line string, position int) int {
	fullNumber := string(line[position])
	log.Println("Position:", position)
	for i := position; i >= 0; i-- {
		if i == position {
			continue
		}
		log.Println("i:", i, "line[i]:", string(line[i]))
		if !isNumber(line, i) {
			break
		}
		fullNumber = string(line[i]) + fullNumber
	}
	for i := position; i <= len(line)-1; i++ {
		if i == position {
			continue
		}
		log.Println("i:", i, "line[i]:", string(line[i]))
		if !isNumber(line, i) {
			break
		}
		fullNumber += string(line[i])
	}
	log.Println("Full number:", fullNumber)
	number, _ := strconv.Atoi(fullNumber)

	return number
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func getNumbersNearSymbol(lines []string, colorArray [][]string, x, y int) []int {
	nearestNumbers := []int{}
	minX := max(x-1, 0)
	maxX := min(x+1, len(lines[0])-1)
	minY := max(y-1, 0)
	maxY := min(y+1, len(lines)-1)
	for i := minY; i <= maxY; i++ {
		isPreviousNumber := false
		for j := minX; j <= maxX; j++ {
			log.Printf("Checking [%d,%d] = %s", j, i, string(lines[i][j]))
			log.Printf("isPreviousNumber: %v", isPreviousNumber)
			if !isNumber(lines[i], j) {
				isPreviousNumber = false
			}

			if i == y && j == x {
				isPreviousNumber = false
				log.Printf("x,y [%d,%d] = %s", j, i, string(lines[i][j]))
				continue
			}
			if isPreviousNumber {
				log.Printf(" skipping as isPreviousNumber: %v", isPreviousNumber)
				continue
			}
			if isNumber(lines[i], j) {
				colorArray[i][j] = color.GreenString(string(lines[i][j]))
				nearestNumbers = append(nearestNumbers, getFullNumber(lines[i], j))
				isPreviousNumber = true
			}
		}
	}
	return nearestNumbers
}

func getAdjacentNumbers(input string) (sum int, colorArray [][]string) {
	log.Println("Input:", input)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		colorLine := []string{}
		for j, _ := range line {
			colorLine = append(colorLine, color.BlueString(string(line[j])))
		}
		colorArray = append(colorArray, colorLine)
	}
	for i, line := range lines {
		for j, _ := range line {
			if isSymbol(line, j) {
				log.Println("Symbol found at", j, i)
				colorArray[i][j] = color.RedString(string(line[j]))
				for _, number := range getNumbersNearSymbol(lines, colorArray, j, i) {
					sum += number
				}

			}
		}

	}
	return sum, colorArray
}

func getGears(input string) (sum int, colorArray [][]string) {
	log.Println("Input:", input)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		colorLine := []string{}
		for j, _ := range line {
			colorLine = append(colorLine, color.BlueString(string(line[j])))
		}
		colorArray = append(colorArray, colorLine)
	}
	for i, line := range lines {
		for j, _ := range line {
			if string(line[j]) == "*" {
				log.Println("Symbol found at", j, i)
				colorArray[i][j] = color.RedString(string(line[j]))
				if len(getNumbersNearSymbol(lines, colorArray, j, i)) != 2 {
					continue
				}
				sum += getNumbersNearSymbol(lines, colorArray, j, i)[0]* getNumbersNearSymbol(lines, colorArray, j, i)[1]

			}
		}

	}
	return sum, colorArray
}

func main() {
	input, _ := os.ReadFile("input.txt")
	sum, ca := getGears(string(input))
	for _, line := range ca {
		fmt.Println(strings.Join(line, ""))
	}
	log.Println("Task1:", sum)

}

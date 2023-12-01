package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func GetNumbers(input string, wordNumsEnabled bool) (numbers string) {
	allNumbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	wordedNumbers := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	foundNumbers := ""
	for i, char := range input {
		if slices.Contains(allNumbers, string(char)) {
			foundNumbers += string(char)
		} else if wordNumsEnabled {
			for key, value := range wordedNumbers {
				if strings.HasPrefix(input[i:], key) {
					foundNumbers += string(value)
				}
			}
		}
	}
	return foundNumbers
}

func Calibrate(input string, wordNumbsEnabled bool) (result int) {
	inputLines := strings.Split(input, "\n")
	for _, line := range inputLines {
		if line != "" {
			result += CalibrateLine(line, wordNumbsEnabled)
		}
	}

	return result
}

func CalibrateLine(input string, wordNumsEnabled bool) (sum int) {
	foundNumbers := GetNumbers(input, wordNumsEnabled)
	// if wordNumsEnabled {
	// 	log.Printf("input: %s, found numbers: %s", input, foundNumbers)
	// }
	firstNumber := string(foundNumbers[0])
	lastNumber := string(foundNumbers[len(foundNumbers)-1])
	sum, _ = strconv.Atoi(firstNumber + lastNumber)
	// log.Println(sum)

	return sum
}

func main() {
	input1, _ := os.ReadFile("input.txt")
	log.Println("Task 1:", Calibrate(string(input1), false))
	log.Println("Task 2:", Calibrate(string(input1), true))
}

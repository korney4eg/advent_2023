package main

import (
	"fmt"
	"github.com/fatih/color"
	"math"
	"os"
	"strconv"
	"strings"
)

type field struct {
	value int
	color string
}

type card struct {
	ID        int
	WinFields []*field
	AllFields []*field
	Count     int
}

func getFields(line string) []*field {
	fields := []*field{}
	for _, fieldString := range strings.Split(line, " ") {
		field := &field{}
		if fieldString == "" {
			continue
		}
		field.value, _ = strconv.Atoi(fieldString)
		field.color = color.WhiteString(string(fieldString))
		fields = append(fields, field)
	}
	return fields
}

func (c *card) Load(input string) {
	c.Count = 1
	cardId := strings.Split(strings.Split(input, ":")[0], " ")[1]
	c.ID, _ = strconv.Atoi(cardId)
	winFields := strings.Split(strings.Split(input, ":")[1], "|")[0]
	c.WinFields = getFields(winFields)
	allFields := strings.Split(strings.Split(input, ":")[1], "|")[1]
	c.AllFields = getFields(allFields)
}

func (c *card) String() string {
	winFields := ""
	for _, field := range c.WinFields {
		winFields += field.color + " "
	}
	allFields := ""
	for _, field := range c.AllFields {
		allFields += field.color + " "
	}
	repeatedCard := strings.Repeat(fmt.Sprintf("%d: [%s]|[%s]", c.ID, winFields, allFields)+"\n", c.Count)
	return repeatedCard[0 : len(repeatedCard)-1]
}

func (c *card) FindWinFields() []field {
	foundFields := []field{}
	for _, winField := range c.WinFields {
		for _, allField := range c.AllFields {
			if winField.value == allField.value {
				winField.color = color.YellowString(strconv.Itoa(winField.value))
				allField.color = color.YellowString(strconv.Itoa(winField.value))
				foundFields = append(foundFields, *winField)
			}
		}
	}
	return foundFields
}

func main() {
	input, _ := os.ReadFile("input.txt")
	sum := 0
	allCards := []card{}
	for _, cardInput := range strings.Split(string(input), "\n") {
		if cardInput == "" {
			continue
		}
		card := card{}
		card.Load(cardInput)
		allCards = append(allCards, card)
		winFields := card.FindWinFields()
		if len(winFields) > 0 {
			sum += int(math.Pow(2, float64(len(winFields)-1)))
		}
		fmt.Println(card.String())
	}
	fmt.Println("Task1: ", sum)
	totalCards := 0
	for i, card := range allCards {
		winFields := card.FindWinFields()
		totalCards += card.Count
		for count := 0; count < card.Count; count++ {
			for j := 1; j <= len(winFields); j++ {
				// fmt.Printf("giving from card %d +1 to card %d, so now card[%d]=%d\n", i, i+j, i+j, allCards[i+j].Count)
				allCards[i+j].Count++
			}
		}
	}
	fmt.Println("Task2: ", totalCards)
}

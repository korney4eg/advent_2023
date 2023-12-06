package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Rule struct {
	sourceRangeStart      int
	destinationRangeStart int
	rangeLength           int
}

type Mapper struct {
	name  string
	rules []Rule
}

func (r *Rule) load(input string) {
	r.destinationRangeStart = getInts(input)[0]
	r.sourceRangeStart = getInts(input)[1]
	r.rangeLength = getInts(input)[2]
}

func (m *Mapper) proceedOne(input int) (output int) {
	for _, rule := range m.rules {
		if isInRange(input, rule.sourceRangeStart, rule.sourceRangeStart+rule.rangeLength) {
			if isInRange(input, rule.sourceRangeStart, rule.sourceRangeStart+rule.rangeLength) {
				foundOutput := rule.destinationRangeStart + input - rule.sourceRangeStart
				// fmt.Printf("input %d in [%d,%d] -> %d\n", input, rule.sourceRangeStart, rule.sourceRangeStart+rule.rangeLength-1, foundOutput)
				return foundOutput
			}
		}

	}
	// fmt.Printf("input %d outside ranges -> %d\n", input, input)
	return input
}

func isInRange(input int, start int, end int) bool {
	return input >= start && input < end
}

func (m *Mapper) proceed(inputs []int) (outputs []int) {
	for _, input := range inputs {
		outputs = append(outputs, m.proceedOne(input))
	}

	return outputs
}

//	func (m *Mapper) proceedOneBack(input int) (output int) {
//		for _, rule := range m.rules {
//			if isInRange(input, rule.destinationRangeStart, rule.destinationRangeStart+rule.rangeLength) {
//				if isInRange(input, rule.sourceRangeStart, rule.sourceRangeStart+rule.rangeLength) {
//					foundOutput := rule.destinationRangeStart + input - rule.sourceRangeStart
//					fmt.Printf("input %d in [%d,%d] -> %d\n", input, rule.sourceRangeStart, rule.sourceRangeStart+rule.rangeLength, foundOutput)
//					return foundOutput
//				}
//			}
//			fmt.Printf("input %d outside ranges -> %d\n", input, input)
//
//		}
//		return input
//	}
func getInts(inputLine string) []int {
	ints := []int{}
	nums := strings.Split(inputLine, " ")
	for _, num := range nums {
		resultNum, _ := strconv.Atoi(num)
		ints = append(ints, resultNum)
	}
	return ints

}

func getSeeds(input string) []int {
	seedsLine := strings.Split(input, "\n")[0]
	seedNums := strings.Split(seedsLine, ": ")[1]

	return getInts(seedNums)

}

func getMappers(input string) (mappers []Mapper) {
	lines := strings.Split(input, "\n")[2:]
	mapper := Mapper{}
	for _, line := range lines {
		if line == "" {
			if mapper.name != "" {
				mappers = append(mappers, mapper)
				sort.SliceStable(mapper.rules, func(i, j int) bool {
					return mapper.rules[i].destinationRangeStart < mapper.rules[j].destinationRangeStart
				})

			}
			mapper = Mapper{}
			continue
		}
		if strings.Contains(line, " map:") {
			mapper.name = strings.Split(line, " map:")[0]
			continue
		}
		rule := Rule{}
		rule.load(line)
		mapper.rules = append(mapper.rules, rule)
	}
	return mappers

}

func getSeedRanged(seeds []int) (rangedSeeds []int) {
	fmt.Println("seeds len", len(seeds))
	firstSeed := true
	seed := 0
	for i := 0; i < len(seeds); i++ {
		if !firstSeed {
			for j := seed; j < seed+seeds[i]; j++ {
			rangedSeeds = append(rangedSeeds, j)
			}
			firstSeed = true
		} else {
			seed = seeds[i]
			firstSeed = false

		}

	}
	return rangedSeeds
}

func main() {
	inputF, _ := os.ReadFile("input.txt")

	seeds := getSeeds(string(inputF))
	fmt.Println("seeds", seeds)

	mappers := getMappers(string(inputF))
	inputs := seeds
	for _, mapper := range mappers {
		fmt.Println(inputs)
		fmt.Printf("%+v\n", mapper.name)
		inputs = mapper.proceed(inputs)
	}
	sort.Ints(inputs)
	fmt.Println("Task1:", inputs[0])

	fmt.Println("seeds", inputs)
	inputs = getSeedRanged(seeds)
	fmt.Println("inputs len", len(inputs))
	for _, mapper := range mappers {
		// fmt.Println(inputs)
		fmt.Printf("%+v\n", mapper.name)
		inputs = mapper.proceed(inputs)
	}
	sort.Ints(inputs)
	fmt.Println("Task2:", inputs[0])
}

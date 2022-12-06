package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func findDuplicates(strs []string) []rune {
	var duplicates []rune
	mapDuplicates := make(map[rune]int)

	fmt.Printf("for strings %s \n", strs)

	for _, str := range strs {
		visitedMap := make(map[rune]bool)
		for _, ch := range str {
			if _, exists := mapDuplicates[ch]; !exists {
				mapDuplicates[ch] = 1
			} else if !visitedMap[ch] {
				mapDuplicates[ch]++

				if mapDuplicates[ch] == len(strs) {
					duplicates = append(duplicates, ch)
				}
			}

			visitedMap[ch] = true
		}
	}

	fmt.Println(duplicates)
	return duplicates
}

func calculatePriority(badges []rune) int {
	var priority int
	for _, badge := range badges {
		if unicode.IsUpper(badge) {
			priority += int(badge) - 38
			continue
		}

		priority += int(badge) - 96
	}

	return priority
}

func main() {
	// strlower := "abcdefghijklmnopqrstuvwxyz"

	// fmt.Printf("lowercases:\n")
	// for _, s := range strlower {
	// 	fmt.Printf("%c - %d\n", s, s)
	// }

	// struper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// fmt.Printf("uppercases:\n")
	// for _, s := range struper {
	// 	fmt.Printf("%c - %d\n", s, s)
	// }

	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic("failed to read file")
	}

	lines := strings.Split(string(input), "\n")

	var priority int
	var priorityGroup int

	for _, rucksackstring := range lines {
		badges := findDuplicates([]string{rucksackstring[:(len(rucksackstring) / 2)], rucksackstring[(len(rucksackstring) / 2):]})

		//fmt.Printf("badges solo: %d\n", badges)

		priority += calculatePriority(badges)
	}

	for i := 0; i < len(lines); i += 3 {
		badges := findDuplicates([]string{lines[i], lines[i+1], lines[i+2]})

		//fmt.Printf("badges group: %d\n", badges)

		priorityGroup += calculatePriority(badges)
	}

	fmt.Printf("priority: %d\n", priority)
	fmt.Printf("group priority: %d\n", priorityGroup)
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic("failed to read file")
	}

	totalScore := 0

	// X lose      A Rock
	// Y draw 	   B Paper
	// Z win       C Scissors
	shapeMap := map[string]string{"X": "R", "Y": "P", "Z": "S", "A": "R", "B": "P", "C": "S"}

	toLose := map[string]string{"R": "S", "S": "P", "P": "R"}
	toWin := map[string]string{"S": "R", "P": "S", "R": "P"}

	shapeScore := map[string]int{"R": 1, "P": 2, "S": 3}

	for _, hand := range strings.Split(string(input), "\n") {
		if string(hand[2]) == "Y" {
			totalScore += 3 + shapeScore[shapeMap[string(hand[0])]]

			continue
		}

		if string(hand[2]) == "Z" {
			totalScore += 6 + shapeScore[toWin[shapeMap[string(hand[0])]]]

			continue
		}

		if string(hand[2]) == "X" {
			totalScore += shapeScore[toLose[shapeMap[string(hand[0])]]]

			continue
		}

	}

	fmt.Print(totalScore)
}

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic("failed to read file")
	}

	splitedInput := strings.Split(string(input), "\n\n")
	var calories = make([]int, len(splitedInput))

	for i, str := range splitedInput {
		for _, cal := range strings.Split(str, "\n") {
			calorie, err := strconv.Atoi(cal)
			if err != nil {
				panic("could not understand number of calories")
			}

			calories[i] += calorie
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))

	fmt.Println(calories)
	fmt.Println(calories[0] + calories[1] + calories[2])
}

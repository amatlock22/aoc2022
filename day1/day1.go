package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}

func part1() int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	maxCals, currentCals := 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			if currentCals > maxCals {
				maxCals = currentCals
			}

			currentCals = 0

			continue
		}

		cals, _ := strconv.Atoi(line)
		currentCals += cals
	}

	return maxCals
}

func part2() int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	allCals := []int{}
	currentCals := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			allCals = append(allCals, currentCals)
			currentCals = 0

			continue
		}

		cals, _ := strconv.Atoi(line)
		currentCals += cals
	}

	sort.Slice(allCals, func(a, b int) bool {
		return allCals[b] < allCals[a]
	})

	return allCals[0] + allCals[1] + allCals[2]
}

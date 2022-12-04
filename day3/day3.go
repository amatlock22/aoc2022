package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}

func part1() rune {
	var totalPriority rune

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rucksack := scanner.Text()
		comp0 := rucksack[0 : len(rucksack)/2]
		comp1 := rucksack[len(rucksack)/2:]
		runes0 := getRunes(comp0)

		for _, r := range comp1 {
			if _, ok := runes0[r]; ok {
				totalPriority += getPriority(r)

				break
			}
		}
	}

	return totalPriority
}

func getPriority(i rune) rune {
	if i <= 90 {
		return i - 38 // uppercase letters
	}

	return i - 96 // lowercase letters
}

func getRunes(compartment string) map[rune]bool {
	runes := make(map[rune]bool, len(compartment))

	for _, char := range compartment {
		runes[char] = true
	}

	return runes
}

func part2() rune {
	var totalPriority rune

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	groupCounter := 0
	rucksacks := make([]string, 3)

	for scanner.Scan() {
		rucksacks[groupCounter] = scanner.Text()

		if groupCounter == 2 {
			runes0 := getRunes(rucksacks[0])
			runes1 := getRunes(rucksacks[1])

			for _, r := range rucksacks[2] {
				if runes0[r] && runes1[r] {
					totalPriority += getPriority(r)

					break
				}
			}

			groupCounter = 0

			continue
		}

		groupCounter++
	}

	return totalPriority
}

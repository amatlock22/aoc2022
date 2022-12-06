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

func part1() int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	for i := range line {
		if !isDuplicate(line[i : i+4]) {
			return i + 4
		}
	}

	return 0
}

func isDuplicate(chars string) bool {
	duplicateMap := make(map[rune]bool)

	for _, c := range chars {
		if _, ok := duplicateMap[c]; ok {
			return true // there is a duplicate
		} else {
			duplicateMap[c] = true
		}
	}

	return false
}

func part2() int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	for i := range line {
		if !isDuplicate(line[i : i+14]) {
			return i + 14
		}
	}

	return 0
}

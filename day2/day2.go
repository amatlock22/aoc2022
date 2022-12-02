package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}

type outcomes struct {
	Win  string
	Lose string
	Draw string
}

var outcomesMap = map[string]outcomes{
	"A": {Win: "Z", Lose: "Y", Draw: "X"},
	"B": {Win: "X", Lose: "Z", Draw: "Y"},
	"C": {Win: "Y", Lose: "X", Draw: "Z"},
}

var valMap = map[string]int{"X": 1, "Y": 2, "Z": 3}

func part1() int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalPoints := 0

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		totalPoints += play(split[0], split[1]) + valMap[split[1]]
	}

	return totalPoints
}

func play(opp, me string) int {
	if outcomesMap[opp].Lose == me {
		return 6 // I win
	}

	if outcomesMap[opp].Win == me {
		return 0 // I lose
	}

	return 3 // Draw
}

func part2() int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalPoints := 0

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		opp := split[0]

		switch split[1] {
		case "X": // lose
			totalPoints += valMap[outcomesMap[opp].Win]
		case "Y": // draw
			totalPoints += valMap[outcomesMap[opp].Draw] + 3
		case "Z": // win
			totalPoints += valMap[outcomesMap[opp].Lose] + 6
		}
	}

	return totalPoints
}

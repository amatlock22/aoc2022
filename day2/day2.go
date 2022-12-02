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
		line := scanner.Text()
		split := strings.Split(line, " ")
		opp := split[0]
		me := split[1]
		outcomePoints := 0

		result := play(opp, me)

		switch result {
		case 0:
			outcomePoints = 3
		case 1:
			outcomePoints = 6
		}

		totalPoints += outcomePoints + valMap[me]
	}

	return totalPoints
}

func play(opp, me string) int {
	if outcomesMap[opp].Lose == me {
		return 1 // I win
	}

	if outcomesMap[opp].Win == me {
		return -1 // I lose
	}

	return 0 // Draw
}

func part2() int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalPoints := 0

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		opp := split[0]
		end := split[1]

		switch end {
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

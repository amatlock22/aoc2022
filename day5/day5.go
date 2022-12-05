package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}

type Move struct {
	Num   int
	Start int
	End   int
}

func part1() string {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	crates := parseCrates(scanner)

	for _, move := range parseMoves(scanner) {
		for i := 0; i < move.Num; i++ {
			start := move.Start - 1
			end := move.End - 1

			crates[end] = append([]string{crates[start][0]}, crates[end]...)
			crates[start] = crates[start][1:]
		}
	}

	return getTopCrates(crates)
}

func part2() string {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	crates := parseCrates(scanner)

	for _, move := range parseMoves(scanner) {
		start := move.Start - 1
		end := move.End - 1

		startCopy := make([]string, len(crates[start]))
		copy(startCopy, crates[start])

		crates[end] = append(startCopy[:move.Num], crates[end]...)
		crates[start] = crates[start][move.Num:]
	}

	return getTopCrates(crates)
}

func parseCrates(scanner *bufio.Scanner) [][]string {
	crates := [][]string{}

	for i := 0; i < 10; i++ { // only 10 crate stacks for problem
		crates = append(crates, make([]string, 0))
	}

	for scanner.Scan() {
		line := scanner.Text()

		if string(line[0:2]) == " 1" {
			scanner.Scan() // skip empty line

			break
		}

		for i := 1; i < len(line); i += 4 {
			val := string(line[i])
			stackNum := (i - 1) / 4

			if val != " " {
				crates[stackNum] = append(crates[stackNum], val)
			}
		}
	}

	return crates
}

func parseMoves(scanner *bufio.Scanner) (moves []Move) {
	for scanner.Scan() {
		moveLine := strings.Split(scanner.Text(), " ")

		moves = append(moves, Move{
			Num:   stringToInt(moveLine[1]),
			Start: stringToInt(moveLine[3]),
			End:   stringToInt(moveLine[5]),
		})
	}

	return moves
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func getTopCrates(crates [][]string) string {
	var topLine string

	for i := 0; i < len(crates); i++ {
		if len(crates[i]) > 0 {
			topLine += crates[i][0]
		}
	}

	return topLine
}

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
	moves := parseMoves(scanner)
	crates = executeMoves(moves, crates, "9000")

	var topLine string

	for i := range crates[0] {
		topLine += crates[findCurrentCrateRow(crates, i)][i]
	}

	return topLine
}

func part2() string {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	crates := parseCrates(scanner)
	moves := parseMoves(scanner)
	crates = executeMoves(moves, crates, "9001")

	var topLine string

	for i := range crates[0] {
		topLine += crates[findCurrentCrateRow(crates, i)][i]
	}

	return topLine
}

func executeMoves(moves []Move, crates [][]string, version string) [][]string {
	for _, move := range moves {
		for i := 0; i < move.Num; i++ {
			start := move.Start - 1
			end := move.End - 1
			currentRow := findCurrentCrateRow(crates, start)
			newRow := findNewCrateRow(crates, end)

			if version == "9001" { // lmao don't judge me
				currentRow += move.Num - 1 - i
			}

			if newRow == -1 { // prepend new row
				tmp := [][]string{make([]string, len(crates[0]))}
				crates = append(tmp, crates...)
				newRow = 0
				currentRow++
			}

			crates[newRow][end] = crates[currentRow][start]
			crates[currentRow][start] = ""
		}
	}

	return crates
}

func findNewCrateRow(crates [][]string, col int) int {
	for i := len(crates) - 1; i >= 0; i-- {
		if crates[i][col] == "" { // find first opening on top of a crate (or floor)
			return i
		}
	}

	return -1
}

func findCurrentCrateRow(crates [][]string, col int) int {
	for i := 0; i < len(crates); i++ {
		if crates[i][col] == "" { // skip until we find a crate
			continue
		}

		return i
	}

	return -1
}

func parseCrates(scanner *bufio.Scanner) (crates [][]string) {
	for scanner.Scan() {
		line := scanner.Text()

		if string(line[0:2]) == " 1" {
			scanner.Scan() // skip empty line

			break
		}

		crateLine := []string{}

		for i := 1; i < len(line); i += 4 {
			val := string(line[i])

			if val == " " {
				val = ""
			}

			crateLine = append(crateLine, val)
		}

		crates = append(crates, crateLine)
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

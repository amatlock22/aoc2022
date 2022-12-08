package main

import (
	"bufio"
	"fmt"
	"os"
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
	trees := [][]int{}

	for scanner.Scan() {
		trees = append(trees, stringArrToIntArr(scanner.Text()))
	}

	width := len(trees[0])
	height := len(trees)
	interiorCount := 0

	for row := 1; row < width-1; row++ {
		for col := 1; col < height-1; col++ {
			if isTallestUp(trees, row, col) || isTallestDown(trees, row, col) ||
				isTallestLeft(trees, row, col) || isTallestRight(trees, row, col) {
				interiorCount++
			}
		}
	}

	return width*2 + height*2 - 4 + interiorCount
}

func isTallestUp(trees [][]int, row, col int) bool {
	for i := row - 1; i >= 0; i-- {
		if trees[i][col] >= trees[row][col] {
			return false
		}
	}

	return true
}

func isTallestDown(trees [][]int, row, col int) bool {
	for i := row + 1; i < len(trees); i++ {
		if trees[i][col] >= trees[row][col] {
			return false
		}
	}

	return true
}

func isTallestLeft(trees [][]int, row, col int) bool {
	for i := col - 1; i >= 0; i-- {
		if trees[row][i] >= trees[row][col] {
			return false
		}
	}

	return true
}

func isTallestRight(trees [][]int, row, col int) bool {
	for i := col + 1; i < len(trees[row]); i++ {
		if trees[row][i] >= trees[row][col] {
			return false
		}
	}

	return true
}

func stringArrToIntArr(line string) []int {
	ints := make([]int, len(line))

	for i := range line {
		num, err := strconv.Atoi(string(line[i]))
		if err != nil {
			panic(err)
		}

		ints[i] = num
	}

	return ints
}

func part2() int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	trees := [][]int{}

	for scanner.Scan() {
		trees = append(trees, stringArrToIntArr(scanner.Text()))
	}

	width := len(trees[0])
	height := len(trees)
	maxScenic := 0

	for row := 0; row < width-1; row++ {
		for col := 0; col < height-1; col++ {
			scenicScore := scenicUp(trees, row, col) * scenicDown(trees, row, col) *
				scenicLeft(trees, row, col) * scenicRight(trees, row, col)

			if scenicScore > maxScenic {
				maxScenic = scenicScore
			}
		}
	}

	return maxScenic
}

func scenicUp(trees [][]int, row, col int) int {
	counter := 0

	for i := row - 1; i >= 0; i-- {
		if trees[i][col] < trees[row][col] {
			counter++
		} else {
			counter++

			return counter
		}
	}

	return counter
}

func scenicDown(trees [][]int, row, col int) int {
	counter := 0

	for i := row + 1; i < len(trees); i++ {
		if trees[i][col] < trees[row][col] {
			counter++
		} else {
			counter++

			return counter
		}
	}

	return counter
}

func scenicLeft(trees [][]int, row, col int) int {
	counter := 0

	for i := col - 1; i >= 0; i-- {
		if trees[row][i] < trees[row][col] {
			counter++
		} else {
			counter++

			return counter
		}
	}

	return counter
}

func scenicRight(trees [][]int, row, col int) int {
	counter := 0

	for i := col + 1; i < len(trees[row]); i++ {
		if trees[row][i] < trees[row][col] {
			counter++
		} else {
			counter++

			return counter
		}
	}

	return counter
}

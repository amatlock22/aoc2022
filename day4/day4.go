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

func part1() int {
	var numPairs int

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if isContained(getAssignments(scanner.Text())) {
			numPairs++
		}
	}

	return numPairs
}

func getAssignments(line string) []int {
	assignments := make([]int, 4)
	pair := strings.Split(line, ",")
	first := strings.Split(pair[0], "-")
	second := strings.Split(pair[1], "-")

	assignments[0], _ = strconv.Atoi(first[0])
	assignments[1], _ = strconv.Atoi(first[1])
	assignments[2], _ = strconv.Atoi(second[0])
	assignments[3], _ = strconv.Atoi(second[1])

	return assignments
}

func isContained(nums []int) bool {
	return (nums[0] <= nums[2] && nums[1] >= nums[3]) || (nums[2] <= nums[0] && nums[3] >= nums[1])
}

func part2() int {
	var numPairs int

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		assignments := getAssignments(scanner.Text())

		if isInRange(assignments[0], assignments[2], assignments[3]) || isInRange(assignments[1], assignments[2], assignments[3]) || isInRange(assignments[2], assignments[0], assignments[1]) || isInRange(assignments[3], assignments[0], assignments[1]) {
			numPairs++
		}
	}

	return numPairs
}

func isInRange(num, lower, upper int) bool {
	return num >= lower && num <= upper
}

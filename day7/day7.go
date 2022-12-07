package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sizes := processDirSizes()
	fmt.Println("Part 1:", part1(sizes))
	fmt.Println("Part 2:", part2(sizes))
}

func part1(sizes map[string]int) int {
	sum := 0

	for _, v := range sizes {
		if v <= 100000 {
			sum += v
		}
	}

	return sum
}

func part2(sizes map[string]int) int {
	minSpace := math.MaxInt
	totalSpace := 70000000
	updateSpace := 30000000
	neededSpace := updateSpace - totalSpace + sizes["root"]

	for _, v := range sizes {
		if v >= neededSpace && v < minSpace {
			minSpace = v
		}
	}

	return minSpace
}

func processDirSizes() map[string]int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sizes := make(map[string]int)
	dirs := []string{"root"}

	scanner.Scan() // skip cd /

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		if split[0] == "$" { // command
			if split[1] == "cd" { // cd
				if split[2] == ".." { // remove last part of directory path
					dirs = dirs[:len(dirs)-1]
				} else { // add to directory path
					dirs = append(dirs, split[2])
				}
			} else { // not cd, don't care
				continue
			}
		} else { // not command
			s, err := strconv.Atoi(split[0])
			if err != nil { // not a size of a file so skip it
				continue
			}

			for i := range dirs { // for every dir in the path of current dir, add the size
				fullPath := strings.Join(dirs[:i+1], "/")
				sizes[fullPath] += s
			}
		}
	}

	return sizes
}

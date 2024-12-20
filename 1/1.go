package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fileData, err := os.ReadFile("1.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(fileData), "\n")
	numLines := len(lines)
	firstCol := make([]int, numLines)
	secondCol := make([]int, numLines)
	for i := 0; i < len(lines); i++ {
		splits := strings.Fields(lines[i])
		first, err1 := strconv.Atoi(splits[0])
		second, err2 := strconv.Atoi(splits[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error converting to int.")
		}
		firstCol[i], secondCol[i] = first, second
	}

	slices.Sort(firstCol)
	slices.Sort(secondCol)

	var diff int
	for i := 0; i < numLines; i++ {
		diff += abs(firstCol[i] - secondCol[i])
	}

	fmt.Println(diff)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

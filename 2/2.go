package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileData, err := os.ReadFile("2.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	reports := strings.Split(string(fileData), "\n")
	var numSafe int
	// Loop through reports (lines) in file
	for _, report := range reports {
		// Split line into levels
		reportLevelsString := strings.Fields(report)
		// Convert each string level to integer level
		reportLevelsInt := make([]int, len(reportLevelsString))
		for i, levelString := range reportLevelsString {
			level, err := strconv.Atoi(levelString)
			if err != nil {
				fmt.Println("Error converting to int.")
			}
			reportLevelsInt[i] = level
		}
		// Check if report is safe
		if isSafe(reportLevelsInt) {
			numSafe++
		}
	}
	fmt.Println(numSafe)
}

func isSafe(report []int) bool {
	levelDirection := 0 // 0 for unknown, 1 for increasing, -1 for decreasing
	for i := 1; i < len(report); i++ {
		if i == 1 {
			if report[i] > report[i-1] {
				levelDirection = 1
			} else if report[i] < report[i-1] {
				levelDirection = -1
			}
		}
		if (report[i] == report[i-1]) || (report[i] > report[i-1] && levelDirection == -1) || (report[i] < report[i-1] && levelDirection == 1) || (abs(report[i]-report[i-1]) > 3) {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

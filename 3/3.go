package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fileData, err := os.ReadFile("3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	r, _ := regexp.Compile(`mul\(\d+,\d+\)`)
	matches := r.FindAllString(string(fileData), -1)

	var runningSum int
	for _, match := range matches {
		r, _ := regexp.Compile(`\d+`)
		mulMatches := r.FindAllString(match, -1)
		leftMul, _ := strconv.Atoi(mulMatches[0])
		rightMul, _ := strconv.Atoi(mulMatches[1])
		runningSum += (leftMul * rightMul)
	}
	fmt.Println(runningSum)
}

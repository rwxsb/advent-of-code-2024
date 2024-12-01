package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	solutionPartOne()
	solutionPartTwo()
}

func getLists() ([]int, []int) {
	file, err := os.Open("input.txt")
	handleError(err)
	sc := bufio.NewScanner(file)
	var leftColumn []int
	var rightColumn []int

	for sc.Scan() {
		line := sc.Text()
		splitString := strings.Split(line, "   ")
		leftValue, err := strconv.Atoi(strings.Trim(splitString[0], " "))
		handleError(err)
		rightValue, err := strconv.Atoi(strings.Trim(splitString[1], " "))
		handleError(err)
		leftColumn = append(leftColumn, leftValue)
		rightColumn = append(rightColumn, rightValue)
	}

	return leftColumn, rightColumn
}

func solutionPartOne() {
	leftColumn, rightColumn := getLists()
	sort.Slice(leftColumn, func(i, j int) bool {
		return leftColumn[i] > leftColumn[j]
	})

	sort.Slice(rightColumn, func(i, j int) bool {
		return rightColumn[i] > rightColumn[j]
	})

	diffSum := 0
	for i := range leftColumn {
		diff := leftColumn[i] - rightColumn[i]
		if diff < 0 {
			diffSum += diff * -1
		} else {
			diffSum += diff
		}
	}

	println(diffSum)
}

func solutionPartTwo() {
	leftColum, rightColumn := getLists()
	rightColumnMap := make(map[string]int)
	for _, v := range rightColumn {
		key := strconv.Itoa(v)
		rightColumnMap[key] += 1
	}

	simScore := 0
	for _, v := range leftColum {
		simScore += v * rightColumnMap[strconv.Itoa(v)]
	}

	println(simScore)
}

package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func getLists() []string {
	file, err := os.Open("sampleInput.txt")
	handleError(err)
	sc := bufio.NewScanner(file)
	var lines []string

	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
	}

	return lines
}

func solutionPartOne() {
	lines := getLists()

	unSafeCount := 0

	for _, v := range lines {
		lineNums := strings.Split(v, " ")
		sign := 0
		for i := 0; i < len(lineNums)-1; i++ {
			curr, _ := strconv.Atoi(lineNums[i])
			next, _ := strconv.Atoi(lineNums[i+1])
			diff := next - curr
			absDiff := math.Abs((float64)(next - curr))
			if absDiff > 3 {
				unSafeCount++
				break
			}

			if (sign > 0 && diff < 0) || (sign < 0 && diff > 0) || (diff == 0) {
				unSafeCount++
				break
			}

			if diff < 0 {
				sign = -1
			} else if diff > 0 {
				sign = 1
			}
		}
	}
	println(len(lines) - unSafeCount)
}

func solutionPartTwo() {
	lines := getLists()

	unSafeCount := 0

	for _, v := range lines {
		lineNums := strings.Split(v, " ")
		monotonicExceptionCount := 0
		diffExceptionCount := 0

		sign := 0
		for i := 0; i < len(lineNums)-2; i++ {
			curr, _ := strconv.Atoi(lineNums[i])
			next, _ := strconv.Atoi(lineNums[i+1])
			nextNext, _ := strconv.Atoi(lineNums[i+2])
			diff := next - curr
			absDiff := math.Abs((float64)(next - curr))
			if absDiff > 3 {
				unSafeCount++
			}

			if (sign > 0 && diff < 0) || (sign < 0 && diff > 0) || (diff == 0) {
				if absDiff <= 3 && math.Abs((float64)(nextNext-curr)) <= 3 {
					continue
				} else {
					unSafeCount++
				}
			}

			if diff < 0 {
				sign = -1
			} else if diff > 0 {
				sign = 1
			}
		}

		println("line", v)
		println("monotonic", monotonicExceptionCount)
		println("diff", diffExceptionCount)
		println("unsafe", unSafeCount)
	}
	println(len(lines) - unSafeCount)
}
func main() {
	solutionPartOne()
	solutionPartTwo()
}

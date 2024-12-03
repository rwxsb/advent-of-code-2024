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
	file, err := os.Open("input.txt")
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

func convertToIntArray(arr []string) []int {
	toReturn := make([]int, len(arr))

	for idx, v := range arr {
		vInt, _ := strconv.Atoi(v)
		toReturn[idx] = vInt
	}

	return toReturn
}

func solutionPartTwo() {
	lines := getLists() // Assumes getLists() reads and returns the input as a slice of strings

	safeCount := 0

	for _, v := range lines {
		lineNums := strings.Split(v, " ")
		lineNumsAsInt := convertToIntArray(lineNums)

		//all increasing
		// all decreasing
		isMonotonic := isMonotonic(lineNumsAsInt)
		// diff is > 1 and < 3
		isJumpWithinLimit := jumpWithinLimit(lineNumsAsInt)
		valid := isMonotonic && isJumpWithinLimit

		if !valid {
			valid = problemDampanable(lineNumsAsInt)
		}

		if valid {
			safeCount++
		} else {
			println("unsafe report: ", v)
		}
	}

	println(safeCount)
}

func problemDampanable(nums []int) bool {
	for i := range nums {
		modifiedNums := make([]int, 0, len(nums)-1)
		modifiedNums = append(modifiedNums, nums[:i]...)
		modifiedNums = append(modifiedNums, nums[i+1:]...)

		if isMonotonic(modifiedNums) && jumpWithinLimit(modifiedNums) {
			return true
		}
	}
	return false
}

func jumpWithinLimit(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i]
		absDiff := math.Abs((float64)(diff))

		if absDiff > 3 || absDiff < 1 {
			return false
		}
	}

	return true
}

func isMonotonic(nums []int) bool {
	increasing := true
	decreasing := true

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			decreasing = false
		}
		if nums[i] > nums[i+1] {
			increasing = false
		}
	}

	return increasing || decreasing
}

func main() {
	solutionPartOne()
	solutionPartTwo()
}

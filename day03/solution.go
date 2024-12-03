package main

import (
	"os"
	"regexp"
	"strconv"
)

func solutionPartOne() {
	buf, _ := os.ReadFile("input.txt")
	str := string(buf)
	r := regexp.MustCompile(`(?m)\s*mul\s*\(\s*(\d+)\s*,\s*(\d+)\s*\)`)

	sum := 0
	for _, match := range r.FindAllStringSubmatch(str, -1) {
		l, _ := strconv.Atoi(match[1])
		r, _ := strconv.Atoi(match[2])
		sum += l * r
	}

	println(sum)
}

func solutionPartTwo() {
	buf, _ := os.ReadFile("input.txt")
	str := string(buf)
	r := regexp.MustCompile(`(?m)(mul)\((\d+),(\d+)\)|(do\(\))|(don't\(\))`)

	sum := 0
	instruction := "do()"
	for _, match := range r.FindAllStringSubmatch(str, -1) {
		if match[0] == "do()" || match[0] == "don't()" {
			instruction = match[0]
		} else if instruction == "do()" {
			l, _ := strconv.Atoi(match[2])
			r, _ := strconv.Atoi(match[3])
			sum += l * r
		}
	}

	println(sum)
}

func main() {
	solutionPartOne()
	solutionPartTwo()
}

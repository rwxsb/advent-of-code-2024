package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func solutionPartOne() {
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)
	var rules []string
	var pages []string

	for sc.Scan() {
		line := sc.Text()

		if len(line) <= 0 {
			break
		}

		rules = append(rules, line)
	}

	for sc.Scan() {
		line := sc.Text()

		pages = append(pages, line)
	}

	sum := 0
	for _, v := range pages {
		val := strings.Split(v, ",")
		valid := true
		for i := 0; i+1 < len(val); i++ {
			s := fmt.Sprintf("%s|%s", val[i], val[i+1])
			if !slices.Contains(rules, s) {
				valid = false
			}
		}

		if valid {
			value := val[len(val)/2]
			valInt, _ := strconv.Atoi(value)
			sum += valInt
		}
	}

	println(sum)

}

func solutionPartTwo() {
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)
	var rules []string
	var pages []string

	for sc.Scan() {
		line := sc.Text()

		if len(line) <= 0 {
			break
		}

		rules = append(rules, line)
	}

	for sc.Scan() {
		line := sc.Text()

		pages = append(pages, line)
	}

	var invalids []string
	for _, v := range pages {
		val := strings.Split(v, ",")
		valid := true
		for i := 0; i+1 < len(val); i++ {
			s := fmt.Sprintf("%s|%s", val[i], val[i+1])
			if !slices.Contains(rules, s) {
				valid = false
			}
		}

		if !valid {
			invalids = append(invalids, v)
		}
	}

	//reorder invalids according to rules
	var sortedInvalids []string
	for _, v := range invalids {
		val := strings.Split(v, ",")

		for k := 0; k < len(val); k++ {
			for i := 0; i+1 < len(val); i++ {
				j := i + 1
				s := fmt.Sprintf("%s|%s", val[i], val[j])
				s1 := fmt.Sprintf("%s|%s", val[j], val[i])
				if !slices.Contains(rules, s) && slices.Contains(rules, s1) {
					val[i], val[j] = val[j], val[i]
				}
			}
		}
		sortedInvalids = append(sortedInvalids, strings.Join(val, ","))
	}

	sum := 0
	for _, v := range sortedInvalids {
		val := strings.Split(v, ",")
		valid := true
		for i := 0; i+1 < len(val); i++ {
			s := fmt.Sprintf("%s|%s", val[i], val[i+1])
			if !slices.Contains(rules, s) {
				valid = false
			}
		}

		if valid {
			value := val[len(val)/2]
			valInt, _ := strconv.Atoi(value)
			sum += valInt
		}
	}

	println(sum)

}

func main() {
	solutionPartOne()
	solutionPartTwo()

}

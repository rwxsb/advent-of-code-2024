package main

import (
	"bufio"
	"os"
	"strings"
)

func countOccurences(str string, substr string) int {
	return strings.Count(str, substr)
}

func getMatrix() [][]string {
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)
	var matrix [][]string

	for sc.Scan() {
		line := sc.Text()
		row := strings.Split(line, "")

		matrix = append(matrix, row)
	}

	return matrix
}

func solutionPartOne() {
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)
	var matrix [][]string

	for sc.Scan() {
		line := sc.Text()
		row := strings.Split(line, "")

		matrix = append(matrix, row)

	}

	sum := 0
	for _, v := range matrix {
		joined := strings.Join(v, "")
		sum += countOccurences(joined, "XMAS")
		sum += countOccurences(joined, "SAMX")
	}

	for k, row := range matrix {
		topToBottom := make([]string, 140)
		for l := range row {
			topToBottom[l] = matrix[l][k]
		}
		joined := strings.Join(topToBottom, "")
		sum += countOccurences(joined, "XMAS")
		sum += countOccurences(joined, "SAMX")
	}

	rows := len(matrix)
	cols := len(matrix[0])
	var diagonals [][]string
	for offset := 0; offset < cols; offset++ {
		var diagonal []string
		for i := 0; i < rows && (i+offset) < cols; i++ {
			diagonal = append(diagonal, matrix[i][i+offset])
		}
		diagonals = append(diagonals, diagonal)
	}

	for offset := 1; offset < rows; offset++ {
		var diagonal []string
		for i := 0; i < cols && (i+offset) < rows; i++ {
			diagonal = append(diagonal, matrix[i+offset][i])
		}
		diagonals = append(diagonals, diagonal)
	}

	for startCol := 0; startCol < cols; startCol++ {
		var diagonal []string
		row, col := 0, startCol
		for row < rows && col >= 0 {
			diagonal = append(diagonal, matrix[row][col])
			row++
			col--
		}
		diagonals = append(diagonals, diagonal)
	}

	for startRow := 1; startRow < rows; startRow++ {
		var diagonal []string
		row, col := startRow, cols-1
		for row < rows && col >= 0 {
			diagonal = append(diagonal, matrix[row][col])
			row++
			col--
		}
		diagonals = append(diagonals, diagonal)
	}

	for _, diagonal := range diagonals {
		joined := strings.Join(diagonal, "")
		sum += countOccurences(joined, "XMAS")
		sum += countOccurences(joined, "SAMX")
	}

	println(sum)
}

func solutionPartTwo() {
	matrix := getMatrix()

	rows := len(matrix)
	cols := len(matrix[0])

	sum := 0
	for row := 0; row+2 < rows; row++ {
		for col := 0; col+2 < cols; col++ {
			var diag1 []string
			var diag2 []string
			for offset := 0; offset < 3; offset++ {
				diag1 = append(diag1, matrix[row+offset][col+offset])
				diag2 = append(diag2, matrix[row+offset][col+2-offset])
			}

			joined1 := strings.Join(diag1, "")
			joined2 := strings.Join(diag2, "")

			if (joined1 == "MAS" || joined1 == "SAM") && (joined2 == "MAS" || joined2 == "SAM") {
				sum++
			}
		}
	}

	println(sum)
}

func main() {
	solutionPartOne()
	solutionPartTwo()
}

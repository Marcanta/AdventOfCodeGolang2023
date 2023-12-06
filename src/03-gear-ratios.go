package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const lineLength = 140

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

var directions = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func isValid(x, y, rows, cols int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

func getAdjacentNumbers(grid [][]byte, x, y, rows, cols int) []int {
	adjacentNumbers := make([]int, 0)
	for _, dir := range directions {
		newX, newY := x+dir[0], y+dir[1]
		if isValid(newX, newY, rows, cols) && grid[newX][newY] != '.' {
			num := int(grid[newX][newY] - '0')
			adjacentNumbers = append(adjacentNumbers, num)
		}
	}
	return adjacentNumbers
}

func findAdjacentSum(grid [][]byte) int {
	symbols := "*#+$"
	rows := len(grid)
	if rows == 0 {
		return 0 // Empty grid, return 0 as the sum
	}
	cols := len(grid[0])
	sumOfAdjacentNumbers := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if contains(symbols, grid[i][j]) {
				adjacentNums := getAdjacentNumbers(grid, i, j, rows, cols)
				for _, num := range adjacentNums {
					sumOfAdjacentNumbers += num
				}
			}
		}
	}

	return sumOfAdjacentNumbers
}

func contains(symbols string, char byte) bool {
	for i := 0; i < len(symbols); i++ {
		if symbols[i] == char {
			return true
		}
	}
	return false
}

func main() {
	input := readFile("../inputs/03.in")

	splittedInput := strings.Split(input, "\n")
	fmt.Printf("splittedInput: %v\n", splittedInput)

	result := findAdjacentSum(splittedInput)
	fmt.Println("Sum of adjacent numbers to symbols:", result)
}

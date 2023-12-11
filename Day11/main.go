package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

}

func calc(rows, columns map[int]int, maps [][]int) int {
	result := 0

	for i, v := range maps {

	}
}

func findIndexesOfGalaxies(input [][]rune) [][]int {
	var result [][]int

	for i, v := range input {
		var temp []int
		for j, char := range v {
			if unicode.IsDigit(char) {
				temp[0] = i
				temp[1] = j
			}
		}
		if len(temp) > 0 {
			result = append(result, temp)
		}
	}

	return result
}

func createColumnsMap(input [][]rune) map[int]int {
	result := make(map[int]int)

	for i, v := range input {
		isEmpty := true
		for j, _ := range v {
			if input[j][i] != '.' {
				isEmpty = false
			}
		}
		if isEmpty {
			result[i] = 2
		} else {
			result[i] = 1
		}
	}

	return result
}

func createRowsMap(input [][]rune) map[int]int {
	result := make(map[int]int)

	for i, v := range input {
		isEmpty := true
		for _, char := range v {
			if char != '.' {
				isEmpty = false
			}
		}
		if isEmpty {
			result[i] = 2
		} else {
			result[i] = 1
		}
	}

	return result
}

func createGalaxyMatrix(input []string) [][]rune {
	var result [][]rune
	for _, v := range input {
		var temp []rune
		for _, char := range v {
			temp = append(temp, char)
		}
		result = append(result, temp)
	}
	return result
}

func read() []string {
	file, err := os.Open("C:\\Aoc2023\\Aoc2023\\Day8\\input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result []string

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return result
}

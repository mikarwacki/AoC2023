package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	input := read()
	matrix := createGalaxyMatrix(input)
	columns := createColumnsMap(matrix)
	rows := createRowsMap(matrix)
	indexes := findIndexesOfGalaxies(matrix)
	res := calc(rows, columns, indexes)
	fmt.Println(res)
}

func calc(rows, columns map[int]int, galaxies [][]int) int {
	result := 0
	var pairs [][]int

	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			temp := []int{i, j}
			pairs = append(pairs, temp)
		}

	}

	fmt.Println(pairs)
	fmt.Println(galaxies)

	for i := range pairs {
		a := galaxies[pairs[i][0]]
		b := galaxies[pairs[i][1]]

		xStart := int(math.Min(float64(a[0]), float64(b[0])))
		xEnd := int(math.Max(float64(a[0]), float64(b[0])))

		fmt.Println("X axis")
		fmt.Printf("%d %d ", xStart, xEnd)
		fmt.Println()

		yStart := int(math.Min(float64(a[1]), float64(b[1])))
		yEnd := int(math.Max(float64(a[1]), float64(b[1])))

		for j := xStart + 1; j < xEnd; j++ {
			fmt.Println(rows[j])
			result += rows[j]
		}

		fmt.Println("Y axis")
		fmt.Printf("%d %d ", yStart, yEnd)
		fmt.Println()

		for j := yStart + 1; j < yEnd; j++ {
			fmt.Println(columns[j])
			result += columns[j]
		}
	}

	return result
}

func findIndexesOfGalaxies(input [][]rune) [][]int {
	var result [][]int

	for i, v := range input {
		for j, char := range v {
			if char == '#' {
				temp := []int{i, j}
				result = append(result, temp)
			}
		}

	}

	return result
}

func createColumnsMap(input [][]rune) map[int]int {
	result := make(map[int]int)

	rows := len(input)
	columns := len(input[0])

	for i := 0; i < columns; i++ {
		isEmpty := true
		for j := 0; j < rows; j++ {
			if input[j][i] == '#' {
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
		for j, _ := range v {
			if input[i][j] == '#' {
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
	file, err := os.Open("C:\\Aoc2023\\Aoc2023\\Day11\\input.txt")
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

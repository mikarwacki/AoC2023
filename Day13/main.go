package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := read()
	patterns := splitPatterns(input)
	res := solve(patterns)
	fmt.Println(res)
}

func solve(patterns [][]string) int {
	res := 0

	for _, pattern := range patterns {
		res += checkRows(pattern)
		res += checkColumns(pattern)
	}

	return res
}

func checkRows(pattern []string) int {
	for i := 0; i < len(pattern); i++ {
		if strings.EqualFold(pattern[i], pattern[i+1]) {
			return i * 100
		}
	}
	return 0
}

func checkColumns(pattern []string) int {
	tpattern := createColumns(pattern)
	fmt.Println(tpattern)
	return checkRows(tpattern)
}

func createColumns(pattern []string) []string {
	res := make([]string, len(pattern[0]))

	for i := 0; i < len(pattern[0]); i++ {
		var temp strings.Builder
		for j := 0; j < len(pattern); j++ {
			temp.WriteByte(pattern[j][i])
		}
		res = append(res, temp.String())
	}

	return res
}

func splitPatterns(input []string) [][]string {
	var res [][]string
	var temp []string
	for _, line := range input {
		if len(line) > 0 {
			temp = append(temp, line)
		} else {
			res = append(res, temp)
			temp = temp[:0]
		}
	}
	return res
}

func read() []string {
	file, err := os.Open("C:\\Aoc2023\\Aoc2023\\Day12\\test.txt")
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

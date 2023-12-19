package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := read()
	lines, nums := formatInput(input)
	fmt.Println(input)
	fmt.Println(lines)
	fmt.Println(nums)
}

func calc(lines []string, nums []int) int {
	return 0
}

func formatInput(input []string) ([]string, [][]int) {
	lines := make([]string, len(input))
	nums := make([][]int, len(input))
	for i := range input {
		var numLine []int
		temp := strings.Fields(input[i])
		fmt.Println(temp)
		num := strings.Split(temp[1], ",")
		for _, v := range num {
			t, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			numLine = append(numLine, t)
		}
		nums[i] = numLine
		lines[i] = temp[0]
	}
	return lines, nums
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

package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions [][]int = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func main() {

}

func recCall(maze []string, count int) int {

}

func route()

func read() []string {
	file, err := os.Open("C:\\Aoc2023\\Aoc2023\\Day10\\test.txt")
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

}

func read() []string {
	file, err := os.Open("C:\\Users\\Mikolaj Karwacki\\Documents\\awesomeProject\\Day3" +
		"\\input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result []string

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return result
}

func sumParts(input []string) int {
	sum := 0
	for _, v := range input {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
		}
		sum += num
	}
	return sum
}

func traverseLine(line string, idx int) []int {
	return nil //TODO
}

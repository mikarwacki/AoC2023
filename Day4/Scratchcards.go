package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sum := checkValues(splitInput(read()))
	fmt.Println(sum)
}

func read() []string {
	file, err := os.Open("C:\\Users\\Mikolaj Karwacki\\Documents\\awesomeProject\\Day4\\input.txt")
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

func splitInput(input []string) [2][]string {
	var firstSplit []string

	for _, v := range input {
		temp := strings.Split(v, ": ")
		firstSplit = append(firstSplit, temp[1])
	}

	var finalRes [2][]string

	for _, v := range firstSplit {
		temp := strings.Split(v, " | ")
		finalRes[0] = append(finalRes[0], temp[0])
		finalRes[1] = append(finalRes[1], temp[1])
	}

	return finalRes
}

func checkValues(input [2][]string) int {
	sum := 0
	winsOnCard := make(map[int]int)
	timesCardIsUsed := make(map[int]int)

	for i, v := range input[0] {
		winning := strings.Fields(v)
		guesses := strings.Fields(input[1][i])
		timesCardIsUsed[i]++
		wins := 0
		for _, v := range guesses {
			if contains(winning, v) {
				wins++
			}
		}
		winsOnCard[i] = wins
		for j := 1; j <= wins; j++ {
			timesCardIsUsed[i+j] += timesCardIsUsed[i]
		}
	}

	for key := range winsOnCard {
		sum += timesCardIsUsed[key]
	}
	return sum
}

func contains(stack []string, needle string) bool {
	for _, v := range stack {
		if v == needle {
			return true
		}
	}
	return false
}

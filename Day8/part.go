package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := read()
	directions, input := getDirections(input)
	elements := getElements(input)
	fmt.Println(len(directions))
	fmt.Println(len(elements))
	res := calc(elements, directions)
	fmt.Println(res)
}

func calc(element map[string][2]string, directions string) int {
	steps := 0
	iter := 0
	keys := createKeys(element)
	fmt.Println(len(keys))
	for {
		for i, key := range keys {
			if iter >= len(directions) {
				iter = 0
			}
			var idx int
			if directions[iter] == 'L' {
				idx = 0
			} else {
				idx = 1
			}
			keys[i] = element[key][idx]
		}
		finished := areAllFinished(keys)

		if finished {
			break
		}

		steps++
		iter++
	}
	return steps
}

func areAllFinished(keys []string) bool {
	result := true
	for _, key := range keys {
		if key[2] != 'Z' {
			result = false
		}
	}
	return result
}

func createKeys(element map[string][2]string) []string {
	var result []string
	for k := range element {
		if k[2] == 'A' {
			result = append(result, k)
		}
	}
	return result
}

func getElements(input []string) map[string][2]string {
	result := make(map[string][2]string)

	for _, v := range input {
		temp := strings.Split(v, " = ")
		temp[1] = strings.ReplaceAll(temp[1], "(", "")
		temp[1] = strings.ReplaceAll(temp[1], ")", "")
		tuple := strings.Split(temp[1], ", ")
		toAdd := [2]string{tuple[0], tuple[1]}
		result[temp[0]] = toAdd
	}

	return result
}

func getDirections(input []string) (string, []string) {
	result := input[0]

	input = input[2:]

	return result, input
}

func read() []string {
	file, err := os.Open("C:\\Users\\Mikolaj Karwacki\\Documents\\awesomeProject\\Day8\\input.txt")
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

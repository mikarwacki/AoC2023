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
	seq := getSequences(input)
	nums := calc(seq)
	sum := sum(nums)
	fmt.Println(sum)
}

func sum(newNums []int) int {
	sum := 0
	for _, v := range newNums {
		sum += v
	}
	return sum
}

func calc(seqs [][]int) []int {
	var result []int
	for _, seq := range seqs {
		num := findNext(seq)
		result = append(result, num)
	}
	return result
}

func findNext(seq []int) int {
	if checkZeros(seq) {
		return 0
	}

	var rec []int
	for i := 0; i < len(seq)-1; i++ {
		num := seq[i+1] - seq[i]
		fmt.Println(num)
		rec = append(rec, num)
	}

	return seq[len(seq)-1] + findNext(rec)
}

func checkZeros(seq []int) bool {
	cond := true
	for _, v := range seq {
		fmt.Println(v)
		if v != 0 {
			cond = false
		}
	}
	return cond
}

func getSequences(input []string) [][]int {
	var result [][]int
	for _, v := range input {
		var toAdd []int
		temp := strings.Fields(v)
		for _, v := range temp {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
			}
			toAdd = append(toAdd, num)
		}
		result = append(result, toAdd)
	}
	return result
}

func read() []string {
	file, err := os.Open("C:\\Users\\Mikolaj Karwacki\\Documents\\awesomeProject\\Day9\\input.txt")
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

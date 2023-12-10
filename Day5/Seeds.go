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
	seeds, input := findSeeds(input)
	soilMap, input := createMaps(input)
	fertMap, input := createMaps(input)
	waterMap, input := createMaps(input)
	lightMap, input := createMaps(input)
	tempMap, input := createMaps(input)
	humMap, input := createMaps(input)
	locationMap, input := createMaps(input)
	maps := [7][][3]int{soilMap, fertMap, waterMap, lightMap, tempMap, humMap, locationMap}
	min := calculateLocations(maps, seeds)
	fmt.Println(min)
}

func read() []string {
	file, err := os.Open("C:\\Users\\Mikolaj Karwacki\\Documents\\awesomeProject\\Day5\\input.txt")
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

func calculateLocations(maps [7][][3]int, seeds []int) int {

	for _, v := range maps {
		flags := createSeedFlags(len(seeds))
		for _, ranges := range v {
			low := ranges[1]
			high := ranges[1] + ranges[2]
			diff := ranges[0] - ranges[1]
			for i, seed := range seeds {
				if seed >= low && seed < high && !flags[i] {
					seeds[i] = seed + diff
					flags[i] = true
				}
			}
		}
	}
	min := seeds[0]

	for _, loc := range seeds {
		if loc < min {
			min = loc
		}
	}
	return min
}

func createSeedFlags(size int) []bool {
	var result []bool
	for i := 0; i < size; i++ {
		result = append(result, false)
	}
	return result
}

func findSeeds(input []string) ([]int, []string) {
	temp := strings.Split(input[0], ": ")
	res := strings.Fields(temp[1])
	var tempSeeds []int
	var seeds []int

	for i, v := range input {
		if v == "seed-to-soil map:" {
			input = input[i+1:]
			break
		}
	}

	for _, s := range res {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
			continue
		}
		tempSeeds = append(tempSeeds, num)
	}

	sum := 0
	for i := 1; i < len(tempSeeds); i = i + 2 {
		//for j := 0; j < tempSeeds[i+1]; j++ {
		//	seeds = append(seeds, tempSeeds[i]+j)
		//}

	}
	fmt.Println(sum)
	return seeds, input
}

func createMaps(input []string) ([][3]int, []string) {
	var soilMap [][]string
	var result [][3]int

	for i, v := range input {
		if len(v) == 0 {
			input = input[i+2:]
			break
		}
		soilMap = append(soilMap, strings.Fields(v))
	}
	for _, row := range soilMap {
		var values [3]int
		for j, str := range row {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println(err)
				continue
			}
			values[j] = num
		}
		result = append(result, values)
	}
	return result, input
}

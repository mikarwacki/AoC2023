package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var card = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 0,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'1': 1,
}

type Game struct {
	hand  string
	price int
}

func main() {
	input := read()
	hands := createHands(input)
	pows := createListsOfGames(hands)
	res := calc(hands, pows)
	fmt.Println(res)
}

func calc(input []Game, pows []Game) int {
	res := 0

	for i, _ := range pows {
		for j := 0; j < len(pows)-i-1; j++ {
			if compareHands(pows[j], pows[j+1]) > 0 {
				temp := pows[j]
				pows[j] = pows[j+1]
				pows[j+1] = temp
				temp2 := input[j]
				input[j] = input[j+1]
				input[j+1] = temp2
			}
		}
	}

	for i, _ := range pows {
		num := input[i].price * (i + 1)
		res += num
	}

	return res
}

func compareHands(h1 Game, h2 Game) int {
	if h1.price > h2.price {
		return 1
	}
	if h1.price < h2.price {
		return -1
	}
	if h1.price == h2.price {
		for i := 0; i < len(h1.hand); i++ {
			if card[h1.hand[i]] > card[h2.hand[i]] {
				return 1
			} else if card[h1.hand[i]] < card[h2.hand[i]] {
				return -1
			}
		}
	}
	return 0
}

func createListsOfGames(input []Game) []Game {
	var result []Game

	for _, v := range input {
		pow := calcHandPow(v.hand)
		result = append(result, Game{v.hand, pow})
	}

	return result
}

func calcHandPow(hand string) int {
	occurance := make(map[byte]int)

	for i := 0; i < len(hand); i++ {
		card := hand[i]
		occurance[card]++
	}

	// for i := 0; i < occurance['J']; i++ {
	//  for key := range occurance {
	//      if key != 'J' {
	//          occurance[key]++
	//      }
	//  }
	// }

	// delete(occurance, 'J')

	max := 0
	pairs := 0

	for _, v := range occurance {
		if v == 2 {
			pairs++
		}
		if v > max {
			max = v
		}
	}

	switch max {
	case 1:
		return 0
	case 3:
		if pairs == 1 {
			return 4
		}
		return 3
	case 4:
		return 5
	case 5:
		return 6
	default:
		return pairs
	}
}

func createHands(input []string) []Game {
	var result []Game
	for _, v := range input {
		temp := strings.Fields(v)
		num, err := strconv.Atoi(temp[1])
		if err != nil {
			continue
		}
		result = append(result, Game{temp[0], num})
	}
	return result
}

func read() []string {
	file, err := os.Open("C:\\Users\\Mikolaj Karwacki\\Documents\\awesomeProject\\Day7\\input.txt")
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

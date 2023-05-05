package main

import (
	"fmt"
	"strings"
)

// rock paper scissors A, B, C as rock, paper, scissors for player 2 and X, Y, Z for player 1, return 6 if player 2 win, 0 if lose and 1 if draw
func win_map() map[string]string {
	win_map := make(map[string]string)
	win_map["A"] = "Z"
	win_map["B"] = "X"
	win_map["C"] = "Y"

	return win_map
}

func equal_map() map[string]string {
	equal_map := make(map[string]string)
	equal_map["A"] = "X"
	equal_map["B"] = "Y"
	equal_map["C"] = "Z"

	return equal_map
}

func lose_map() map[string]string {
	lose_map := make(map[string]string)
	lose_map["A"] = "Y"
	lose_map["B"] = "Z"
	lose_map["C"] = "X"

	return lose_map
}

func input_map() map[string]int {
	input_map := make(map[string]int)
	input_map["X"] = 1
	input_map["Y"] = 2
	input_map["Z"] = 3

	return input_map
}

func nwin_map() map[string]int {
	nwin_map := make(map[string]int)
	nwin_map["X"] = 0
	nwin_map["Y"] = 3
	nwin_map["Z"] = 6

	return nwin_map
}

func RockPaperScissors(player1 string, player2 string) int {
	if equal_map()[player1] == player2 {
		return 3
	}

	if win_map()[player1] == player2 {
		return 0
	}

	return 6
}

func RockPaperScissors2(player2 string) int {
	return nwin_map()[player2]
}

func DayTwo() {
	lines := ReadFile("input_2")

	var total_score int = 0

	for _, line := range lines {
		words := strings.Split(line, " ")
		total_score += RockPaperScissors(words[0], words[1]) + input_map()[words[1]]
	}

	fmt.Printf("1. The total score is %d\n", total_score)
}

func DayTwo2() {
	lines := ReadFile("input_2")

	var total_score int = 0

	for _, line := range lines {
		words := strings.Split(line, " ")
		var input_value int = 0
		if words[1] == "X" {
			input_value = input_map()[win_map()[words[0]]]
		} else if words[1] == "Y" {
			input_value = input_map()[equal_map()[words[0]]]
		} else if words[1] == "Z" {
			input_value = input_map()[lose_map()[words[0]]]
		}

		total_score += RockPaperScissors2(words[1]) + input_value
	}

	fmt.Printf("2. The total score is %d\n", total_score)
}

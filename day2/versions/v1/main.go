package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	var text = string(bytes)

	var lines = strings.Split(text, "\n")

	for _, v := range lines {
		var splits = strings.Split(v, ": ")
		game := splits[0]
		rounds := splits[1]

		// Get game number
		var game_number_str = strings.Split(game, " ")[1]
		game_number, err := strconv.Atoi(game_number_str)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Create array of rounds in game
		rounds_array := strings.Split(rounds, "; ")

		round1_raw := strings.Split(rounds_array[0], ", ")
		round2_raw := strings.Split(rounds_array[1], ", ")
		round3_raw := strings.Split(rounds_array[2], ", ")

		// Create arrays holding number of balls of each color pulled that round (red, blue, green)
		var round1 = make([]int, 3)
		var round2 = make([]int, 3)
		var round3 = make([]int, 3)

		for _, v := range round1_raw {
			colors := strings.Split(v, " ")
			amount, err := strconv.Atoi(colors[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			switch colors[1] {
			case "red":
				round1[0] = amount
			case "blue":
				round1[1] = amount
			case "green":
				round1[2] = amount
			default:
				fmt.Println("Invalid color")
				return
			}
		}
		for _, v := range round2_raw {
			colors := strings.Split(v, " ")
			amount, err := strconv.Atoi(colors[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			switch colors[1] {
			case "red":
				round2[0] = amount
			case "blue":
				round2[1] = amount
			case "green":
				round2[2] = amount
			default:
				fmt.Println("Invalid color")
				return
			}
		}
		for _, v := range round3_raw {
			colors := strings.Split(v, " ")
			amount, err := strconv.Atoi(colors[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			switch colors[1] {
			case "red":
				round3[0] = amount
			case "blue":
				round3[1] = amount
			case "green":
				round3[2] = amount
			default:
				fmt.Println("Invalid color")
				return
			}
		}
		fmt.Printf("Game %d: Round 1: %v, Round 2: %v, Round 3: %v\n", game_number, round1, round2, round3)
	}
}

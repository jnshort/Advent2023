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

	var sum = 0

	var text = string(bytes)

	var lines = strings.Split(text, "\n")

	for _, v := range lines {
		var splits = strings.Split(v, ": ")
		game := splits[0]
		rounds_str := splits[1]

		// Get game number
		var game_number_str = strings.Split(game, " ")[1]
		game_number, err := strconv.Atoi(game_number_str)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(game_number)

		// Create array of rounds in game
		rounds_array := strings.Split(rounds_str, "; ")

		// Create arrays holding number of balls of each color pulled that round (red, blue, green)
		rounds := make([][]int, len(rounds_array))

		var valid_game = true

		for i := range rounds {
			rounds[i] = make([]int, 3)
			round_raw := strings.Split(rounds_array[i], ", ")

			for _, v := range round_raw {
				// Your logic here
				colors := strings.Split(v, " ")
				amount, err := strconv.Atoi(colors[0])
				if err != nil {
					fmt.Println(err)
					return
				}
				switch colors[1] {
				case "red":
					if amount > 12 {
						valid_game = false
						break
					}
					rounds[i][0] = amount
				case "blue":
					if amount > 14 {
						valid_game = false
						break
					}
					rounds[i][1] = amount
				case "green":
					if amount > 13 {
						valid_game = false
						break
					}
					rounds[i][2] = amount
				default:
					fmt.Println("Invalid color")
					return
				}
				fmt.Println(rounds[i])
			}

		}
		if valid_game == true {
			sum += game_number
		}
		fmt.Println(valid_game)

		valid_game = true
	}
	fmt.Println(sum)
}

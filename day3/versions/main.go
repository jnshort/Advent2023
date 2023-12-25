package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day3/input.txt")
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

	var engine = string(bytes)

	var lines = strings.Split(engine, "\n")

	var runes = make([][]rune, len(lines))
	// for line in engine, create a list of runes in the line
	for i, v := range lines {
		runes[i] = []rune(v)
	}

	valid_numbers := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	for i := range runes {
		for j := range runes[i] {
			for _, v := range valid_numbers {
				if runes[i][j] == v {
					var number_string = string(runes[i][j])
					var k = j + 1
					for k < len(runes[i]) && contains(valid_numbers, runes[i][k]) {
						number_string += string(runes[i][k])
						k++
					}
					var number, err = strconv.Atoi(number_string)
					if err != nil {
						fmt.Println(err)
						return
					}
					var length = len(number_string)
					fmt.Println(number, length)

				}
			}
		}
	}
}

func contains(s []rune, e rune) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

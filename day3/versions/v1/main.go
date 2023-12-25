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

	var sum = 0

	var engine = string(bytes)

	var lines = strings.Split(engine, "\n")

	var runes = make([][]rune, len(lines))
	// for line in engine, create a list of runes in the line
	for i, v := range lines {
		runes[i] = []rune(v)
	}

	valid_numbers := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	var checked = make([][]bool, len(runes))
	for i := range checked {
		checked[i] = make([]bool, len(runes[i]))
	}

	for i := range runes {
		for j := range runes[i] {
			for _, v := range valid_numbers {
				if runes[i][j] == v && !checked[i][j] {
					checked[i][j] = true
					var number_string = []rune{runes[i][j]}
					var k = j + 1
					for k < len(runes[i]) && contains(valid_numbers, runes[i][k]) {
						number_string = append(number_string, runes[i][k])
						checked[i][k] = true
						k++
					}
					var number_str = string(number_string)
					var number, err = strconv.Atoi(number_str)
					if err != nil {
						fmt.Println(err)
						return
					}
					var length = len(number_string)

					valid_part := valid_part_number(i, j, length, runes)
					if valid_part {
						sum += number
						fmt.Println(number, sum)
					}
				}
			}
		}
	}
	fmt.Println(sum)
}

func contains(s []rune, e rune) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// takes in i, j, and length of number
// checks all surounding indexes and returns true if a symbol is found that is not '.' or a number
func valid_part_number(i int, j int, length int, runes [][]rune) bool {
	valid_numbers := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	var indexes = []rune{}

	top_left := []int{i - 1, j - 1}
	bottom_right := []int{i + 1, j + length + 1}

	for k := max(0, top_left[0]); k <= min(len(runes)-1, bottom_right[0]); k++ {
		for l := max(0, top_left[1]); l <= min(len(runes[k])-1, bottom_right[1]); l++ {
			if runes[k][l] != '.' && !contains(valid_numbers, runes[k][l]) {
				indexes = append(indexes, runes[k][l])
			}
		}
	}

	var valid = false
	if len(indexes) != 0 {
		valid = true
	}
	return valid
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

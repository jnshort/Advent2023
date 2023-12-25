package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Part struct {
	Number int
	i      int
	j      int
	length int
}

func (p Part) span() [][]int {
	var span = [][]int{}
	for k := 0; k < p.length; k++ {
		span = append(span, []int{p.i, p.j + k})
	}
	return span
}

func (p Part) equals(other Part) bool {
	return p.i == other.i && p.j == other.j
}

func contains_part(s []Part, e Part) bool {
	for _, a := range s {
		if a.equals(e) {
			return true
		}
	}
	return false
}

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

	parts := []Part{}

	for i := range runes {
		for j := range runes[i] {
			if contains(valid_numbers, runes[i][j]) && !checked[i][j] {
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
					parts = append(parts, Part{number, i, j, length})
				}
			}
		}
	}
	fmt.Println("Total: ", sum)

	var gear_ratio_sum = 0
	for i := range runes {
		for j := range runes[i] {
			if runes[i][j] == '*' {
				gear_ratio, err := gear_ratio(i, j, parts, runes)
				if err != nil {
				} else {
					gear_ratio_sum += gear_ratio
				}
			}
		}
	}
	fmt.Println("Gear Ratio: ", gear_ratio_sum)
}

func gear_ratio(i int, j int, parts []Part, runes [][]rune) (int, error) {
	var err = error(nil)
	var gear_ratio int

	// find unique parts that have at least 1 cell adjacent to the gear
	var connected_parts = []Part{}

	//for part in part list
	for _, part := range parts {
		// for 8 cells sourounding the gear at i, j
		for k := max(0, i-1); k <= min(len(runes)-1, i+1); k++ {
			for l := max(0, j-1); l <= min(len(runes[k])-1, j+1); l++ {
				// if part.span() contains the cell
				for _, span := range part.span() {
					if span[0] == k && span[1] == l {
						// if part is not in connected_parts
						if !contains_part(connected_parts, part) {
							// add part to connected_parts
							connected_parts = append(connected_parts, part)
						}
					}
				}

			}
		}

	}

	if len(connected_parts) == 2 {
		gear_ratio = connected_parts[0].Number * connected_parts[1].Number
	} else {
		gear_ratio = 0
		err = fmt.Errorf("Gear at position (%d, %d) is not connected to 2 parts", i, j)
	}
	return gear_ratio, err
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
	valid_numbers := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '\r'}
	var indexes = []rune{}

	top_left := []int{i - 1, j - 1}
	bottom_right := []int{i + 1, j + length}

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

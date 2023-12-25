package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
)

func main() {
	file, err := os.Open("go_day_1/input.txt")
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
	text = strings.Replace(text, "one", "1", -1)
	text = strings.Replace(text, "two", "2", -1)
	text = strings.Replace(text, "three", "3", -1)
	text = strings.Replace(text, "four", "4", -1)
	text = strings.Replace(text, "five", "5", -1)
	text = strings.Replace(text, "six", "6", -1)
	text = strings.Replace(text, "seven", "7", -1)
	text = strings.Replace(text, "eight", "8", -1)
	text = strings.Replace(text, "nine", "9", -1)

	// split text by new line
	var lines = strings.Split(text, "\n")

	var sum = 0

	// for each line find first and last number

	for _, line := range lines {
		var line_first = ""
		var line_last = ""
		var i int = 0
		for line_first == "" {
			if line[i] >= '0' && line[i] <= '9' {
				line_first = string(line[i])
			}else{
				i++
			}
		}
		var j int = int(len(line) - 1)
		for line_last == "" {
			if line[j] >= '0' && line[j] <= '9' {
				line_last = string(line[j])
			}else{
			j--
			}
		}
		var line_number_str = line_first + line_last
		var line_number_int, err = strconv.Atoi(line_number_str)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(line_number_int)
		sum += line_number_int
	}
	fmt.Println(sum)
}
package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day1/input.txt")
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

	// split text by new line
	var lines = strings.Split(text, "\n")

	var sum = 0

	// for each line find first and last number as a single digit or string of characters
	var validSubStrings = [18]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, v1 := range lines {
		var firstIndexArray = [18]int{}
		var lastIndexArray = [18]int{}
		for i := range firstIndexArray {
			firstIndexArray[i] = -1
			lastIndexArray[i] = -1
		}

		for i, v2 := range validSubStrings {
			var index1 = strings.Index(v1, v2)
			if index1 != -1 {
				firstIndexArray[i] = index1
			}
			var index2 = strings.LastIndex(v1, v2)
			if index2 != -1 {
				lastIndexArray[i] = index2
			}
		}
		var first_ind = 1000000
		var first_val = 1000000
		var last_ind = -1
		var last_val = -1
		for i, v := range firstIndexArray {
			if v != -1 {
				if v < first_val {
					first_ind = i
					first_val = v
				}
			}
		}
		for i, v := range lastIndexArray {
			if v != -1 {
				if v > last_val {
					last_ind = i
					last_val = v
				}
			}
		}
		var first = (first_ind % 9) + 1
		var last = (last_ind % 9) + 1

		firstStr := strconv.Itoa(first)
		lastStr := strconv.Itoa(last)

		concatStr := firstStr + lastStr

		fullNumber, err := strconv.Atoi(concatStr)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%d \n", fullNumber)
		sum += fullNumber
	}

	fmt.Println(sum)
}

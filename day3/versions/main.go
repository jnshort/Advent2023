package main

import (
	"fmt"
	"io"
	"os"
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

	for i := range runes {
		for j := range runes[i] {
			fmt.Print(string(runes[i][j]))
		}
		fmt.Println()
	}
}

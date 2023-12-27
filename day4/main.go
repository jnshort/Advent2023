package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	card_id         int
	winning_numbers []int
	chosen_numbers  []int
	copies          int
}

func (c Card) String() string {
	return fmt.Sprintf("Winning numbers: %v\nChosen numbers: %v", c.winning_numbers, c.chosen_numbers)
}

func (c Card) Points() int {
	//return 1 for a single matching number, and double for each additional matching number
	matches := 0
	points := 0
	var found []bool = make([]bool, len(c.winning_numbers))
	for index, winning_num := range c.winning_numbers {
		if contains(c.chosen_numbers, winning_num) && !found[index] {
			matches++
			found[index] = true
		}
	}
	switch matches {
	case 0:
		points = 0
	case 1:
		points = 1
	default:
		points = 1 << (matches - 1)
	}
	return points
}

func (c Card) Matches() int {
	var matches []int
	for _, winning_num := range c.winning_numbers {
		if contains(c.chosen_numbers, winning_num) {
			matches = append(matches, winning_num)
		}
	}
	return len(matches)
}

func process_card_list(cards []Card) int {
	var scratchcards = 0
	for i, card := range cards {
		winning_cards := card.Matches()
		for j := i+1; j < i+winning_cards+1; j++ {
			if j > len(cards) {
				break
			}
			cards[j].copies += card.copies
		}
		scratchcards += card.copies
	}
	return scratchcards
}

func main() {
	var file, err = os.Open("day4/input.txt")
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
	contents := string(bytes)

	lines := strings.Split(contents, "\n")
	var pointsSum = 0
	var card_id = 0
	cards := make([]Card, 0)

	for _, line := range lines {
		numbers := strings.Split(line, ": ")[1]
		raw_lists := strings.Split(numbers, " | ")
		winning_numbers, chosen_numbers := normalize_string_lists(raw_lists[0]), normalize_string_lists(raw_lists[1])
		var card Card
		card, err = create_card(card_id, winning_numbers, chosen_numbers)
		if err != nil {
			fmt.Println(err)
			return
		}
		pointsSum += card.Points()
		cards = append(cards, card)
		card_id++
	}
	fmt.Println("Sum of points: ", pointsSum)

	fmt.Println("number of cards: ", len(cards))

	fmt.Println("Number of scratchcards: ", process_card_list(cards))

}

func normalize_string_lists(number_list string) string {
	runeSlice := []rune(number_list)
	for i := 0; i < len(runeSlice); i += 3 {
		if runeSlice[i] == ' ' {
			runeSlice[i] = '0'
		}
	}
	return string(runeSlice)
}

func create_card(card_id int, winning_numbers string, chosen_numbers string) (Card, error) {
	winning_nums_as_strings := strings.Split(winning_numbers, " ")
	chosen_nums_as_strings := strings.Split(chosen_numbers, " ")
	var winning = make([]int, len(winning_nums_as_strings))
	var chosen = make([]int, len(chosen_nums_as_strings))
	var err = error(nil)
	for i, num := range winning_nums_as_strings {
		winning[i], err = strconv.Atoi(num)
		if err != nil {
			fmt.Println(err)
			return Card{}, err
		}
	}
	for i, num := range chosen_nums_as_strings {
		chosen[i], err = strconv.Atoi(num)
	}
	return Card{card_id, winning, chosen, 1}, err
}

func contains(slice []int, num int) bool {
	for _, n := range slice {
		if n == num {
			return true
		}
	}
	return false
}

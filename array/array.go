package main

import "fmt"

// array is of fixed length
//slice can grow and shrink, should be of same type

func main() {
	cards := []string{"Ace of Diamnods", newCard()}
	cards = append(cards, "Six of Spades")
	for i, card := range cards {
		fmt.Println(i, card)
	}
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamond"
}

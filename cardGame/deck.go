package main

import "fmt"

type deck []string

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Diamond", "Hearts", "Spades", "club"}
	cardValues := []string{"Ace", "two", "three", "four"}

	for _, carcardSuit := range cardSuits {
		for _, cardcardValue := range cardValues {
			cards = append(cards, cardcardValue+" of "+carcardSuit)
		}
	}
	return cards
}

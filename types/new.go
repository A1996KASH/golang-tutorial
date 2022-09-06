package main

import "fmt"

type cards []string

func (c cards) print() {
	for i, card := range c {
		fmt.Println(card, i)
	}
}

func newCards() cards {
	return cards{"tets", "Akash"}
}

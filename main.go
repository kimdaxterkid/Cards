package main

import "fmt"

func main() {
	//var card = "Ace of Spades"
	cards := []string{"Ace of Diamonds", newCard()}

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}

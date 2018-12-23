package main

func main() {
	cards := newDeck()
	cards.shuffle()
	//hand, remainCards := deal(cards, 5)
	//cards.saveToFile("MyCards")
	//cards := newDeckFromFile("My Cards")
	cards.print()
}

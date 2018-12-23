package main

func main() {
	//cards := newDeck()
	//hand, remainCards := deal(cards, 5)
	//cards.saveToFile("MyCards")
	cards := newDeckFromFile("MyCards")
	cards.print()
}

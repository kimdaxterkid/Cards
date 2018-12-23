package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// all deck type varibale could use print() function
// d is abbreviation for deck, usually use the first letter of the type, this case - 'd'
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// there is the parenthesis for two returns
// but in the return syntax, there is no need of the parenthesis
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// strings package has the easy function for me - join
func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

// io/ioutil package to save file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a cal to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	stringSlice := strings.Split(string(byteSlice), ",")
	return deck(stringSlice)
}

// rand and time to be used together to create real random number
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

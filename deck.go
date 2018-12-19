package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// all deck type varibale could use print() function
// d is abbreviation for deck, usually use the first letter of the type, this case - 'd'
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

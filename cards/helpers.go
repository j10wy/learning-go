package main

func newCard() string {
	return "Joker"
}

func newDeck() deck {
	cards := deck{}

	// A Slice is an array that can grow and shrink.
	// An Array in Go has a fixed length
	// ex Array: [2]string{"one", "two"}
	// ex Slice: []string{"one", "two", "three", "four"}
	cardSuites := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			// The append function does not modify the exisiting slice,
			// it returns a new slice.
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

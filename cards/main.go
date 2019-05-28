package main

func main() {

	hand, cards := deal(newDeck(), 3)

	hand.print()
	cards.print()

}

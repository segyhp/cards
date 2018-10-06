package main

func main() {

	// cards := newDeck()

	// fmt.Println(cards.toString())

	// cards.saveToFile("myCards")

	cards := newDeckFromFile("myCarsds")
	cards.print()
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
}

package main

func main() {
	cards := newDeck()
	cards.printCards()

	hand, remainingCards := deal(cards, 5)
	hand.printCards()
	remainingCards.printCards()

}

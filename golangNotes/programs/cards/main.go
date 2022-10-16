package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.printCards()

	//cardsFromFile := newDeckFromFile("my_cards")
	//cardsFromFile.printCards()

	hand, remainingCards := deal(cards, 5)
	hand.printCards()
	remainingCards.printCards()

}

package main

func main() {
	playingCards := newDeckFromFile("playing_cards")
	// playingCards.shuffle()
	playingCards.print()

	hand, _ := deal(playingCards, 2)
	countHand(hand)

}

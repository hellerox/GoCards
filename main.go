package main

func main() {
	cards := newDeck()
	cards.saveToFile("godpdf")
	cards2 := newDeckFromFile("godpdf")
	cards2.print()
}

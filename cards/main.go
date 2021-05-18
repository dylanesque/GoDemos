// https://tour.golang.org/basics/5

// Required to run the program, this groups functions, 
// is comprised of all files in the same directory
// Makes an executable package

package main

func main() {
	// More verbose variable declaration: var card string = "Ace of Spades"
	// colon-equals syntax is only used for new variables
	// reassignments are done without the colon
	// card := "Ace of Spades"
	// variables can be initialized outside of function scope, but not assigned values

	// Slices in Go are like arrays, but mutable
	// They can only hold one type of primitive value
	cards := newDeckFromFile("my_cards")
	cards.print()
	// We add new items via appending
	// Notable: this creates a new item, it doesn't mutate the original slice
	// cards = append(cards, "six of spades")

	// Syntax like the below deminstrated is how we deal with multiple return values in Go
	// hand, remainingCards := deal(cards, 5)


}

// While the 'main' package must be explicitly declared, files within the same package can call functions
// defined in other files without needing to manually import them

func newCard() string {
	return "five of diamonds"
}

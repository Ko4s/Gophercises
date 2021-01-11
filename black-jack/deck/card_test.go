package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {

	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Nine, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Club})
	fmt.Println(Card{Rank: King, Suit: Diamond})
	fmt.Println(Card{Rank: Queen, Suit: Diamond})
	fmt.Println(Card{Rank: Ace, Suit: Spade})

	// Output:
	// Ace of Hearts
	// Nine of Hearts
	// Two of Clubs
	// King of Diamonds
	// Queen of Diamonds
	// Ace of Spades
}

func TestNew(t *testing.T) {
	cards := New()

	//12 ranks * 4 suits
	if len(cards) != 52 {
		t.Error("Wrong number of cards in a new deck.")
	}
}

package deck

import (
	"fmt"
	"sort"
)

//Rank card rank type
//go:generate stringer -type=Rank
type Rank uint

//Carsds constants
const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

//Suit card type
//go:generate stringer -type=Suit
type Suit uint

//Clubs Cards Colors constants
const (
	Club Suit = iota
	Diamond
	Heart
	Spade
	Joker //special case
)

//CardRank splice of cards figures
var ranks = [...]Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

const (
	minRank = Ace
	maxRank = King
)

//CardSuit splice of cards colors
var suits = [...]Suit{Club, Diamond, Heart, Spade}

//Card type
type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %ss", c.Rank, c.Suit)
}

//New creates a deck of cards
func New(opts ...func([]Card) []Card) []Card {

	cards := []Card{}

	for _, suit := range suits {
		for _, rank := range ranks {
			card := Card{Suit: suit, Rank: rank}
			cards = append(cards, card)
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}

	return cards
}

//DefaultSort is a DefaultSort :)
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

//Less a default compare function
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

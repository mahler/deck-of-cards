//go:generate stringer -type=Rank,Suit
package deck

import "fmt"

// Suit holds the four suits of a normal set of playing cards.
type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker // This is a special case...
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

// Rank holds the 13 different card ranks in a set of playing cards.
type Rank uint8

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

const minRank = Ace
const maxRank = King

// Card consists of a suit and a rank
type Card struct {
	Suit
	Rank
}

func New() []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}

	return cards
}

func (c Card) String() string {
	if c.Suit == Joker {
		return "Joker"
	}

	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

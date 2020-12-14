//go:generate stringer -type=Rank,Suit
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

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

func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}

	return cards
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func (c Card) String() string {
	if c.Suit == Joker {
		return "Joker"
	}

	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(cards))
	for i, j := range perm {
		ret[i] = cards[j]
	}
	return cards
}

func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Rank: Rank(i),
				Suit: Joker,
			})
		}
		return cards
	}
}

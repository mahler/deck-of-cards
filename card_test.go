package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Nine, Suit: Club})
	fmt.Println(Card{Rank: Jack, Suit: Diamond})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Nine of Clubs
	// Jack of Diamonds
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	// 13 ranks in 4 suits
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in a new deck")
	}
}
func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	// 13 ranks in 4 suits
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in a new deck")
	}

	expectedCard := Card{Rank: Ace, Suit: Spade}
	if cards[0] != expectedCard {
		t.Error("Expected Ace of Spades as first card, recieved:", cards[0])
	}
}

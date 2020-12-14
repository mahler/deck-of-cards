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

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Error("Expected 3 jokers, received:", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}

	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Expected all twos and threes to be filtered out.")
		}
	}
}

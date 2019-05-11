package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Nine, Suit: Diamond})
	fmt.Println(Card{Rank: Jack, Suit: Club})
	fmt.Println(Card{Suit: Joker})
	// Output:
	// Ace of Hearts
	// Two of Spades
	// Nine of Diamonds
	// Jack of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	ranksTimesSuits := 13 * 4 // 13 ranks * 4 suits
	if len(cards) != ranksTimesSuits {
		t.Error("Wrong number of cards in a new deck.")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	expected := Card{Rank: Ace, Suit: Spade}
	if cards[0] != expected {
		t.Errorf("Expected %s first card.Received: %s", expected, cards[0])
	}
}
func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	expected := Card{Rank: Ace, Suit: Spade}
	if cards[0] != expected {
		t.Errorf("Expected %s first card.Received: %s", expected, cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	expected := 3
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}

	if count != expected {
		t.Errorf("Expected %d jokers.Received: %d", expected, count)
	}

}

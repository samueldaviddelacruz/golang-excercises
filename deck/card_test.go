package deck

import (
	"fmt"
	"math/rand"
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

func TestShuggle(t *testing.T) {
	//make shuffle rand deterministic
	//First call to shuffleRand.Perm(52) should be:
	// [40 35 50 0 44....]
	shuffleRand = rand.New(rand.NewSource(0))
	orig := New()

	first := orig[40]
	second := orig[35]

	cards := New(Shuffle)
	if cards[0] != first {
		t.Errorf("Expected first card to be %s .Received: %s", first, cards[0])

	}
	if cards[1] != second {
		t.Errorf("Expected second card to be %s .Received: %s", second, cards[1])

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

func TestFilter(t *testing.T) {
	predicate := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(predicate))
	for _, card := range cards {
		if card.Rank == Two || card.Rank == Three {
			t.Error("Expected all twos and threes to be filtered out")
		}
	}

}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	ranksTimesSuitsTimesDecks := 13 * 4 * 3 // 13 ranks * 4 suits * 3 decks
	if len(cards) != ranksTimesSuitsTimesDecks {
		t.Errorf("Expected %d cards.Received: %d", ranksTimesSuitsTimesDecks, len(cards))

	}

}

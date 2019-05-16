package main

import (
	"fmt"
	"strings"

	"github.com/samueldaviddelacruz/golang-exercises/deck"
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}
func (h Hand) DealerString() string {
	return h[0].String() + ", **HIDDEN**"
}

func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		if c.Rank == deck.Ace {
			//ace is currently worth 1, and we are changing to be worth 11
			//11 - 1 = 10
			return minScore + 10
		}
	}
	return minScore
}

func (h Hand) MinScore() int {
	score := 0
	for _, c := range h {
		score += min(int(c.Rank), 10)
	}

	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Shuffle(gs GameState) GameState {
	result := clone(gs)
	result.Deck = deck.New(deck.Deck(3), deck.Shuffle)
	return result
}
func Deal(gs GameState) GameState {
	result := clone(gs)
	result.Player = make(Hand, 0, 5)
	result.Dealer = make(Hand, 0, 5)

	var card deck.Card
	for i := 0; i < 2; i++ {
		card, result.Deck = draw(result.Deck)
		result.Player = append(result.Player, card)
		card, result.Deck = draw(result.Deck)
		result.Dealer = append(result.Dealer, card)
	}
	result.State = StatePlayerTurn

	return result
}
func Hit(gs GameState) GameState {
	result := clone(gs)
	hand := result.CurrentPlayer()
	var card deck.Card
	card, result.Deck = draw(result.Deck)
	*hand = append(*hand, card)
	return result
}
func Stand(gs GameState) GameState {
	result := clone(gs)

	result.State++

	return result
}
func EndHand(gs GameState) GameState {
	result := clone(gs)
	pScore, dScore := result.Player.Score(), result.Dealer.Score()
	fmt.Scanln("==FINAL HANDS==")

	fmt.Println("Player:", result.Player, "\nScore:", pScore)
	fmt.Println("Dealer:", result.Dealer, "\nScore:", dScore)

	switch {
	case pScore > 21:
		fmt.Println("You busted")
	case dScore > 21:
		fmt.Println("Dealer busted")
	case pScore > dScore:
		fmt.Println("You win")
	case dScore > pScore:
		fmt.Println("You lose")
	case dScore == pScore:
		fmt.Println("Draw")
	}
	fmt.Println()
	result.Player = nil
	result.Dealer = nil
	return result
}
func main() {
	var gs GameState
	gs = Shuffle(gs)
	for i := 0; i < 10; i++ {
		gs = Deal(gs)

		var input string
		for gs.State == StatePlayerTurn {
			fmt.Println("Player:", gs.Player)
			fmt.Println("Dealer:", gs.Dealer.DealerString())

			fmt.Println("What will you do? (h)it, (s)stand")

			fmt.Scanf("%s\n", &input)
			switch input {
			case "h":
				gs = Hit(gs)
			case "s":
				gs = Stand(gs)
			default:
				fmt.Println("Invalid option:", input)
			}
		}
		for gs.State == StateDealerTurn {
			//if dealer score <= 16, we hit
			//if dealer has soft 17, then we hit.

			if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
				gs = Hit(gs)
			} else {
				gs = Stand(gs)
			}
		}
		gs = EndHand(gs)
	}

}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

type State int8

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

type GameState struct {
	Deck   []deck.Card
	State  State
	Player Hand
	Dealer Hand
}

func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.State {
	case StatePlayerTurn:
		return &gs.Player
	case StateDealerTurn:
		return &gs.Dealer
	default:
		panic("it isn't currently any player's turn")
	}
}

func clone(gs GameState) GameState {
	result := GameState{
		Deck:   make([]deck.Card, len(gs.Deck)),
		State:  gs.State,
		Player: make(Hand, len(gs.Player)),
		Dealer: make(Hand, len(gs.Dealer)),
	}
	copy(result.Deck, gs.Deck)
	copy(result.Player, gs.Player)
	copy(result.Dealer, gs.Dealer)
	return result
}

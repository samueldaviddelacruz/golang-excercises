package blackjack

import (
	"fmt"

	"github.com/samueldaviddelacruz/golang-exercises/deck"
)

type AI interface {
	Results(hands [][]deck.Card, dealer []deck.Card)
	Bet(shuffled bool) int
	Play(hand []deck.Card, dealer deck.Card) Move
}
type dealerAI struct{}

func (ai dealerAI) Play(hand []deck.Card, dealer deck.Card) Move {
	dScore := Score(hand...)
	if dScore <= 16 || (dScore == 17 && Soft(hand...)) {
		return MoveHit
	}
	return MoveStand
}

func (ai dealerAI) Bet(shuffled bool) int {
	return 1
}
func (ai dealerAI) Results(hands [][]deck.Card, dealer []deck.Card) {
	//noop
}
func HumanAI() AI {
	return humanAI{}
}

type humanAI struct{}

func (ai humanAI) Bet(shuffled bool) int {
	if shuffled {
		fmt.Println("The deck was just shuffled")
	}
	fmt.Println("What would you liek to bet?")
	var bet int
	fmt.Scanf("%d\n", &bet)
	return bet
}

func (ai humanAI) Play(hand []deck.Card, dealer deck.Card) Move {
	for {
		fmt.Println("Player:", hand)
		fmt.Println("Dealer:", dealer)

		fmt.Println("What will you do? (h)it, (s)stand, s(p)lit")
		var input string
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		case "p":
			return MoveSplit
		default:
			fmt.Println("Invalid option:", input)
		}
	}

}
func (ai humanAI) Results(hands [][]deck.Card, dealer []deck.Card) {
	fmt.Scanln("==FINAL HANDS==")
	for _, h := range hands {
		fmt.Println(" ", h)
	}
	fmt.Println("Player:", hands)
	fmt.Println("Dealer:", dealer)
}

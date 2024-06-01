package gocards

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
)

type Pile []Card

func NewPile() *Pile {
	return &Pile{}
}

func (p Pile) String() string {
	var sb strings.Builder
	for _, card := range p {
		sb.WriteString(fmt.Sprintf(" %s", card))
	}
	return sb.String()
}

func (pile *Pile) AddCard(card Card, idx int) {
	if idx == -1 {
		// log.Printf("idx is %d", idx)
		idx = len(*pile)
	}
	if len(*pile) == idx {
		*pile = append(*pile, card)
		// log.Print("idx is ", pile)
	} else {
		*pile = append((*pile)[:idx+1], (*pile)[idx:]...)
		(*pile)[idx] = card
	}
}

func (pile *Pile) AddCardToEnd(card Card) {
	pile.AddCard(card, -1)
}

func (pile *Pile) AddCardToStart(card Card) {
	pile.AddCard(card, 0)
}

func (pile *Pile) RemoveCard(idx int) (bool, Card) {
	if idx < len(*pile) && idx >= 0 {
		removedCard := (*pile)[idx]
		*pile = append((*pile)[:idx], (*pile)[idx+1:]...)
		return true, removedCard
	}
	return false, Card{}
}

func (pile *Pile) RemoveCards(idxs []int) Pile {
	removedPile := Pile{}
	for _, idx := range idxs {
		removed, card := pile.RemoveCard(idx)
		if removed {
			removedPile.AddCardToEnd(card)
		}
	}
	return removedPile
}

func (pile Pile) SearchCard(c Card) (bool, int) {
	for i, card := range pile {
		if card.Eq(c) {
			return true, i
		}
	}
	return false, -1
}

func (pile Pile) Shuffle() {
	rand.Shuffle(len(pile), func(i, j int) {
		pile[i], pile[j] = pile[j], pile[i]
	})
}

func (pile Pile) ShowCard(idx int) {
	pile[idx].Show = true
}

func (pile Pile) ShowCards(idxs []int) {
	for _, idx := range idxs {
		pile.ShowCard(idx)
	}
}

func (pile Pile) ShowAllCards() {
	for idx := range pile {
		pile.ShowCard(idx)
	}
}

func (a Pile) Len() int      { return len(a) }
func (a Pile) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Pile) Less(i, j int) bool {
	return a[i].Suit < a[j].Suit && a[i].Number < a[j].Number
}

func GetDeck(joker int) Pile {
	pile := Pile{}
	for _, suit := range [...]Suit{SPADES, HEARTS, DIAMONDS, CLUBS} {
		for _, num := range [...]Number{
			ACE, TWO, THREE, FOUR, FIVE, SIX, SEVEN,
			EIGHT, NINE, TEN, JACK, QUEEN, KING,
		} {
			pile.AddCardToEnd(Card{Suit: suit, Number: num})
		}
	}
	i := 0
	for i < joker {
		pile.AddCardToEnd(Card{Suit: NOSUIT, Number: JOKER})
		i += 1
	}
	return pile
}

func GetDecks(n int, joker int) Pile {
	log.Printf("Getting %d decks with %d jokers", n, joker)
	deck := Pile{}
	i := 0
	for i < n {
		deck = append(deck, GetDeck(0)...)
		i += 1
	}
	i = 0
	for i < joker {
		deck.AddCardToEnd(Card{Suit: NOSUIT, Number: JOKER})
		i += 1
	}
	// fmt.Println(deck)
	log.Printf("Deck - %s", deck)
	return deck
}

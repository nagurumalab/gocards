package gocards

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
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

// TODO: Rewrite to handle the whole range of reverse index
func (pile *Pile) PutCard(card Card, idx int) {
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

func (pile *Pile) PutCardEnd(card Card) {
	pile.PutCard(card, -1)
}

func (pile *Pile) PutCardStart(card Card) {
	pile.PutCard(card, 0)
}

var ErrRemoveEmptyPile = errors.New("can't remove a card from empty pile")

// TODO: Rewrite to handle the whole range of reverse index
func (pile *Pile) GetCard(idx int) (Card, error) {
	if idx == -1 {
		idx = len(*pile)
	}
	if idx < len(*pile) && idx >= 0 {
		removedCard := (*pile)[idx]
		*pile = append((*pile)[:idx], (*pile)[idx+1:]...)
		return removedCard, nil
	}
	return Card{}, ErrRemoveEmptyPile
}

func (pile *Pile) GetCardStart() (Card, error) {
	return pile.GetCard(0)
}

func (pile *Pile) GetCardEnd() (Card, error) {
	return pile.GetCard(-1)
}

// TODO: Rewite this whole function better
func (pile *Pile) GetCards(idxs []int) Pile {
	removedPile := Pile{}
	// FIXME: BUG. Once a card is removed, the idxs change places and becomes invalid

	for _, idx := range idxs {
		card, err := pile.GetCard(idx)
		if err != nil {
			if errors.Is(err, ErrRemoveEmptyPile) {
				log.Info().Msgf("No more cards in the pile - %d", removedPile.Len())
			}
			break
		}
		removedPile.PutCardEnd(card)
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

func (pile Pile) HideCard(idx int) {
	pile[idx].Show = false
}

func (pile Pile) HideAllCards() {
	for idx := range pile {
		pile.HideCard(idx)
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
			pile.PutCardEnd(Card{Suit: suit, Number: num})
		}
	}
	i := 0
	for i < joker {
		pile.PutCardEnd(Card{Suit: NOSUITE, Number: JOKER})
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
		deck.PutCardEnd(Card{Suit: NOSUITE, Number: JOKER})
		i += 1
	}
	// fmt.Println(deck)
	log.Printf("Deck - %s", deck)
	return deck
}

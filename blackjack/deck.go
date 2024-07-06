package blackjack

import (
	pbblackjack "github.com/anmho/blackjack/gen/proto/blackjack"
	"math/rand"
)

const (
	CardsInDeck = 52
)

type Deck struct {
	Cards []Card
	rand  *rand.Rand
}

var (
	Ranks = []pbblackjack.Rank{
		pbblackjack.Rank_ACE,
		pbblackjack.Rank_ONE,
		pbblackjack.Rank_TWO,
		pbblackjack.Rank_THREE,
		pbblackjack.Rank_FOUR,
		pbblackjack.Rank_FIVE,
		pbblackjack.Rank_SIX,
		pbblackjack.Rank_SEVEN,
		pbblackjack.Rank_EIGHT,
		pbblackjack.Rank_NINE,
		pbblackjack.Rank_JACK,
		pbblackjack.Rank_QUEEN,
		pbblackjack.Rank_KING,
	}

	Suits = []pbblackjack.Suit{
		pbblackjack.Suit_SPADES,
		pbblackjack.Suit_HEARTS,
		pbblackjack.Suit_DIAMONDS,
		pbblackjack.Suit_CLUBS,
	}
)

// Shuffles all Cards in the deck
func (d *Deck) Shuffle() {
	// maybe we can add shuffle types
	// riffle, random, cut etc

	// empty all of the Cards
	var cards []Card
	for len(d.Cards) > 0 {
		card := d.DrawCard()
		cards = append(cards, card)
	}
	shuffleOrder := rand.Perm(len(cards))

	for i := range shuffleOrder {
		card := cards[i]
		d.AddCard(card)
	}
}

// DrawCard draws a card from the top of the deck (face down) and removes the card from the deck
func (d *Deck) DrawCard() Card {
	// deck may be empty. in this case return an error
	top := d.Cards[len(d.Cards)-1]
	d.Cards = d.Cards[:len(d.Cards)-1]
	return top
}

// AddCard Adds a card to the bottom of the deck
func (d *Deck) AddCard(card Card) {
	d.Cards = append([]Card{card}, d.Cards...)
}

func (d *Deck) Len() int {
	return len(d.Cards)
}

func NewDeck() *Deck {
	var cards []Card

	for _, rank := range Ranks {
		for _, suit := range Suits {
			card := Card{
				Suit: suit,
				Rank: rank,
			}
			cards = append(cards, card)
		}
	}

	return &Deck{
		Cards: cards,
		rand:  rand.New(rand.NewSource(rand.Int63())),
	}
}

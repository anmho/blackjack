package blackjack

import (
	pb_blackjack "github.com/anmho/blackjack/gen/proto/blackjack"
)

type Card struct {
	Suit pb_blackjack.Suit
	Rank pb_blackjack.Rank
}

// Value returns the value of the card. Aces have two values

func (c Card) IsAce() bool {
	return c.Rank == pb_blackjack.Rank_ACE
}

func (c Card) ToProto() *pb_blackjack.Card {
	return &pb_blackjack.Card{
		Rank: c.Rank,
		Suit: c.Suit,
	}
}

type Hand struct {
	Cards []Card
}

func NewHand() *Hand {
	cards := make([]Card, 0)
	return &Hand{
		Cards: cards,
	}
}

func (h *Hand) AddCard(card Card) {
	h.Cards = append(h.Cards, card)
}

func (h *Hand) ToProto() *pb_blackjack.Hand {
	var cards []*pb_blackjack.Card
	for _, card := range h.Cards {
		cards = append(cards, card.ToProto())
	}
	return &pb_blackjack.Hand{
		Cards: cards,
	}
}

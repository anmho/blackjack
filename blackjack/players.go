package blackjack

import (
	pbblackjack "github.com/anmho/blackjack/gen/proto/blackjack"
)

const (
	// InitialBankroll is initial bankroll in dollars
	InitialBankroll = 500
)

type Player struct {
	ID   string
	Hand *Hand
}

func NewPlayer(id string) *Player {
	hand := NewHand()
	return &Player{
		ID:   id,
		Hand: hand,
	}
}

func (p *Player) ToProto() *pbblackjack.Player {
	return &pbblackjack.Player{
		Id:       p.ID,
		Hand:     p.Hand.ToProto(),
		Wager:    0,
		Bankroll: InitialBankroll,
	}
}

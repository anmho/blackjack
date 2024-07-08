package blackjack

import (
	pbblackjack "github.com/anmho/blackjack/gen/proto/blackjack"
	"github.com/google/uuid"
)

const (
	// InitialBankroll is initial bankroll in dollars
	InitialBankroll = 500
)

type Player struct {
	ID          uuid.UUID
	DisplayName string
	Hand        *Hand
	Wager       int
	Bankroll    int
}

func NewPlayer(displayName string) *Player {
	hand := NewHand()
	return &Player{
		ID:          uuid.New(),
		DisplayName: displayName,
		Hand:        hand,
	}
}

func (p *Player) ToPlayerProto() *pbblackjack.Player {
	return &pbblackjack.Player{
		Id:       p.ID.String(),
		Hand:     p.Hand.ToProto(),
		Wager:    0,
		Bankroll: InitialBankroll,
	}
}

func (p *Player) ToDealerProto() *pbblackjack.Dealer {
	return &pbblackjack.Dealer{
		Hand: p.Hand.ToProto(),
	}
}

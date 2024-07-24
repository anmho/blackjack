package blackjack

import (
	"fmt"
	pbblackjack "github.com/anmho/blackjack/gen/proto/blackjack"
	"github.com/google/uuid"
)

type PlayerStatus int

const (
	BrowsingStatus PlayerStatus = iota
	PlayingStatus
)

const (
	// InitialBankroll is initial bankroll in dollars
	InitialBankroll = 500
)

type Player struct {
	ID uuid.UUID
	// DisplayName is the name of the user
	Status      PlayerStatus
	DisplayName string
	Hand        *Hand
	Wager       int
	Bankroll    int
}

// NewPlayer creates a new player
// The user should have no display name initially.
// DisplayName is only set once they join an actual game.
// It must be unique per that game
func NewPlayer() *Player {
	hand := NewHand()
	// give them a temporary name?
	playerID := uuid.New()
	tempName := fmt.Sprintf("player-%s", playerID.String())
	return &Player{
		ID:          playerID,
		DisplayName: tempName,
		Hand:        hand,
		Status:      BrowsingStatus,
		Bankroll:    InitialBankroll,
	}
}

func NewDealer() *Player {
	player := NewPlayer()
	player.SetPlayingStatus(PlayingStatus)
	player.SetDisplayName("dealer")
	return player
}
func (p *Player) ToPlayerProto() *pbblackjack.Player {
	return &pbblackjack.Player{
		Id:       p.ID.String(),
		Hand:     p.Hand.ToProto(),
		Wager:    0,
		Bankroll: InitialBankroll,
	}
}

func (p *Player) SetDisplayName(displayName string) {
	p.DisplayName = displayName
}

func (p *Player) SetPlayingStatus(status PlayerStatus) {
	p.Status = status
}

func (p *Player) ToDealerProto() *pbblackjack.Dealer {
	return &pbblackjack.Dealer{
		Hand: p.Hand.ToProto(),
	}
}

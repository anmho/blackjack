package blackjack

import (
	pb_blackjack "github.com/anmho/blackjack/gen/proto/blackjack"
	"github.com/google/uuid"
)

type Player struct {
	hand []pb_blackjack.Card
}

type Game struct {
	GameID  uuid.UUID
	Players []*Player
}

func NewGame() Game {
	players := make([]*Player, 0)

	gameID := uuid.New()
	return Game{
		GameID:  gameID,
		Players: players,
	}
}

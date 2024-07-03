package blackjack

import (
	"errors"
	pbblackjack "github.com/anmho/blackjack/gen/proto/blackjack"
	"github.com/google/uuid"
)

type Game struct {
	GameID  uuid.UUID
	Players []*Player
}

const (
	MaxPlayers = 10
)

var (
	ErrMaxPlayersReached = errors.New("max players reached")
	ErrPlayerNotFound    = errors.New("player not found")
)

func (g *Game) ToProto() *pbblackjack.Game {
	var players []*pbblackjack.Player
	for _, player := range g.Players {
		playerProto := player.ToProto()
		players = append(players, playerProto)
	}

	return &pbblackjack.Game{
		Id:      g.GameID.String(),
		Players: players,
		Dealer:  nil,
	}
}

func (g *Game) IsFull() bool {
	return len(g.Players) == MaxPlayers
}
func (g *Game) Join() (*Player, error) {
	if g.IsFull() {
		return nil, ErrMaxPlayersReached
	}
	player := NewPlayer(uuid.NewString())
	g.Players = append(g.Players, player)
	return player, nil
}

func (g *Game) FindPlayerByID(id string) (*Player, error) {
	for i := range len(g.Players) {
		if g.Players[i].ID == id {
			return g.Players[i], nil
		}
	}
	return nil, ErrPlayerNotFound
}

func (g *Game) Leave(id string) error {
	//for i := range len(g.Players) {
	//	if g.Players[i].ID == player.ID {
	//		g.Players[i] = nil
	//		return nil
	//	}
	//}
	return ErrPlayerNotFound
}
func NewGame() *Game {
	var players []*Player

	gameID := uuid.New()
	return &Game{
		GameID:  gameID,
		Players: players,
	}
}

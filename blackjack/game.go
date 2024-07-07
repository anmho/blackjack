package blackjack

import (
	pbblackjack "github.com/anmho/blackjack/gen/proto/blackjack"
	"github.com/google/uuid"
)

type GameState int

const (
	StartedState GameState = iota
	WaitingState
)

type Game struct {
	Dealer  *Player
	GameID  uuid.UUID
	Players []*Player
	Deck    *Deck
}

const (
	MaxPlayers = 10
)

func NewGame() *Game {
	var players []*Player
	return &Game{
		GameID:  uuid.New(),
		Players: players,
		Deck:    NewDeck(),
		Dealer:  NewPlayer("dealer"),
	}
}

func (g *Game) ToProto() *pbblackjack.Game {
	var players []*pbblackjack.Player
	for _, player := range g.Players {
		playerProto := player.ToPlayerProto()
		players = append(players, playerProto)
	}
	dealer := NewPlayer("dealer")

	return &pbblackjack.Game{
		Id:      g.GameID.String(),
		Players: players,
		Dealer:  dealer.ToDealerProto(),
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

// Start starts the game if it isn't already started.
// The player starting the game must have previously joined the game.
// If game is already started return ErrInvalidState.
func (g *Game) Start() error {
	return nil
}

func (g *Game) Leave(playerID string) error {
	//for i := range len(g.Players) {
	//	if g.Players[i].ID == player.ID {
	//		g.Players[i] = nil
	//		return nil
	//	}
	//}
	return ErrPlayerNotFound
}

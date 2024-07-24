package blackjack

import (
	"errors"
	pbblackjack "github.com/anmho/blackjack/gen/proto/blackjack"
	"github.com/google/uuid"
)

type Game interface {
	Start(playerID uuid.UUID) error
	Leave(playerID uuid.UUID) error
	IsFull() bool
	ToProto() *pbblackjack.Game
	IsJoined(playerID uuid.UUID) bool
	ID() uuid.UUID
	Players() []*Player
	ActivePlayers() int
	MaxPlayers() int

	// Player actions
	Join(player *Player) error
	MakeWager(playerID uuid.UUID, wager int) error
}

type game struct {
	Dealer        *Player
	GameID        uuid.UUID
	seats         [MaxPlayers]*Player
	Deck          *Deck
	Status        pbblackjack.GameStatus
	activePlayers int
}

const (
	MaxPlayers = 10
)

func NewGame() Game {
	var seats [MaxPlayers]*Player
	dealer := NewDealer()
	return &game{
		GameID: uuid.New(),
		seats:  seats,
		Status: pbblackjack.GameStatus_NOT_STARTED,
		Deck:   NewDeck(),
		Dealer: dealer,
	}
}

func (g *game) MakeWager(playerID uuid.UUID, wager int) error {
	//TODO implement me
	panic("implement me")
}

// Wagers returns the wagers of all Players in the game
func (g *game) Wagers() map[uuid.UUID]int {
	wagers := map[uuid.UUID]int{}

	for _, player := range g.Players() {
		wagers[player.ID] = player.Wager
	}
	return wagers
}

func (g *game) ActivePlayers() int {
	return g.activePlayers
}

func (g *game) MaxPlayers() int {
	return MaxPlayers
}

func (g *game) ID() uuid.UUID {
	return g.GameID
}

// Players returns a slice of the Players currently joined in the game with no gaps
func (g *game) Players() []*Player {
	var players []*Player
	for _, player := range g.seats {
		if player != nil {
			players = append(players, player)
		}
	}
	return players
}

func (g *game) ToProto() *pbblackjack.Game {
	var players []*pbblackjack.Player
	for _, player := range g.Players() {
		playerProto := player.ToPlayerProto()
		players = append(players, playerProto)
	}
	dealer := NewPlayer()
	dealer.SetDisplayName("dealer")

	return &pbblackjack.Game{
		Id:      g.GameID.String(),
		Players: players,
		Dealer:  dealer.ToDealerProto(),
		Status:  g.Status,
	}
}

func (g *game) IsFull() bool {
	return g.ActivePlayers() == MaxPlayers
}

func (g *game) Join(player *Player) error {
	if player == nil {
		return errors.New("player is nil")
	}
	if g.IsFull() {
		return ErrMaxPlayersReached
	}

	// find the first seat
	for seat, _ := range g.Players() {
		if g.seats[seat] == nil {
			g.seats[seat] = player
			break
		}
	}
	return nil
}

func (g *game) FindPlayerByID(id string) (*Player, error) {
	//for i := range len(g.Players) {
	for _, player := range g.Players() {
		if player.ID.String() == id {
			return player, nil
		}
	}
	return nil, ErrPlayerNotFound
}

// Start starts the games if it isn't already started.
// The player starting the games must have previously joined the games.
// If games is already started return ErrInvalidState.
func (g *game) Start(playerID uuid.UUID) error {
	g.Status = pbblackjack.GameStatus_IN_PROGRESS
	return nil
}

func (g *game) Leave(playerID uuid.UUID) error {
	// set their spot to null
	for i := range len(g.seats) {
		if g.seats[i].ID == playerID {
			g.seats[i] = nil
			return nil
		}
	}
	return ErrPlayerNotFound
}

func (g *game) IsJoined(playerID uuid.UUID) bool {
	for _, player := range g.Players() {
		if player.ID.String() == playerID.String() {
			return true
		}
	}

	return false
}

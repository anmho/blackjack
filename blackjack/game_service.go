package blackjack

import (
	"fmt"
	"github.com/google/uuid"
)

type GameService interface {
	GetGames() []Game
	// GetGameByID Gets the game by id.
	GetGameByID(gameID uuid.UUID) (Game, error)
	// AddGame opens up a new game.
	AddGame() error
	// GetJoinedGame gets the game a player has joined.
	GetJoinedGame(playerID uuid.UUID) (Game, error)
	// LeaveGame leaves the player's currently joined game.
	LeaveGame(playerID uuid.UUID) error
}
type gameService struct {
	games map[uuid.UUID]Game
	// PlayerID to GameID mappings
	joinedGames map[uuid.UUID]uuid.UUID
}

func (g *gameService) GetGames() []Game {
	games := make([]Game, 0)

	for _, game := range g.games {
		if game != nil {
			fmt.Printf("%+v\n", game)
			games = append(games, game)
		}
	}
	return games
}

func (g *gameService) GetGameByID(gameID uuid.UUID) (Game, error) {
	game, ok := g.games[gameID]
	if !ok || game == nil {
		return nil, ErrGameNotFound
	}
	return g.games[gameID], nil
}

func (g *gameService) AddGame() error {
	newGame := NewGame()
	g.games[newGame.ID()] = newGame
	return nil
}

func (g *gameService) GetJoinedGame(playerID uuid.UUID) (Game, error) {
	gameID := g.joinedGames[playerID]
	joinedGame, err := g.GetGameByID(gameID)
	return joinedGame, err
}

func (g *gameService) LeaveGame(playerID uuid.UUID) error {
	joinedGame, ok := g.games[playerID]
	if !ok || joinedGame == nil {
		return ErrNotJoined
	}

	delete(g.games, playerID)

	return nil
}

func NewGameService() GameService {
	return &gameService{
		games:       make(map[uuid.UUID]Game),
		joinedGames: make(map[uuid.UUID]uuid.UUID),
	}
}

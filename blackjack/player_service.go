package blackjack

import "github.com/google/uuid"

type PlayerService interface {
	CreatePlayer() *Player
}

type playerService struct {
	Players map[uuid.UUID]*Player
}

func NewPlayerService() PlayerService {
	players := make(map[uuid.UUID]*Player)
	return &playerService{
		Players: players,
	}
}

func (ps *playerService) CreatePlayer() *Player {
	p := NewPlayer()
	ps.Players[p.ID] = p
	return p
}

func (ps *playerService) Disconnect(playerID uuid.UUID) {

}

func (ps *playerService) GetPlayerByID(playerID uuid.UUID) *Player {
	player, ok := ps.Players[playerID]
	if !ok {
		return nil
	}
	return player
}

func (ps *playerService) PlayerExists(playerID uuid.UUID) bool {
	_, exists := ps.Players[playerID]
	return exists
}

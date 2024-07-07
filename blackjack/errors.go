package blackjack

import "errors"

var (
	ErrNotJoined         = errors.New("player has not joined a game")
	ErrNotFound          = errors.New("resource not found")
	ErrInvalidState      = errors.New("game is in an invalid state")
	ErrMaxPlayersReached = errors.New("max players reached")
	ErrPlayerNotFound    = errors.New("player not found")
)

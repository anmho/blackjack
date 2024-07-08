package blackjack

import "errors"

var (
	ErrNotJoined         = errors.New("player has not joined a games")
	ErrNotFound          = errors.New("resource not found")
	ErrInvalidState      = errors.New("games is in an invalid state")
	ErrMaxPlayersReached = errors.New("max players reached")
	ErrPlayerNotFound    = errors.New("player not found")
)

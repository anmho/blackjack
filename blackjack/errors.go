package blackjack

import "errors"

var (
	ErrNotJoined = errors.New("player has not joined a game")
	ErrNotFound  = errors.New("resource not found")
)

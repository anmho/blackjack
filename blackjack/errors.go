package blackjack

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNotJoined         = errors.New("player has not joined a games")
	ErrNotFound          = errors.New("resource not found")
	ErrInvalidState      = errors.New("games is in an invalid state")
	ErrMaxPlayersReached = errors.New("max Players reached")
	ErrPlayerNotFound    = errors.New("player not found")

	ErrNilRequest   = status.Errorf(codes.InvalidArgument, "nil request")
	ErrGameNotFound = errors.New("games not found")

	ErrGameExists = errors.New("game already exists")

	ErrUnsupportedResultType = errors.New("unsupported blackjack result type")
)

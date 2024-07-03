package blackjack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGame_IsFull(t *testing.T) {

	game := NewGame()

	assert.False(t, game.IsFull(), "is false")
}

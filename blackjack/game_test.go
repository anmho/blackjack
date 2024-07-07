package blackjack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGame_IsFull(t *testing.T) {
	game := NewGame()

	t.Run("new game is not full (no players)", func(t *testing.T) {
		assert.False(t, game.IsFull())
	})
}

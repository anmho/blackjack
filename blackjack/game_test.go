package blackjack

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGame_Join(t *testing.T) {
	// join the games

	var err error

	g := NewGame()

	var players []uuid.UUID
	for range g.MaxPlayers() {
		players = append(players, uuid.New())
	}

	for i := range players[:10] {
		err = g.Join(players[i])
		require.NoError(t, err)
		assert.True(t, g.IsJoined(players[i]))
	}

	assert.False(t, g.IsFull())
}

func TestGame_Leave(t *testing.T) {

}

func TestGame_Start(t *testing.T) {

}

func TestGame_IsFull(t *testing.T) {
	game := NewGame()

	t.Run("new games is not full (no players)", func(t *testing.T) {
		assert.False(t, game.IsFull())
	})
}

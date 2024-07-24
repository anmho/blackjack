package blackjack

import (
	pbblackjack "github.com/anmho/blackjack/gen/proto/blackjack"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGame_IsJoined(t *testing.T) {
	g := NewGame()
	p := NewPlayer()

	err := g.Join(p)
	require.NoError(t, err)

	g.IsJoined(p.ID)
}

//
//func TestGame_Join(t *testing.T) {
//	// join the games
//
//	var err error
//
//	g := NewGame()
//
//	var Players []*Player
//	for i := range g.MaxPlayers() {
//		Players = append(Players, NewPlayer(fmt.Sprintf("joey%d", i)))
//	}
//
//	for i := range Players[:10] {
//		err = g.Join(Players[i])
//		require.NoError(t, err)
//		assert.True(t, g.IsJoined(Players[i].ID))
//	}
//
//	assert.False(t, g.IsFull())
//}

func TestGame_ToProto(t *testing.T) {
	g := NewGame()

	pbGame := g.ToProto()

	assert.Equal(t, pbblackjack.GameStatus_NOT_STARTED, pbGame.Status)
}

func TestGame_Leave(t *testing.T) {

}

func TestGame_Start(t *testing.T) {

}

func TestGame_IsFull(t *testing.T) {
	game := NewGame()

	t.Run("new games is not full (no Players)", func(t *testing.T) {
		assert.False(t, game.IsFull())
	})
}

package blackjack

import (
	pbblackjack "github.com/anmho/blackjack/gen/proto/blackjack"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGame_IsJoined(t *testing.T) {
	g := NewGame()
	p := NewPlayer("joseph")

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
//	var players []*Player
//	for i := range g.MaxPlayers() {
//		players = append(players, NewPlayer(fmt.Sprintf("joey%d", i)))
//	}
//
//	for i := range players[:10] {
//		err = g.Join(players[i])
//		require.NoError(t, err)
//		assert.True(t, g.IsJoined(players[i].ID))
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

	t.Run("new games is not full (no players)", func(t *testing.T) {
		assert.False(t, game.IsFull())
	})
}

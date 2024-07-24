package blackjack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type TestFixture struct {
	t               *testing.T
	BlackjackServer Server
}

func NewTestFixture(t *testing.T) *TestFixture {
	gameSvc := NewGameService()
	playerSvc := NewPlayerService()
	blackjackSvc := MakeBlackjackService(gameSvc, playerSvc)
	require.NotNil(t, gameSvc)
	require.NotNil(t, blackjackSvc)

	tf := &TestFixture{
		t:               t,
		BlackjackServer: blackjackSvc,
	}
	return tf
}

func (tf *TestFixture) getServer() *server {
	s, ok := tf.BlackjackServer.(*server)
	require.True(tf.t, ok)
	require.NotNil(tf.t, s)
	return s
}

func (tf *TestFixture) CreateGame() Game {
	s := tf.getServer()
	g, err := s.gameService.CreateGame()
	require.NoError(tf.t, err)
	return g
}

func (tf *TestFixture) CreatePlayer() *Player {
	s := tf.getServer()
	p := s.playerService.CreatePlayer()
	require.NotNil(tf.t, p)
	return p
}

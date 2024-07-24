package blackjack

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func makeTestGameService(t *testing.T) *gameService {
	s, ok := NewGameService().(*gameService)
	require.True(t, ok)
	return s
}

func TestNewGameService(t *testing.T) {
	s := NewGameService()
	require.NotNil(t, s)

	games := s.GetGames()

	assert.NotNil(t, games)
	assert.Len(t, games, 0)
}

func TestGameService_CreateGame(t *testing.T) {
	t.Run("add should add games", func(t *testing.T) {
		s := NewGameService()
		assert.Equal(t, 0, len(s.GetGames()))
		var err error
		var games []Game
		g, err := s.CreateGame()
		require.NoError(t, err)
		assert.NotNil(t, g)
		games = s.GetGames()
		assert.Len(t, games, 1)
	})
}

func TestGameService_GetGames(t *testing.T) {
	t.Run("should get games", func(t *testing.T) {
		g1, g2, g3 := uuid.New(), uuid.New(), uuid.New()
		var games map[uuid.UUID]Game
		games = map[uuid.UUID]Game{
			g1: NewGame(),
			g2: NewGame(),
			g3: NewGame(),
		}

		s := makeTestGameService(t)
		s.games = games

		assert.Equal(t, 3, len(s.GetGames()))
	})
}

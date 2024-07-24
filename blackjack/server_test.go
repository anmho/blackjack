package blackjack

import (
	pbblackjack "github.com/anmho/blackjack/gen/proto/blackjack"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestServer_ViewGames(t *testing.T) {
	tests := []struct {
		testName string
		numGames int
		request  *pbblackjack.ViewGamesRequest

		expectedGames int
	}{
		{
			testName: "zero games",
			numGames: 0,
			request:  &pbblackjack.ViewGamesRequest{},

			expectedGames: 0,
		},
		{
			testName: "1 game",
			numGames: 1,
			request:  &pbblackjack.ViewGamesRequest{},

			expectedGames: 1,
		},
		{
			testName: "2 games",
			numGames: 2,
			request:  &pbblackjack.ViewGamesRequest{},

			expectedGames: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			tf := NewTestFixture(t)
			for range test.numGames {
				g := tf.CreateGame()
				require.NotNil(t, g)
			}

			s, ok := tf.BlackjackServer.(*server)
			require.True(t, ok)

			result, err := s.ViewGames(test.request)
			require.NoError(t, err)
			assert.NotNil(t, result)
			assert.NotNil(t, result.Games)
			assert.Len(t, result.Games, test.expectedGames)
		})
	}

}

func TestServer_JoinGame(t *testing.T) {
	tf := NewTestFixture(t)
	g := tf.CreateGame()
	require.NotNil(t, g)
	p := tf.CreatePlayer()
	// Attempt to join the game we just made
	result, err := tf.BlackjackServer.JoinGame(&pbblackjack.JoinGameRequest{
		GameId:   g.ID().String(),
		Username: "Joe mama",
	}, &p)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.True(t, result.Success)

}

func TestServer_StartGame(t *testing.T) {}

func TestServer_LeaveGame(t *testing.T) {}

func TestServer_ViewCurrentGame(t *testing.T) {}

func Test_MakeBlackjackResponse(t *testing.T) {}

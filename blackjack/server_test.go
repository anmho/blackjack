package blackjack

import (
	"testing"
)

func makeTestServer() *server {
	s, _ := NewService().(*server)
	return s
}

func TestServer_ViewGames(t *testing.T) {

	//s := makeTestServer()
	//
	//require.NotNil(t, s)
	//s.StartGame()

}

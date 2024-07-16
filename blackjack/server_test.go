package blackjack

import (
	"testing"
)

func makeTestServer() *server {
	s, _ := MakeBlackjackService().(*server)
	return s
}

func TestServer_ViewGames(t *testing.T) {

	//s := makeTestServer()
	//
	//require.NotNil(t, s)
	//s.StartGame()

}

package blackjack

import (
	"errors"
	pbblackjack "github.com/anmho/blackjack/gen/proto/blackjack"
	"io"
	"log/slog"
)

var (
	ErrGameNotFound = errors.New("game not found")
)

var _ pbblackjack.BlackJackServiceServer = (*service)(nil)

type service struct {
	pbblackjack.UnimplementedBlackJackServiceServer
	games map[string]*Game
}

func NewService() pbblackjack.BlackJackServiceServer {
	const NumGames = 8
	// initialize 8 games
	games := make(map[string]*Game)
	for range NumGames {
		game := NewGame()
		games[game.GameID.String()] = game
	}

	s := &service{
		games: games,
	}
	return s
}

func (s *service) Connect(connectServer pbblackjack.BlackJackService_ConnectServer) error {
	//var player *Player

	for {
		var err error
		var request *pbblackjack.BlackjackRequest
		var response pbblackjack.BlackjackResponse
		request, err = connectServer.Recv()
		if err != nil {
			if err == io.EOF {
				// we should remove them from the game
				break
			}
			slog.Error(err.Error())
			break
		}

		var viewGamesResult *pbblackjack.ViewGamesResult
		var joinGameResult *pbblackjack.JoinGameResult
		var startGameResult *pbblackjack.StartGameResult
		var leaveGameResult *pbblackjack.LeaveGameResult

		switch request := request.Request.(type) {
		case *pbblackjack.BlackjackRequest_ViewGamesRequest:
			// A connected player can always view the status of currently played games
			viewGamesResult, err = s.viewGames(request.ViewGamesRequest)
		case *pbblackjack.BlackjackRequest_JoinGameRequest:
			// A player can join a game if they are not already in a game
			joinGameResult, err = s.joinGame(request.JoinGameRequest)
		case *pbblackjack.BlackjackRequest_StartGameRequest:
			// A player must join a game before starting that game. The game must not be started.
			startGameResult, err = s.startGame(request.StartGameRequest)
		case *pbblackjack.BlackjackRequest_LeaveGameRequest:
			// A player must have joined a game before leaving a game.
			leaveGameResult, err = s.leaveGame(request.LeaveGameRequest)
		}

		if err != nil {
			return err
		}

		if viewGamesResult != nil {
			response.Result = &pbblackjack.BlackjackResponse_ViewGamesResult{
				ViewGamesResult: viewGamesResult,
			}
		} else if joinGameResult != nil {
			response.Result = &pbblackjack.BlackjackResponse_JoinGameResult{
				JoinGameResult: joinGameResult,
			}
		} else if startGameResult != nil {
			response.Result = &pbblackjack.BlackjackResponse_StartGameResult{
				StartGameResult: startGameResult,
			}
		} else if leaveGameResult != nil {
			response.Result = &pbblackjack.BlackjackResponse_LeaveGameResult{
				LeaveGameResult: leaveGameResult,
			}
		}

		err = connectServer.Send(&response)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) viewGames(request *pbblackjack.ViewGamesRequest) (*pbblackjack.ViewGamesResult, error) {
	return nil, nil
}

func (s *service) joinGame(request *pbblackjack.JoinGameRequest) (*pbblackjack.JoinGameResult, error) {
	return nil, nil
}

func (s *service) startGame(request *pbblackjack.StartGameRequest) (*pbblackjack.StartGameResult, error) {
	return nil, nil
}

// leaveGame leaves the player's currently joined game. If the player is not currently joined a game we return ErrNotJoined.
func (s *service) leaveGame(request *pbblackjack.LeaveGameRequest) (*pbblackjack.LeaveGameResult, error) {
	return nil, nil
}

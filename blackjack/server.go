package blackjack

import (
	"errors"
	pbblackjack "github.com/anmho/blackjack/gen/proto/blackjack"
	"io"
	"log/slog"
)

const NumGames = 8

var (
	ErrGameNotFound = errors.New("games not found")
)

var _ pbblackjack.BlackJackServiceServer = (*service)(nil)

type service struct {
	pbblackjack.UnimplementedBlackJackServiceServer
	games map[string]Game
	// joinedGames tracks the games each player is currently joined as. userID -> gameID
	joinedGames map[string]string
}

func NewService() pbblackjack.BlackJackServiceServer {

	// initialize 8 games
	games := make(map[string]Game)
	for range NumGames {
		game := NewGame()
		games[game.ID().String()] = game
	}

	s := &service{
		games: games,
	}
	return s
}

func (s *service) Connect(stream pbblackjack.BlackJackService_ConnectServer) error {
	var player *Player
	for {
		var err error
		var request *pbblackjack.BlackjackRequest
		var response pbblackjack.BlackjackResponse
		request, err = stream.Recv()
		if err != nil {
			if err == io.EOF {
				// we should remove them from the games
				break
			}
			slog.Error(err.Error())
			break
		}

		var viewGamesResult *pbblackjack.ViewGamesResult
		var joinGameResult *pbblackjack.JoinGameResult
		var startGameResult *pbblackjack.StartGameResult
		var leaveGameResult *pbblackjack.LeaveGameResult
		var viewCurrentGameResult *pbblackjack.ViewCurrentGameResult

		switch request := request.Request.(type) {
		case *pbblackjack.BlackjackRequest_ViewGamesRequest:
			// A connected player can always view the status of currently played games
			viewGamesResult, err = s.viewGames(request.ViewGamesRequest)
		case *pbblackjack.BlackjackRequest_JoinGameRequest:
			// A player can join a games if they are not already in a games
			joinGameResult, err = s.joinGame(request.JoinGameRequest, &player)
		case *pbblackjack.BlackjackRequest_StartGameRequest:
			// A player must join a games before starting that games. The games must not be started.
			startGameResult, err = s.startGame(request.StartGameRequest, &player)
		case *pbblackjack.BlackjackRequest_LeaveGameRequest:
			// A player must have joined a games before leaving a games.
			leaveGameResult, err = s.leaveGame(request.LeaveGameRequest, &player)
		case *pbblackjack.BlackjackRequest_ViewCurrentGameRequest:
			viewCurrentGameResult, err = s.viewCurrentGame(request.ViewCurrentGameRequest, &player)
		}

		if err != nil {
			return err
		}

		setResponseResult(
			&response,
			viewGamesResult,
			joinGameResult,
			startGameResult,
			leaveGameResult,
			viewCurrentGameResult,
		)

		err = stream.Send(&response)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) viewGames(request *pbblackjack.ViewGamesRequest) (*pbblackjack.ViewGamesResult, error) {

	var games []*pbblackjack.Game
	for _, game := range s.games {

		games = append(games, game.ToProto())
	}

	result := &pbblackjack.ViewGamesResult{
		Games: games,
	}

	return result, nil
}

func (s *service) joinGame(request *pbblackjack.JoinGameRequest, player **Player) (*pbblackjack.JoinGameResult, error) {
	p := NewPlayer(request.Username)
	*player = p
	return &pbblackjack.JoinGameResult{
		Success: true,
	}, nil
}

func (s *service) startGame(request *pbblackjack.StartGameRequest, player **Player) (*pbblackjack.StartGameResult, error) {

	return nil, nil
}

// leaveGame leaves the player's currently joined games. If the player is not currently joined a games we return ErrNotJoined.
func (s *service) leaveGame(request *pbblackjack.LeaveGameRequest, player **Player) (*pbblackjack.LeaveGameResult, error) {
	return nil, nil
}

func (s *service) viewCurrentGame(request *pbblackjack.ViewCurrentGameRequest, player **Player) (*pbblackjack.ViewCurrentGameResult, error) {
	if player == nil || *player == nil {
		return nil, ErrNotJoined
	}

	pbPlayer := (*player).ToPlayerProto()

	result := pbblackjack.ViewCurrentGameResult{
		Player: pbPlayer,
		Game: &pbblackjack.Game{
			Id:      "",
			Players: nil,
			Dealer: &pbblackjack.Dealer{
				Hand: &pbblackjack.Hand{
					Cards: nil,
				},
			},
		},
	}
	return &result, nil
}

func setResponseResult(response *pbblackjack.BlackjackResponse, results ...interface{}) {
	for _, result := range results {
		switch v := result.(type) {
		case *pbblackjack.ViewGamesResult:
			if v != nil {
				response.Result = &pbblackjack.BlackjackResponse_ViewGamesResult{ViewGamesResult: v}
			}
		case *pbblackjack.JoinGameResult:
			if v != nil {
				response.Result = &pbblackjack.BlackjackResponse_JoinGameResult{JoinGameResult: v}
			}
		case *pbblackjack.StartGameResult:
			if v != nil {
				response.Result = &pbblackjack.BlackjackResponse_StartGameResult{StartGameResult: v}
			}
		case *pbblackjack.LeaveGameResult:
			if v != nil {
				response.Result = &pbblackjack.BlackjackResponse_LeaveGameResult{LeaveGameResult: v}
			}
		case *pbblackjack.ViewCurrentGameResult:
			if v != nil {
				response.Result = &pbblackjack.BlackjackResponse_ViewCurrentGameResult{ViewCurrentGameResult: v}
			}
		}
	}
}

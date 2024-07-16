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

var _ pbblackjack.BlackJackServiceServer = (*server)(nil)

type server struct {
	pbblackjack.UnimplementedBlackJackServiceServer
	gameService GameService
}

func MakeBlackjackService() pbblackjack.BlackJackServiceServer {
	gameSvc := NewGameService()
	s := &server{
		gameService: gameSvc,
	}
	return s
}

// Connect connects a players to the game and enters the control loop.
func (s *server) Connect(stream pbblackjack.BlackJackService_ConnectServer) error {
	var player *Player
	for {
		var err error
		var response *pbblackjack.BlackjackResponse
		request, err := stream.Recv()
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
			viewGamesResult, err = s.ViewGames(request.ViewGamesRequest)
			response = MakeBlackjackResponse(viewGamesResult)
		case *pbblackjack.BlackjackRequest_JoinGameRequest:
			// A player can join a games if they are not already in a games
			joinGameResult, err = s.JoinGame(request.JoinGameRequest, &player)
			response = MakeBlackjackResponse(joinGameResult)
		case *pbblackjack.BlackjackRequest_StartGameRequest:
			// A player must join a games before starting that games. The games must not be started.
			startGameResult, err = s.StartGame(request.StartGameRequest, &player)
			response = MakeBlackjackResponse(startGameResult)
		case *pbblackjack.BlackjackRequest_LeaveGameRequest:
			// A player must have joined a games before leaving a games.
			leaveGameResult, err = s.LeaveGame(request.LeaveGameRequest, &player)
			response = MakeBlackjackResponse(leaveGameResult)
		case *pbblackjack.BlackjackRequest_ViewCurrentGameRequest:
			viewCurrentGameResult, err = s.ViewCurrentGame(request.ViewCurrentGameRequest, &player)
			response = MakeBlackjackResponse(viewCurrentGameResult)
		}

		if err != nil {
			return err
		}

		err = stream.Send(response)
		if err != nil {
			return err
		}
	}
	return nil
}

// ViewGames displays the games currently available on the blackjack server.
func (s *server) ViewGames(request *pbblackjack.ViewGamesRequest) (*pbblackjack.ViewGamesResult, error) {
	result := &pbblackjack.ViewGamesResult{}

	return result, nil
}

// JoinGame joins the currently connected user to the game.
func (s *server) JoinGame(request *pbblackjack.JoinGameRequest, player **Player) (*pbblackjack.JoinGameResult, error) {
	p := NewPlayer(request.Username)
	*player = p
	return &pbblackjack.JoinGameResult{
		Success: true,
	}, nil
}

// StartGame starts the game the user is currently connected to.
func (s *server) StartGame(request *pbblackjack.StartGameRequest, player **Player) (*pbblackjack.StartGameResult, error) {

	return nil, nil
}

// LeaveGame leaves the player's currently joined games. If the player is not currently joined a games we return ErrNotJoined.
func (s *server) LeaveGame(request *pbblackjack.LeaveGameRequest, player **Player) (*pbblackjack.LeaveGameResult, error) {
	return nil, nil
}

func (s *server) ViewCurrentGame(request *pbblackjack.ViewCurrentGameRequest, player **Player) (*pbblackjack.ViewCurrentGameResult, error) {
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

func MakeBlackjackResponse(result interface{}) *pbblackjack.BlackjackResponse {
	var response pbblackjack.BlackjackResponse
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
	return &response
}

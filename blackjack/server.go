package blackjack

import (
	"context"
	"errors"
	"fmt"
	pbblackjack "github.com/anmho/blackjack/gen/proto/blackjack"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log/slog"
)

var (
	ErrGameNotFound = errors.New("game not found")
)

type service struct {
	pbblackjack.UnimplementedBlackJackServiceServer
	games map[string]*Game
}

func (s *service) ViewGames(ctx context.Context, request *pbblackjack.ViewGamesRequest) (*pbblackjack.ViewGamesResponse, error) {
	var games []*pbblackjack.Game
	for _, game := range s.games {
		gameProto := game.ToProto()
		games = append(games, gameProto)
	}

	response := &pbblackjack.ViewGamesResponse{
		Games: games,
	}
	return response, nil
}

func (s *service) getGame(id string) (*Game, error) {
	game, ok := s.games[id]
	if !ok {
		return nil, ErrGameNotFound
	}
	return game, nil
}

func (s *service) handleJoinGame(req *pbblackjack.BlackjackRequest_JoinGameRequest) (*Player, error) {
	gameID := req.JoinGameRequest.GameId
	game, err := s.getGame(gameID)
	if err != nil {
		return nil, err
	}
	player, err := game.Join()
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (s *service) PlayGame(gameServer pbblackjack.BlackJackService_PlayGameServer) error {
	var playerID string
	for {
		fmt.Printf("player id %s\n", playerID)
		req, err := gameServer.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			slog.Error(err.Error())
			break
		}

		var response pbblackjack.BlackjackResponse
		var joinGameResult *pbblackjack.BlackjackResponse_JoinGameResult
		var makeMoveResult *pbblackjack.BlackjackResponse_BlackjackResult

		switch req := req.Action.(type) {
		case *pbblackjack.BlackjackRequest_JoinGameRequest:
			if playerID != "" {
				return status.Errorf(codes.PermissionDenied, "cannot join game you are already joined")
			}
			player, err := s.handleJoinGame(req)
			if err != nil {
				return status.Errorf(codes.Internal, err.Error())
			}

			playerID = player.ID
		case *pbblackjack.BlackjackRequest_StartGameRequest:
		case *pbblackjack.BlackjackRequest_BlackjackMove:
		}

		if joinGameResult != nil {
			response.Result = joinGameResult
		} else if makeMoveResult != nil {
			response.Result = makeMoveResult
		} else if makeMoveResult != nil {
			response.Result = makeMoveResult
		}
		err = gameServer.Send(&response)
		if err != nil {
			return err
		}
	}
	return nil
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

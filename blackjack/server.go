package blackjack

import (
	context "context"
	"fmt"
	pb_blackjack "github.com/anmho/blackjack/gen/proto/blackjack"
	"io"
	"log/slog"
)

type service struct {
	pb_blackjack.UnimplementedBlackJackServiceServer
}

func (s *service) ViewGames(ctx context.Context, request *pb_blackjack.ViewGamesRequest) (*pb_blackjack.ViewGamesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) PlayGame(gameServer pb_blackjack.BlackJackService_PlayGameServer) error {
	for {
		fmt.Println("hello")
		req, err := gameServer.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			slog.Error("error ", err.Error())
			break
		}
		fmt.Printf("%+v\n", req.GetJoinGame())
	}
	return nil
}

func NewService() pb_blackjack.BlackJackServiceServer {
	s := &service{}
	return s
}

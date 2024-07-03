package main

import (
	"fmt"
	"github.com/anmho/blackjack/blackjack"
	pb_blackjack "github.com/anmho/blackjack/gen/proto/blackjack"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = 50051
)

func main() {
	blackjackService := blackjack.NewService()

	s := grpc.NewServer()
	pb_blackjack.RegisterBlackJackServiceServer(s, blackjackService)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln(err)
	}
	reflection.Register(s)
	fmt.Printf("listening on %d\n", port)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}
}

package blackjack

import pb_blackjack "github.com/anmho/blackjack/gen/proto/blackjack"

type Hand struct {
	Cards []*pb_blackjack.Card
}

func NewHand() *Hand {
	cards := make([]*pb_blackjack.Card, 0)
	return &Hand{
		Cards: cards,
	}
}

func (h *Hand) ToProto() *pb_blackjack.Hand {
	return &pb_blackjack.Hand{
		Cards: h.Cards,
	}
}

package blackjack

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlayer_ToDealerProto(t *testing.T) {
	t.Parallel()
	t.Run("valid dealer proto", func(t *testing.T) {
		player := Player{
			ID:          uuid.New(),
			DisplayName: "joey moms",
			Hand:        NewHand(),
			Wager:       500,
			Bankroll:    2,
		}

		pbDealer := player.ToDealerProto()
		assert.NotNil(t, pbDealer.GetHand())

	})
}

func TestPlayer_ToPlayerProto(t *testing.T) {

}

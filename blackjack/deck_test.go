package blackjack

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/rand"
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	assert.Equal(t, 52, deck.Len())
}
func TestDeck_Shuffle(t *testing.T) {
	deck := NewDeck()
	originalCards := deck.Cards

	// To prevent flakiness lets set a non-random seed
	r := rand.New(rand.NewSource(0))
	deck.rand = r
	require.Equal(t, 52, deck.Len())
	deck.Shuffle()
	assert.Equal(t, 52, deck.Len())

	for i := range deck.Len() {
		reflect.DeepEqual(originalCards[i], deck.Cards[i])
	}
}

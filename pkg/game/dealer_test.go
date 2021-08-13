package game_test

import (
	"testing"

	"example.com/pkg/game"
	"example.com/pkg/game/gamefakes"
	"github.com/stretchr/testify/assert"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . GamePlayer
//counterfeiter:generate . GameDeck

func TestDealerDeal(t *testing.T) {
	fakePlayer1 := new(gamefakes.FakeGamePlayer)
	fakeDeck := new(gamefakes.FakeGameDeck)
	players := []game.GamePlayer{fakePlayer1}

	testObject := game.Dealer{players, fakeDeck}

	c1 := game.Card{2, game.HEARTS}
	c2 := game.Card{3, game.HEARTS}
	c3 := game.Card{4, game.HEARTS}
	c4 := game.Card{5, game.HEARTS}
	c5 := game.Card{6, game.HEARTS}
	fakeDeck.PopCardReturnsOnCall(0, c1, nil)
	fakeDeck.PopCardReturnsOnCall(1, c2, nil)
	fakeDeck.PopCardReturnsOnCall(2, c3, nil)
	fakeDeck.PopCardReturnsOnCall(3, c4, nil)
	fakeDeck.PopCardReturnsOnCall(4, c5, nil)

	testObject.Deal()

	assert.Equal(t, 5, fakePlayer1.TakeCallCount())
}

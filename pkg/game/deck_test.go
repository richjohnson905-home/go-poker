package game

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeckSize52(t *testing.T) {
	deck := MakeDeck()
	deck.Shuffle()
	//fmt.Println(shuffledCards)
	assert.Equal(t, 52, len(deck.Cards))
}

func TestDeckSize51AfterOneDeal(t *testing.T) {
	deck := MakeDeck()
	deck.Shuffle()
	c, _ := deck.PopCard()
	fmt.Println(c)
	assert.Equal(t, 51, len(deck.Cards))
}

func TestErrorWhenDealEmptyDeck(t *testing.T) {
	d := Deck{}
	_, e := d.PopCard()

	// fmt.Println(c)
	assert.True(t, e != nil, "Should have been an error")
}

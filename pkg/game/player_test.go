package game

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandHasFiveCardsAfterDeal(t *testing.T) {
	c1 := Card{Suit: HEARTS, Value: 2}
	c2 := Card{Suit: HEARTS, Value: 3}
	c3 := Card{Suit: HEARTS, Value: 4}
	c4 := Card{Suit: HEARTS, Value: 5}
	c5 := Card{Suit: HEARTS, Value: 6}

	testObject := Player{Name: "Bob"}
	testObject.Take(c1)
	testObject.Take(c2)
	testObject.Take(c3)
	testObject.Take(c4)
	testObject.Take(c5)

	assert.Equal(t, 5, len(testObject.Hand))
}

func TestPlayerHandHasCorrectHandSizeAfterEachDeal(t *testing.T) {
	var tests = []struct {
		c    Card
		want int
	}{
		{Card{Suit: HEARTS, Value: 2}, 1},
		{Card{Suit: HEARTS, Value: 3}, 2},
		{Card{Suit: HEARTS, Value: 4}, 3},
		{Card{Suit: HEARTS, Value: 5}, 4},
		{Card{Suit: HEARTS, Value: 6}, 5},
	}
	testObject := Player{Name: "Bob"}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.c)
		t.Run(testname, func(t *testing.T) {
			testObject.Take(tt.c)
			ans := len(testObject.Hand)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

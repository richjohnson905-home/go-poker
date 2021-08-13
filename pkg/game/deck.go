package game

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

type Suit int

const (
	HEARTS Suit = iota
	DIAMONDS
	CLUBS
	SPADES
)

type Card struct {
	Value int
	Suit  Suit
}

type Deck struct {
	Cards []Card
}

func faceCardValue(v int64) string {
	if v > 10 {
		switch v {
		case 11:
			return "J"
		case 12:
			return "Q"
		case 13:
			return "K"
		case 14:
			return "A"
		}
	}
	return strconv.FormatInt(v, 10)
}

func (c Card) String() string {
	switch c.Suit {
	case 0:
		return faceCardValue(int64(c.Value)) + ":HEARTS"
	case 1:
		return faceCardValue(int64(c.Value)) + ":DIAMONDS"
	case 2:
		return faceCardValue(int64(c.Value)) + ":CLUBS"
	case 3:
		return faceCardValue(int64(c.Value)) + ":SPADES"
	}
	return "UNK"
}

func (deck *Deck) Shuffle() {
	cards := deck.Cards
	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))
	for i := range cards {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
}

func MakeDeck() *Deck {
	suites := []Suit{HEARTS, DIAMONDS, CLUBS, SPADES}
	values := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	deck := Deck{}
	for _, suite := range suites {
		for _, value := range values {
			deck.Cards = append(deck.Cards, Card{Suit: suite, Value: value})
		}
	}

	return &deck
}

func RemoveCard(cards []Card) []Card {
	return append(cards[:0], cards[1:]...)
}

func (deck *Deck) PopCard() (Card, error) {
	if len(deck.Cards) > 0 {
		c := deck.Cards[0]
		deck.Cards = RemoveCard(deck.Cards)
		return c, nil
	} else {
		return Card{}, errors.New("empty deck error")
	}
}

package game

import (
	"fmt"
	"sort"
)

type Player struct {
	Name string
	Hand []Card
}

func (p *Player) Take(c Card) {
	p.Hand = append(p.Hand, c)
}

func (p *Player) SortHand() {
	sort.Slice(p.Hand, func(i, j int) bool {
		return p.Hand[i].Value < p.Hand[j].Value
	})
}

func (p Player) GetHandEvaluated2() PokerHand {
	ph, _ := GetPokerHand(p.Hand)
	return ph
}

func (p Player) GetDiscards() []int {
	_, discards := GetPokerHand(p.Hand)
	return discards
}

func (p *Player) Discard(s []Card, c Card) []Card {

	for index, v := range p.Hand {
		if v.Value == c.Value && v.Suit == c.Suit {
			// Found!
			return append(s[:index], s[index+1:]...)
		}
	}
	fmt.Println("DEV ERROR: DISCARD did not find card")
	return []Card{}
}

func PokerString(pokerHand PokerHand) string {
	switch pokerHand {
	case HIGH_CARD:
		return "High Card"
	case PAIR:
		return "Pair"
	case TWO_PAIR:
		return "Two Pair"
	case TRIPS:
		return "Trips"
	case STRAIGHT:
		return "Straight"
	case FLUSH:
		return "Flush"
	case FULL_HOUSE:
		return "Full House"
	case QUAD:
		return "4 of a kind"
	case STRAIGHT_FLUSH:
		return "Straight Flush"
	case ROYAL_FLUSH:
		return "Royal Flush"
	}

	return "UNK"
}
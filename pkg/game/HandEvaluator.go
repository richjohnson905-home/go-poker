package game

import (
	"sort"
)

type PokerHand int

const (
	HIGH_CARD PokerHand = iota
	PAIR
	TWO_PAIR
	TRIPS
	STRAIGHT
	FLUSH
	FULL_HOUSE
	QUAD
	STRAIGHT_FLUSH
	ROYAL_FLUSH
)

func GetPokerHand(sortedHand []Card) (PokerHand, []int) {
	if isRoyal(sortedHand) {
		return ROYAL_FLUSH, []int{}
	} else if isStraightFlush(sortedHand) {
		return STRAIGHT_FLUSH, []int{}
	} else if isQuad(sortedHand) {
		return QUAD, []int{}
	} else if isStraightFlush(sortedHand) {
		return STRAIGHT_FLUSH, []int{}
	} else if isQuad(sortedHand) {
		return QUAD, []int{}
	} else if isFullHouse(sortedHand) {
		return FULL_HOUSE, []int{}
	} else if isFlush(sortedHand) {
		return FLUSH, []int{}
	} else if isStraight(sortedHand) {
		return STRAIGHT, []int{}
	} else if flag, a := isTrips(sortedHand); flag {
		return TRIPS, a
	} else if flag, s := isTwoPair(sortedHand); flag {
		return TWO_PAIR, s
	} else if flag, d := isPair(sortedHand); flag {
		return PAIR, d
	}
	highCardDiscards := getHighCardDiscards(sortedHand)
	return HIGH_CARD, highCardDiscards
}


func SortSuit(sortedHand []Card) {
	sort.Slice(sortedHand, func(i, j int) bool {
		return sortedHand[i].Suit < sortedHand[j].Suit
	})
}

func SortHand(sortedHand []Card) {
	sort.Slice(sortedHand, func(i, j int) bool {
		return sortedHand[i].Value < sortedHand[j].Value
	})
}
func getHighCardDiscards(sortedHand []Card) []int {
	SortSuit(sortedHand)

	firstFourFlush := (sortedHand[0].Suit == sortedHand[1].Suit) &&
		(sortedHand[0].Suit == sortedHand[2].Suit) &&
		(sortedHand[0].Suit == sortedHand[3].Suit)
	lastFourFlush := (sortedHand[1].Suit == sortedHand[2].Suit) &&
		(sortedHand[1].Suit == sortedHand[3].Suit) &&
		(sortedHand[1].Suit == sortedHand[4].Suit)
	if firstFourFlush {
		return []int{4}
	} else if lastFourFlush {
		return []int{0}
	}

	SortHand(sortedHand)
	firstFourStraight := (sortedHand[3].Value - sortedHand[2].Value == 1) &&
		(sortedHand[2].Value - sortedHand[1].Value == 1) &&
		(sortedHand[1].Value - sortedHand[0].Value == 1)
	lastFourStraight := (sortedHand[4].Value - sortedHand[3].Value == 1) &&
		(sortedHand[3].Value - sortedHand[2].Value == 1) &&
		(sortedHand[2].Value - sortedHand[1].Value == 1)
	if firstFourStraight {
		return []int{4}
	} else if lastFourStraight {
		return []int{0}
	}
	return []int{0, 1, 2}
}

func isRoyal(sortedHand []Card) bool {
	return isStraight(sortedHand) &&
		isFlush(sortedHand) &&
		sortedHand[0].Value == 10
}

func isFullHouse(sortedHand []Card) bool {
	return (are2CardsEqual(sortedHand[0], sortedHand[1]) &&
		are3CardsEqual(sortedHand[2], sortedHand[3], sortedHand[4])) ||
		(are2CardsEqual(sortedHand[3], sortedHand[4]) &&
			are3CardsEqual(sortedHand[0], sortedHand[1], sortedHand[2]))
}

func isStraightFlush(sortedHand []Card) bool {
	return isStraight(sortedHand) && isFlush(sortedHand)
}

func isQuad(sortedHand []Card) bool {
	return are4CardsEqual(
		sortedHand[0],
		sortedHand[1],
		sortedHand[2],
		sortedHand[3]) ||
		are4CardsEqual(
			sortedHand[1],
			sortedHand[2],
			sortedHand[3],
			sortedHand[4])
}

func isFlush(sortedHand []Card) bool {
	return (sortedHand[0].Suit == sortedHand[1].Suit) &&
		(sortedHand[0].Suit == sortedHand[2].Suit) &&
		(sortedHand[0].Suit == sortedHand[3].Suit) &&
		(sortedHand[0].Suit == sortedHand[4].Suit)
}

func isStraight(sortedHand []Card) bool {
	return (sortedHand[4].Value - sortedHand[3].Value == 1) &&
		(sortedHand[3].Value - sortedHand[2].Value == 1) &&
		(sortedHand[2].Value - sortedHand[1].Value == 1) &&
		(sortedHand[1].Value - sortedHand[0].Value == 1)
}

func isTrips(sortedHand []Card) (bool, []int) {
	firstThree := are3CardsEqual(sortedHand[0], sortedHand[1], sortedHand[2])
	middleThree := are3CardsEqual(sortedHand[1], sortedHand[2], sortedHand[3])
	lastThree := are3CardsEqual(sortedHand[2], sortedHand[3], sortedHand[4])
	if firstThree {
		return true, []int{3, 4}
	} else if middleThree {
		return true, []int{0, 4}
	} else if lastThree {
		return true, []int{0, 1}
	}
	return false, []int{}
}

func isTwoPair(sortedHand []Card) (bool, []int) {
	firstFour := are2CardsEqual(sortedHand[0], sortedHand[1]) && are2CardsEqual(sortedHand[2], sortedHand[3])
	firstLast := are2CardsEqual(sortedHand[0], sortedHand[1]) && are2CardsEqual(sortedHand[3], sortedHand[4])
	lastFour :=	are2CardsEqual(sortedHand[1], sortedHand[2]) && are2CardsEqual(sortedHand[3], sortedHand[4])
	if firstFour {
		return true, []int{4}
	} else if firstLast {
		return true, []int{2}
	} else if lastFour {
		return true, []int{0}
	}
	return false, []int{}
}

func isPair(sortedHand []Card) (bool, []int) {
	firstTwo := are2CardsEqual(sortedHand[0], sortedHand[1])
	secondTwo := are2CardsEqual(sortedHand[1], sortedHand[2])
	thirdTwo :=	are2CardsEqual(sortedHand[2], sortedHand[3])
	fourthTwo := are2CardsEqual(sortedHand[3], sortedHand[4])
	if firstTwo {
		return true, []int{2, 3, 4}
	} else if secondTwo {
		return true, []int{0, 3, 4}
	} else if thirdTwo {
		return true, []int{0, 1, 4}
	}else if fourthTwo {
		return true, []int{0, 1, 2}
	}
	return false, []int{}
}

func are2CardsEqual(c1 Card, c2 Card) bool {
	return c1.Value == c2.Value
}

func are3CardsEqual(c1 Card, c2 Card, c3 Card) bool {
	return c1.Value == c2.Value && c2.Value == c3.Value
}

func are4CardsEqual(c1 Card, c2 Card, c3 Card, c4 Card) bool {
	return c1.Value == c2.Value && c2.Value == c3.Value && c3.Value == c4.Value
}


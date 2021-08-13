package game

import "testing"

func TestGetPokerHand(t *testing.T) {
	type args struct {
		sortedHand []Card
	}
	tests := []struct {
		name string
		args args
		want PokerHand
	}{
		{"GetPair",
			args{
				[]Card{
					{Suit: DIAMONDS, Value: 10},
					{Suit: DIAMONDS, Value: 3},
					{Suit: CLUBS, Value: 3},
					{Suit: DIAMONDS, Value: 5},
					{Suit: DIAMONDS, Value: 6},
				},
			},
			PAIR},
		{"Get Royal Flush",
			args{
				[]Card{
					{Suit: DIAMONDS, Value: 10},
					{Suit: DIAMONDS, Value: 11},
					{Suit: DIAMONDS, Value: 12},
					{Suit: DIAMONDS, Value: 13},
					{Suit: DIAMONDS, Value: 14},
				},
			},
			ROYAL_FLUSH},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := GetPokerHand(tt.args.sortedHand); got != tt.want {
				t.Errorf("GetPokerHand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_are2CardsEqual(t *testing.T) {
	type args struct {
		c1 Card
		c2 Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Equal",
			args{
			Card{Suit: DIAMONDS, Value: 2},
			Card{Suit: HEARTS, Value: 2},
			},
			true},
		{"Not Equal",
			args{
			Card{Suit: DIAMONDS, Value: 2},
			Card{Suit: HEARTS, Value: 3},
			},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := are2CardsEqual(tt.args.c1, tt.args.c2); got != tt.want {
				t.Errorf("are2CardsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_are3CardsEqual(t *testing.T) {
	type args struct {
		c1 Card
		c2 Card
		c3 Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Equal",
			args{
				Card{Suit: DIAMONDS, Value: 2},
				Card{Suit: HEARTS, Value: 2},
				Card{Suit: HEARTS, Value: 2},
			},
			true},
		{"Not Equal",
			args{
				Card{Suit: DIAMONDS, Value: 2},
				Card{Suit: HEARTS, Value: 3},
				Card{Suit: HEARTS, Value: 3},
			},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := are3CardsEqual(tt.args.c1, tt.args.c2, tt.args.c3); got != tt.want {
				t.Errorf("are3CardsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isFlush(t *testing.T) {
	type args struct {
		sortedHand []Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"isFlush",
			args{
				 []Card{
				 	{Suit: DIAMONDS, Value: 2},
				 	{Suit: DIAMONDS, Value: 3},
				 	{Suit: DIAMONDS, Value: 4},
				 	{Suit: DIAMONDS, Value: 5},
				 	{Suit: DIAMONDS, Value: 6},
				 },
			},
			true},
		{"Not Flush",
			args{
				[]Card{
					{Suit: SPADES, Value: 2},
					{Suit: DIAMONDS, Value: 3},
					{Suit: DIAMONDS, Value: 4},
					{Suit: DIAMONDS, Value: 5},
					{Suit: DIAMONDS, Value: 6},
				},
			},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFlush(tt.args.sortedHand); got != tt.want {
				t.Errorf("isFlush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isFullHouse(t *testing.T) {
	type args struct {
		sortedHand []Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"isFullHouse",
			args{
				[]Card{
					{Suit: DIAMONDS, Value: 2},
					{Suit: SPADES, Value: 2},
					{Suit: CLUBS, Value: 2},
					{Suit: DIAMONDS, Value: 5},
					{Suit: HEARTS, Value: 5},
				},
			},
			true},
		{"Not FullHouse",
			args{
				[]Card{
					{Suit: SPADES, Value: 2},
					{Suit: DIAMONDS, Value: 3},
					{Suit: DIAMONDS, Value: 4},
					{Suit: DIAMONDS, Value: 5},
					{Suit: DIAMONDS, Value: 6},
				},
			},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFullHouse(tt.args.sortedHand); got != tt.want {
				t.Errorf("isFullHouse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPair(t *testing.T) {
	type args struct {
		sortedHand []Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"isPair",
			args{
				[]Card{
					{Suit: DIAMONDS, Value: 2},
					{Suit: SPADES, Value: 3},
					{Suit: CLUBS, Value: 4},
					{Suit: DIAMONDS, Value: 5},
					{Suit: HEARTS, Value: 5},
				},
			},
			true},
		{"Not Pair",
			args{
				[]Card{
					{Suit: SPADES, Value: 2},
					{Suit: DIAMONDS, Value: 3},
					{Suit: DIAMONDS, Value: 4},
					{Suit: DIAMONDS, Value: 5},
					{Suit: DIAMONDS, Value: 6},
				},
			},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := isPair(tt.args.sortedHand); got != tt.want {
				t.Errorf("isPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isQuad(t *testing.T) {
	type args struct {
		sortedHand []Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"isQuad",
			args{
				[]Card{
					{Suit: DIAMONDS, Value: 2},
					{Suit: SPADES, Value: 5},
					{Suit: CLUBS, Value: 5},
					{Suit: DIAMONDS, Value: 5},
					{Suit: HEARTS, Value: 5},
				},
			},
			true},
		{"Not Quad",
			args{
				[]Card{
					{Suit: SPADES, Value: 2},
					{Suit: DIAMONDS, Value: 3},
					{Suit: DIAMONDS, Value: 4},
					{Suit: DIAMONDS, Value: 5},
					{Suit: DIAMONDS, Value: 6},
				},
			},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isQuad(tt.args.sortedHand); got != tt.want {
				t.Errorf("isQuad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isRoyal(t *testing.T) {
	type args struct {
		sortedHand []Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"isRoyal",
			args{
				[]Card{
					{Suit: DIAMONDS, Value: 10},
					{Suit: DIAMONDS, Value: 11},
					{Suit: DIAMONDS, Value: 12},
					{Suit: DIAMONDS, Value: 13},
					{Suit: DIAMONDS, Value: 14},
				},
			},
			true},
		{"Not Royal",
			args{
				[]Card{
					{Suit: SPADES, Value: 2},
					{Suit: DIAMONDS, Value: 3},
					{Suit: DIAMONDS, Value: 4},
					{Suit: DIAMONDS, Value: 5},
					{Suit: DIAMONDS, Value: 6},
				},
			},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRoyal(tt.args.sortedHand); got != tt.want {
				t.Errorf("isRoyal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isStraight(t *testing.T) {
	type args struct {
		sortedHand []Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"isStraight",
			args{
				[]Card{
					{Suit: SPADES, Value: 3},
					{Suit: SPADES, Value: 4},
					{Suit: DIAMONDS, Value: 5},
					{Suit: DIAMONDS, Value: 6},
					{Suit: DIAMONDS, Value: 7},
				},
			},
			true},
		{"Not Straight",
			args{
				[]Card{
					{Suit: SPADES, Value: 2},
					{Suit: DIAMONDS, Value: 3},
					{Suit: DIAMONDS, Value: 4},
					{Suit: DIAMONDS, Value: 5},
					{Suit: DIAMONDS, Value: 7},
				},
			},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStraight(tt.args.sortedHand); got != tt.want {
				t.Errorf("isStraight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isStraightFlush(t *testing.T) {
	type args struct {
		sortedHand []Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"isStraightFlush",
			args{
				[]Card{
					{Suit: DIAMONDS, Value: 3},
					{Suit: DIAMONDS, Value: 4},
					{Suit: DIAMONDS, Value: 5},
					{Suit: DIAMONDS, Value: 6},
					{Suit: DIAMONDS, Value: 7},
				},
			},
			true},
		{"Not StraightFlush",
			args{
				[]Card{
					{Suit: SPADES, Value: 2},
					{Suit: DIAMONDS, Value: 3},
					{Suit: DIAMONDS, Value: 4},
					{Suit: DIAMONDS, Value: 5},
					{Suit: DIAMONDS, Value: 6},
				},
			},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStraightFlush(tt.args.sortedHand); got != tt.want {
				t.Errorf("isStraightFlush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isTrips(t *testing.T) {
	type retVal struct {
		flag bool
		discards []int
	}
	type args struct {
		sortedHand []Card
	}
	tests := []struct {
		name string
		args args
		want retVal
	}{
		{"isTrips_FirstThree",
			args{
				[]Card{
					{Suit: DIAMONDS, Value: 4},
					{Suit: DIAMONDS, Value: 4},
					{Suit: DIAMONDS, Value: 4},
					{Suit: DIAMONDS, Value: 3},
					{Suit: DIAMONDS, Value: 7},
				},
			},
			retVal{true, []int{3, 4}}},
		{"isTrips_middleThree",
			args{
				[]Card{
					{Suit: DIAMONDS, Value: 3},
					{Suit: DIAMONDS, Value: 4},
					{Suit: DIAMONDS, Value: 4},
					{Suit: DIAMONDS, Value: 4},
					{Suit: DIAMONDS, Value: 7},
				},
			},
			retVal{true, []int{0, 4}}},
		{"isTrips_lastThree",
			args{
				[]Card{
					{Suit: DIAMONDS, Value: 3},
					{Suit: DIAMONDS, Value: 7},
					{Suit: DIAMONDS, Value: 4},
					{Suit: DIAMONDS, Value: 4},
					{Suit: DIAMONDS, Value: 4},
				},
			},
			retVal{true, []int{0, 1}}},
		{"Not Trips",
			args{
				[]Card{
					{Suit: SPADES, Value: 2},
					{Suit: DIAMONDS, Value: 3},
					{Suit: DIAMONDS, Value: 4},
					{Suit: DIAMONDS, Value: 5},
					{Suit: DIAMONDS, Value: 6},
				},
			},
			retVal{false, []int{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag, discards := isTrips(tt.args.sortedHand)
			if flag != tt.want.flag {
				t.Errorf("isTrips() = %v, want %v", flag, tt.want.flag)
			}
			if len(discards) != len(tt.want.discards) {
				t.Errorf("isTrips() = %v, want %v", len(discards), len(tt.want.discards))
			}
		})
	}
}

func Test_isTwoPair(t *testing.T) {
	type args struct {
		sortedHand []Card
	}
	type retVal struct {
		flag bool
		discards []int
	}
	tests := []struct {
		name string
		args args
		want retVal
	}{
		{"isTwoPair_firstFour",
			args{
				[]Card{
					{Suit: DIAMONDS, Value: 4},
					{Suit: DIAMONDS, Value: 4},
					{Suit: DIAMONDS, Value: 7},
					{Suit: DIAMONDS, Value: 7},
					{Suit: DIAMONDS, Value: 3},
				},
			},
			retVal{true, []int{4}}},
		{"isTwoPair_firstLast",
			args{
				[]Card{
					{Suit: DIAMONDS, Value: 4},
					{Suit: DIAMONDS, Value: 4},
					{Suit: DIAMONDS, Value: 3},
					{Suit: DIAMONDS, Value: 7},
					{Suit: DIAMONDS, Value: 7},
				},
			},
			retVal{true, []int{2}}},
		{"isTwoPair_lastFour",
			args{
				[]Card{
					{Suit: DIAMONDS, Value: 3},
					{Suit: DIAMONDS, Value: 4},
					{Suit: DIAMONDS, Value: 4},
					{Suit: DIAMONDS, Value: 7},
					{Suit: DIAMONDS, Value: 7},
				},
			},
			retVal{true, []int{0}}},
		{"Not TwoPair",
			args{
				[]Card{
					{Suit: SPADES, Value: 2},
					{Suit: DIAMONDS, Value: 3},
					{Suit: DIAMONDS, Value: 4},
					{Suit: DIAMONDS, Value: 5},
					{Suit: DIAMONDS, Value: 6},
				},
			},
			retVal{false, []int{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag, discards := isTwoPair(tt.args.sortedHand)
			if flag != tt.want.flag {
				t.Errorf("isTwoPair() = %v, want %v", flag, tt.want.flag)
			}
			if len(discards) != len(tt.want.discards) {
				t.Errorf("isTwoPair() = %v, want %v", len(discards), len(tt.want.discards))
			}
		})
	}
}

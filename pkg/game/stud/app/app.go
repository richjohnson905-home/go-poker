package app

import (
	"example.com/pkg/game"
	"fmt"
)

func RunGame(p1 game.Player, p2 game.Player) {
	fmt.Println("5 CARD STUD POKER")
	pokerHandFrequencyMap := make(map[game.PokerHand]float64)
	rounds := 100000
	for i := 0; i < rounds; i++ {
		playRound(&p1, &p2, &pokerHandFrequencyMap)
	}
	game.PrettyPrintResult(&pokerHandFrequencyMap, float64(rounds * 10.0))
}

func playRound(p1 *game.Player, p2 *game.Player, pokerMap *map[game.PokerHand]float64) {
	players := make([]*game.Player, 2)
	players[0] = p1
	players[1] = p2
	deck := game.MakeDeck()
	deck.Shuffle()
	dealer := game.MakeDealer2(players, deck)
	r := game.Round{players, dealer}
	for i := 0; i < 5; i++ {
		r.InitialDeal()
		ph1 := r.Eval(p1)
		ph2 := r.Eval(p2)
		(*pokerMap)[ph1]++
		(*pokerMap)[ph2]++
		p1.Hand = []game.Card{}
		p2.Hand = []game.Card{}
	}
}

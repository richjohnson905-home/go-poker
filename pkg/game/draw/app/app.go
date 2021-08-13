package app

import (
	"example.com/pkg/game"
	"fmt"
)

func RunGame(p1 game.Player, p2 game.Player) {
	fmt.Println("5-CARD DRAW POKER")
	pokerHandFrequencyMap := make(map[game.PokerHand]float64)
	rounds := 166666
	for i := 0; i < rounds; i++ {
		playRound(&p1, &p2, &pokerHandFrequencyMap)
	}
	game.PrettyPrintResult(&pokerHandFrequencyMap, float64(rounds * 6.0))
}

func playRound(p1 *game.Player, p2 *game.Player, pokerMap *map[game.PokerHand]float64) {
	players := make([]*game.Player, 2)
	players[0] = p1
	players[1] = p2
	deck := game.MakeDeck()
	deck.Shuffle()
	dealer := game.MakeDealer2(players, deck)
	r := game.Round{players, dealer}
	for i := 0; i < 3; i++ {
		//fmt.Println("NEW ROUND")
		r.InitialDeal()

		r.Draw(r.Players[0], dealer)
		ph1 := r.Eval(r.Players[0])
		(*pokerMap)[ph1]++
		r.Players[0].Hand = []game.Card{}

		r.Draw(r.Players[1], dealer)
		ph2 := r.Eval(r.Players[1])
		(*pokerMap)[ph2]++
		r.Players[1].Hand = []game.Card{}
	}
}
package game

import "fmt"

type Round struct {
	Players []*Player
	Dealer *Dealer2
}

func(r Round) InitialDeal() {
	r.Dealer.Deal2(r.Players[0], r.Players[1])
	for i := range r.Players {
		r.Players[i].SortHand()
	}
}

func(r Round) Draw(p *Player, dealer *Dealer2) {

	discards := p.GetDiscards()
	cardsToDelete := make([]Card, len(discards))
	for idx, d := range discards {
		cardsToDelete[idx] = p.Hand[d]
	}
	for _, c := range cardsToDelete {
		p.Hand = p.Discard(p.Hand, c)
	}
	for i := 0; i < len(discards); i++ {
		p.Take(dealer.OneCard())
	}
	p.SortHand()
}

func (r Round) Eval(p *Player) PokerHand {
	ph := p.GetHandEvaluated2()
	if ph > QUAD {
		fmt.Println("Before Player:", p.Name, "Hand:", p.Hand)
	}
	return ph
}

func PrettyPrintResult(pokerMap *map[PokerHand]float64, total float64) {
	//fmt.Println("HS: ", (*pokerMap)[HIGH_CARD])
	fmt.Println("Total Hands Played:\t", total)
	fmt.Printf("High Card \t%f\n", 		((*pokerMap)[HIGH_CARD] / total) * 100)
	fmt.Printf("Pair \t\t%f\n", 			((*pokerMap)[PAIR] / total) * 100)
	fmt.Printf("Two Pair \t%f\n", 		((*pokerMap)[TWO_PAIR] / total) * 100)
	fmt.Printf("Trips \t\t%f\n", 		((*pokerMap)[TRIPS] / total) * 100)
	fmt.Printf("Straight \t%f\n", 		((*pokerMap)[STRAIGHT] / total) * 100)
	fmt.Printf("Flush \t\t%f\n", 		((*pokerMap)[FLUSH] / total) * 100)
	fmt.Printf("Full House \t%f\n", 		((*pokerMap)[FULL_HOUSE] / total) * 100)
	fmt.Printf("4 of a kind \t%f\n", 	((*pokerMap)[QUAD] / total) * 100)
	fmt.Printf("Straight Flush \t%f\n", 	((*pokerMap)[STRAIGHT_FLUSH] / total) * 100)
	fmt.Printf("Royal Flush \t%f\n", 	((*pokerMap)[ROYAL_FLUSH] / total) * 100)
}
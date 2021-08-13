package game

type Dealer2 struct {
	Players []*Player
	Deck    *Deck
}

func MakeDealer2(players []*Player, deck *Deck) *Dealer2 {
	return &Dealer2 {
		Players: players,
		Deck: deck,
	}
}

func (d Dealer2) Deal2(p1 *Player, p2 *Player) {
	i := 0
	for i < 5 {
		for _, p := range d.Players {
			c, _ := d.Deck.PopCard()
			p.Take(c)
		}
		i++
	}
}

func (dealer Dealer2) OneCard() Card {
	//fmt.Println("DECK SIZE:", len(dealer.Deck.Cards))
	c, _ := dealer.Deck.PopCard()
	return c
}
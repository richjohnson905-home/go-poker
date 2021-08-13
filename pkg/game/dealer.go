package game

type Dealer struct {
	Players []GamePlayer
	Deck    GameDeck
}

type GamePlayer interface {
	Take(c Card)
}

type GameDeck interface {
	PopCard() (Card, error)
}

func MakeDealer(players []GamePlayer, deck GameDeck) *Dealer {
	return &Dealer {
		Players: players,
		Deck: deck,
	}
}

func (dealer Dealer) Deal() {
	i := 0
	for i < 5 {
		for _, p := range dealer.Players {
			c, _ := dealer.Deck.PopCard()
			p.Take(c)
		}
		i++
	}
}



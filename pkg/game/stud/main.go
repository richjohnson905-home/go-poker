package main

import (
	"example.com/pkg/game"
	app2 "example.com/pkg/game/stud/app"
)

func main() {
	p1 := game.Player{"Bob", []game.Card{}}
	p2 := game.Player{"Larry", []game.Card{}}
	app2.RunGame(p1, p2)
}
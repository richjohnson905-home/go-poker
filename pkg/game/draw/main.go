package main

import (
	"example.com/pkg/game"
	"example.com/pkg/game/draw/app"
)

func main() {
	p1 := game.Player{"Bill", []game.Card{}}
	p2 := game.Player{"Tony", []game.Card{}}
	app.RunGame(p1, p2)
}
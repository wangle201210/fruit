package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.NewWithID("iwangle.me.fruit")
	game := newSlotGame(a)

	game.run()
}

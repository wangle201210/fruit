package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"testing"
)

func TestSpin(t *testing.T) {
	a := app.New()
	w := a.NewWindow("fruit 777")
	game := newSlotGame(w)
	_ = game.panel() // 面板
	winCount := 0
	award := 0
	times := 2000000
	creditInit, _ := game.userInfo.credit.Get()

	for i := 0; i < times; i++ {
		game.startSpin()
		award, _ = game.userInfo.award.Get()
		if award > 0 {
			winCount++
		}
	}
	credit, _ := game.userInfo.credit.Get()
	bet, _ := game.userInfo.bet.Get()
	fmt.Printf("winCount:%v,winRt:%v, credit:%v,creditInit%v, rtp:%v",
		winCount, float64(winCount)/float64(times), credit, creditInit, float64(credit-creditInit+times*bet)/float64(times*bet))
}

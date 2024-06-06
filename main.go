package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.NewWithID("iwangle.me.fruit")
	w := a.NewWindow("fruit 777")
	game := newSlotGame(w)
	panel := game.panel()   // 面板
	header := game.header() // 头部
	footer := game.footer() // 页脚
	history := game.history()
	a.Lifecycle().SetOnStopped(func() {
		game.saveData() // 退出时保存数据
	})
	w.SetContent(container.NewVBox(
		header,
		panel,
		footer,
		history,
	))
	w.Resize(fyne.NewSize(800, 800))
	w.ShowAndRun()
}

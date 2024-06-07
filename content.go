package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func (g *slotGame) panel() *fyne.Container {
	grid := container.NewGridWithColumns(columns)
	for colIdx := 0; colIdx < columns; colIdx++ {
		col := container.NewVBox()
		for rowIdx := 0; rowIdx < rows; rowIdx++ {
			col.Add(g.labels[rowIdx][colIdx])
		}
		grid.Add(col)
	}
	g.spin() // 初始面板
	return grid
}

func (g *slotGame) header() *fyne.Container {
	creditLabel := widget.NewLabelWithData(binding.IntToStringWithFormat(g.userInfo.credit, "credit: %d"))
	awardLabel := widget.NewLabelWithData(binding.IntToStringWithFormat(g.userInfo.award, "award: %d"))
	payButton := widget.NewButton("pay table", func() {
		g.payTable.Show()
	})
	hisButton := widget.NewButton("history", func() {
		g.history.Show()
	})
	header := container.NewHBox(
		creditLabel,
		layout.NewSpacer(),
		awardLabel,
		payButton,
		hisButton,
	)
	return header
}

func (g *slotGame) footer() *fyne.Container {
	spinButton := widget.NewButton("Spin", func() {
		g.startSpin()
	})
	bl := widget.NewSelect(betList, func(s string) {
		i, _ := strconv.Atoi(s)
		g.userInfo.bet.Set(i)
	})
	bl.SetSelectedIndex(0)
	bl.Resize(bl.Size().AddWidthHeight(100, 0))
	button := container.New(&customLayout{}, spinButton)
	text := canvas.NewText("  bet:", nil)
	footer := container.NewHBox(
		text,
		bl,
		layout.NewSpacer(),
		button,
	)
	return footer
}

func (g *slotGame) payTableInit() {
	pt := g.app.NewWindow("pay table")
	list := widget.NewTable(
		func() (int, int) {
			return len(symbols), 2
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Col == 0 {
				o.(*widget.Label).SetText(symbols[i.Row])
				o.(*widget.Label).TextStyle = fyne.TextStyle{Bold: true}
			} else {
				o.(*widget.Label).SetText(fmt.Sprintf("%d * bet / %d", awardList[symbols[i.Row]], len(lines)))
			}
		})
	list.SetColumnWidth(0, 100) // 设置第一列宽度为100
	list.SetColumnWidth(1, 100) // 设置第一列宽度为100

	pt.SetContent(list)
	pt.Resize(fyne.NewSize(300, 300))
	pt.SetCloseIntercept(func() {
		pt.Hide()
	})
	g.payTable = pt
}

func (g *slotGame) historyInit() {
	hisWin := g.app.NewWindow("history")
	g.history = hisWin
	his := widget.NewTable(
		func() (int, int) {
			return len(g.userInfo.history), len(g.userInfo.history[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Row == 0 {
				o.(*widget.Label).SetText(g.userInfo.history[i.Row][i.Col])
			} else {
				o.(*widget.Label).SetText(g.userInfo.history[len(g.userInfo.history)-i.Row][i.Col])
			}
		})
	his.SetColumnWidth(0, 100) // 设置第一列宽度为100
	his.SetColumnWidth(1, 80)  // 设置第二列宽度为150
	his.SetColumnWidth(2, 80)  // 设置第三列宽度为200
	scrollContainer := container.NewVScroll(his)
	scrollContainer.SetMinSize(fyne.NewSize(300, 300))
	hisWin.Resize(fyne.NewSize(300, 300))
	hisWin.SetCloseIntercept(func() {
		hisWin.Hide()
	})
	hisWin.SetContent(scrollContainer)
}

package main

import (
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
	header := container.NewHBox(
		creditLabel,
		layout.NewSpacer(),
		awardLabel,
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
	footer.Resize(footer.Size().AddWidthHeight(0, 100))
	footer.Move(fyne.NewPos(0, 10))
	return footer
}

func (g *slotGame) history() *fyne.Container {
	his := widget.NewTable(
		func() (int, int) {
			return len(g.userInfo.history), len(g.userInfo.history[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("宽内容")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(g.userInfo.history[i.Row][i.Col])
		})
	list := container.New(&listLayout{}, his)
	con := container.NewVBox(list)
	return con
}

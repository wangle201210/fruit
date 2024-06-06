package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var (
	buttonWidth = float32(200)
)

// 定义一个自定义布局
type customLayout struct{}

func (c *customLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if len(objects) == 0 {
		return
	}
	// 假设我们希望设置按钮的宽度为 200 像素
	// 计算按钮的位置
	x := float32(0)
	for _, obj := range objects {
		switch obj.(type) {
		case *widget.Button:
			obj.Resize(fyne.NewSize(buttonWidth, size.Height))
			// obj.MinSize() = obj.Size()
		}
		obj.Move(fyne.NewPos(x, 0))
		x += obj.Size().Width
	}
}

func (c *customLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	width := float32(0)
	height := float32(0)

	for _, obj := range objects {
		childSize := obj.MinSize()
		width += childSize.Width
		if childSize.Height > height {
			height = childSize.Height
		}
	}
	width = buttonWidth + 10
	height = 50
	return fyne.NewSize(width, height)
}

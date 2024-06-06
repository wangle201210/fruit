package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var (
	listHeight = float32(200)
)

// 定义一个自定义布局
type listLayout struct{}

func (c *listLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if len(objects) == 0 {
		return
	}
	// 假设我们希望设置按钮的宽度为 200 像素
	// 计算按钮的位置
	// x := float32(0)
	for _, obj := range objects {
		switch obj.(type) {
		case *widget.Button:
			obj.Resize(fyne.NewSize(size.Width, listHeight))
		}
		// obj.Move(fyne.NewPos(x, 0))
		// x += obj.Size().Width
	}
}

func (c *listLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	width := float32(0)
	height := float32(0)

	for _, obj := range objects {
		childSize := obj.MinSize()
		width += childSize.Width
	}
	height = listHeight
	return fyne.NewSize(width, height)
}

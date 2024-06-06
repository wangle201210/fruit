package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
	"strconv"
	"sync"
	"time"
)

const (
	rows    = 3
	columns = 3
)

var symbols = []string{"🍎", "🍒", "🍋", "🍊", "🍉", "⭐", "7"}
var awardList = map[string]int{
	"🍎": 5, "🍒": 10, "🍋": 20, "🍊": 40, "🍉": 80, "⭐": 200, "7": 500,
}
var reels = []int{0, 1, 2, 0, 1, 2, 3, 1, 4, 5, 1, 6, 0, 1, 2, 1, 2, 0, 1, 2, 3, 1, 4, 2, 5, 2, 1, 6}
var weight = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

var betList = []string{"7", "14", "21", "28", "35", "42"}

var (
	lines = [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 4, 8},
		{2, 4, 6},
		{0, 7, 2},
		{6, 1, 8},
	}
)

var (
	debug           = false
	reelDelay       = time.Millisecond * 300
	calcDelay       = time.Millisecond * 200
	winLineDelay    = time.Millisecond * 200
	spinColumnDelay = time.Millisecond * 50
)

type slotGame struct {
	sync.Mutex
	windows  fyne.Window
	labels   [][]*canvas.Text
	wg       *sync.WaitGroup
	userInfo *userInfo
}

type userInfo struct {
	award   binding.Int
	credit  binding.Int
	bet     binding.Int
	history [][]string
}

func newSlotGame(w fyne.Window) *slotGame {
	game := &slotGame{
		windows: w,
		labels:  make([][]*canvas.Text, rows),
	}
	game.loadData() // 加载以前的数据
	for i := 0; i < rows; i++ {
		game.labels[i] = make([]*canvas.Text, columns)
		for j := 0; j < columns; j++ {
			game.labels[i][j] = canvas.NewText("", nil)
		}
	}

	return game
}

func (g *slotGame) startSpin() {
	g.spin()
	// 结算
	g.settle()
}

func (g *slotGame) spin() {
	g.Lock()
	defer g.Unlock()
	g.wg = &sync.WaitGroup{}
	for i := 0; i < columns; i++ {
		g.wg.Add(1)
		time.Sleep(reelDelay)
		go g.spinColumn(i)
	}
	g.wg.Wait()
}

func (g *slotGame) settle() {
	award := g.calc()
	credit, _ := g.userInfo.credit.Get()
	bet, _ := g.userInfo.bet.Get()
	g.userInfo.award.Set(award)
	g.userInfo.credit.Set(credit + award - bet)
	g.userInfo.history = append(g.userInfo.history, []string{strconv.Itoa(award), strconv.Itoa(credit), strconv.Itoa(bet)})
}

func (g *slotGame) calc() (award int) {
	time.Sleep(calcDelay)
	for _, line := range lines {
		x, y := indexToXY(line[0])
		icon := g.labels[x][y].Text
		num := 0
		for _, i := range line {
			x, y = indexToXY(i)
			if g.labels[x][y].Text == icon {
				num++
			}
		}
		if num == 3 {
			bet, _ := g.userInfo.bet.Get()
			award += bet * awardList[icon] / len(lines)
			// println(icon, award, bet, awardList[icon], len(lines))
			g.winLine(line)
		}
	}
	return
}

func (g *slotGame) winLine(line []int) {
	idx := line[0]
	x, y := indexToXY(idx)
	text := g.labels[x][y].Text
	for j := 0; j < 3; j++ {
		for _, i := range line {
			x, y = indexToXY(i)
			g.labels[x][y].Text = " "
		}
		g.windows.Content().Refresh()
		time.Sleep(winLineDelay)
		for _, i := range line {
			x, y = indexToXY(i)
			g.labels[x][y].Text = text
		}
		g.windows.Content().Refresh()
		time.Sleep(winLineDelay)
	}
}

func (g *slotGame) spinColumn(col int) {
	defer g.wg.Done()
	for i := 0; i < 20; i++ { // 简单的轮数控制
		reelIdx := randIndex()
		for j := 0; j < rows; j++ {
			iconIndex := (reelIdx + j) % len(reels)
			g.labels[j][col].Text = symbols[reels[iconIndex]]
			textStyle(g.labels[j][col])
		}
		time.Sleep(spinColumnDelay)
	}
}

func randIndex() int {
	reelIdx := WeightedChoice(weight)
	return reelIdx
}

func textStyle(text *canvas.Text) {
	text.TextSize = 50
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	text.Refresh()
}
func indexToXY(index int) (x, y int) {
	x = index / columns
	y = index % columns
	return
}

func init() {
	if debug {
		reelDelay = 0
		calcDelay = 0
		winLineDelay = 0
		spinColumnDelay = 0
	}
}

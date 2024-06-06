package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"io"
	"os"
	"path"
)

const fileName = "data.json"

type data struct {
	Award  int
	Credit int
	Bet    int
}

func user2data(info *userInfo) (res *data) {
	res = new(data)
	res.Bet, _ = info.bet.Get()
	res.Credit, _ = info.credit.Get()
	res.Award, _ = info.award.Get()
	return
}

func data2user(res *data) (info *userInfo) {
	if res.Bet == 0 {
		res.Bet = 5
	}
	if res.Credit == 0 {
		res.Credit = 1e5
	}
	info = &userInfo{
		credit: binding.NewInt(),
		bet:    binding.NewInt(),
		award:  binding.NewInt(),
	}
	// info.award.Set(res.Award)
	info.credit.Set(res.Credit)
	info.bet.Set(res.Bet)
	info.history = append(info.history, []string{"bet", "award", "credit"})
	return
}

// saveData 用于保存数据
func (g *slotGame) saveData() {
	res := user2data(g.userInfo)
	marshal, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Marshal err:", err)
	}
	file, err := os.Create(getDataPath())
	if err != nil {
		fmt.Println("无法创建文件:", err)
		return
	}
	defer file.Close()
	file.WriteString(string(marshal))
}

// loadData 用于加载数据
func (g *slotGame) loadData() {
	res := new(data)
	defer func() {
		g.userInfo = data2user(res) // 最后赋值就行
	}()
	file, err := os.Open(getDataPath())
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()
	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("读取文件错误:", err)
		return
	}
	if len(byteValue) > 0 {
		err = json.Unmarshal(byteValue, res)
		if err != nil {
			fmt.Println("Unmarshal err:", err)
			return
		}
	}
}

func getDataPath() string {
	storageRootURI := fyne.CurrentApp().Storage().RootURI()
	println(storageRootURI.Path())
	return path.Join(storageRootURI.Path(), fileName)
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

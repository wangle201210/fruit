# fruit 777
## 简介
使用fyne做的简易版水果机  
亲测 安卓、mac、windows 均可正常使用
![img.png](content.png)
## 运行及打包
本地运行：`go run .`  
下面三个都需要提前安装依赖
安卓打包：`fyne package -os android -icon icon.png -appID iwangle.me.fruit`  
windows: `fyne-cross windows -app-id iwangle.me.fruit`  
mac: `fyne-cross darwin -app-id iwangle.me.fruit`  

## 注意事项
构建安卓应用时若出现
`../../../go/pkg/mod/fyne.io/fyne/v2@v2.4.5/internal/driver/mobile/app/android.go:520:6: could not determine kind of name for C.ALooper_pollAll
`需要找到对应行把`ALooper_pollAll`改为`ALooper_pollOnce`

## 功能列表
- [x] 基础界面
- [x] 滚动效果
- [x] 关闭后再次打开能读到上次数据
- [x] 奖线播放
- [x] 数值调控 rtp: 0.96, winRate 0.5+
- [x] 中奖历史记录
- [x] 界面可查看奖励配置
- [ ] 添加音效
- [ ] 添加背景 (加了发现不好看)
- [ ] 自动spin

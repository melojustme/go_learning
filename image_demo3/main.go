package main

import (
	"strconv"
	"time"

	"github.com/fogleman/gg"
)

func main() {
	var p []Platform
	var pf Platform

	pf.PlatformName = "哈哈哈哈哈哈"
	pf.RegisterCount = 11112333
	p = append(p, pf)
	pf.PlatformName = "哈哈哈哈哈哈"
	pf.RegisterCount = 11112333
	p = append(p, pf)
	pf.PlatformName = "哈哈哈哈哈哈"
	pf.RegisterCount = 11112333
	p = append(p, pf)
	pf.PlatformName = "哈哈哈哈哈哈"
	pf.RegisterCount = 11112333
	p = append(p, pf)
	pf.PlatformName = "哈哈哈哈哈哈"
	pf.RegisterCount = 11112333
	p = append(p, pf)
	PlatformWechatImage("", p)
}

func PlatformDrawLine(dc *gg.Context, w, h float64, platformName, registerCount string) (float64, float64) {
	dc.DrawString(platformName, w, h)
	w += 300
	dc.DrawString(registerCount, w, h)
	h += 40
	w = 100
	return w, h
}

// 注册量
type Platform struct {
	Id            int
	PlatformName  string
	RegisterCount int
}

func PlatformWechatImage(touser string, registerList []Platform) {
	dc := gg.NewContext(500, 200+len(registerList)*40)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("./字体管家淘淘体.ttf", 18); err != nil {
		return
	}
	w := 100.0
	h := 50.0
	dc.DrawString("时间:"+time.Now().Format("2006-01-02 15:04:05"), w, h)
	h += 60
	w, h = PlatformDrawLine(dc, w, h, "名称", "数量")

	var registerTotal int
	for _, v := range registerList {
		registerTotal += v.RegisterCount
		platformName := ""
		if v.PlatformName == "" {
			platformName = "未获取名称"
		} else {
			platformName = v.PlatformName
		}
		w, h = PlatformDrawLine(dc, w, h, platformName, strconv.Itoa(v.RegisterCount))
	}
	PlatformDrawLine(dc, w, h, "合计", strconv.Itoa(registerTotal))

	//保存图片
	dc.SavePNG("platform.png")

}

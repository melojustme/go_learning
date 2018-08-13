package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang/freetype"
)

func main() {
	var s Signer

	s.FontSize = 20 // 字体尺寸
	s.Dpi = 72      ///Users/goworkspace/mygo/src/go_learning/i1/方正苏新诗柳楷简体.TTF
	s.fontPath = "方正苏新诗柳楷简体.TTF"
	text := `
	生活可以漂泊，可以孤独，但灵魂必须有所归依。\r\n
	对未来真正的慷慨，就是把一切都献给现在。\r\n
	年轻是一种资源，但不努力就浪费了。\r\n
	斩断自己的退路，才能更好地赢得出路。\r\n
	一见钟情，钟的不是情，是脸。\r\n
	世上只有骗子是真心的，因为他是真心骗你的。\r\n
	我要让全世界的人都知道，我很低调。\r\n
	唐僧再厉害，也不过是一个耍猴的。\r\n
	一句话欢喜，一句话泪流，都是你。\r\n
	`
	img, err := s.DrawStringImage(text)
	// 以PNG格式保存文件
	imgcounter := 123
	imgfile, _ := os.Create(fmt.Sprintf("%03d.png", imgcounter))
	err = png.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}
}

type Signer struct {
	fontPath string
	Dpi      float64
	FontSize float64
}

// 画一个带有text的图片
func (this *Signer) DrawStringImage(text string) (image.Image, error) {
	fontBytes, err := ioutil.ReadFile(this.fontPath)
	if err != nil {
		return nil, err
	}

	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}

	fg, bg := image.White, image.Black
	rgba := image.NewRGBA(image.Rect(0, 0, 900, 900))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)

	c := freetype.NewContext()
	c.SetDPI(this.Dpi)
	c.SetFont(font)
	c.SetFontSize(this.FontSize)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)

	// Draw the text.
	pt := freetype.Pt(10, 10+int(c.PointToFixed(12)>>8))
	for _, s := range strings.Split(text, "\r\n") {
		_, err = c.DrawString(s, pt)
		pt.Y += c.PointToFixed(12 * 1.5)
	}

	return rgba, nil
}

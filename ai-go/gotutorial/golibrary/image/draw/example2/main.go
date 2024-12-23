/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 11:48:37
*/
package main

import (
	"bufio"
	"flag"
	"github.com/golang/freetype"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func Drawing(text, text1, text2 []string) {

	var (
		dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
		fontfile = flag.String("fontfile", "./src/sms_v1/etc/阿丽达黑体.ttf", "filename of the ttf font")
		hinting  = flag.String("hinting", "none", "none | full")
		size     = flag.Float64("size", 20, "font size in points")
		spacing  = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
		wonb     = flag.Bool("whiteonblack", false, "white text on a black background")
	)

	flag.Parse()

	// Read the font data.
	fontBytes, err := os.ReadFile(*fontfile)
	if err != nil {
		log.Println(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	// Initialize the context.
	fg, bg := image.Black, image.White
	//ruler := color.RGBA{220, 220, 220, 200}
	if *wonb {
		fg, bg = image.White, image.Black
		//ruler = color.RGBA{220, 220, 220, 200}
	}

	// f1, err := os.Open("aa.png")
	// if err != nil {
	//    panic(err)
	// }
	// defer f1.Close()
	// m1, err := png.Decode(f1)
	//if err != nil {
	// panic(err)
	//}
	//bounds := m1.Bounds()

	rgba := image.NewRGBA(image.Rect(0, 0, 640, 480))

	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	//draw.DrawMask(rgba, bounds, bg, image.ZP,, draw.Src)

	c := freetype.NewContext()
	c.SetDPI(*dpi)
	c.SetFont(f)
	c.SetFontSize(*size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	switch *hinting {
	default:
		c.SetHinting(font.HintingNone)
	case "full":
		c.SetHinting(font.HintingFull)
	}

	// Draw the guidelines.
	//for y := 0; y < 480; y++ {
	//      for x := 0; x < 640; x++ {
	//                // 设置某个点的颜色，依次是 RGBA
	//                 rgba.Set(x, y, color.RGBA{uint8(x), uint8(y), 0, 255})
	//            }
	//      }

	for y := 100; y < 130; y++ {
		for x := 10; x < 620; x++ {
			// 设置某个点的颜色，依次是 RGBA
			rgba.Set(x, y, color.RGBA{190, 190, 190, 255})
		}
	}

	for y := 130; y < 160; y++ {
		for x := 10; x < 620; x++ {
			// 设置某个点的颜色，依次是 RGBA
			rgba.Set(x, y, color.RGBA{28, 172, 255, 255})
		}
	}

	for y := 160; y < 190; y++ {
		for x := 10; x < 620; x++ {
			// 设置某个点的颜色，依次是 RGBA
			rgba.Set(x, y, color.RGBA{190, 190, 190, 255})
		}
	}

	for y := 190; y < 220; y++ {
		for x := 10; x < 620; x++ {
			// 设置某个点的颜色，依次是 RGBA
			rgba.Set(x, y, color.RGBA{28, 172, 255, 255})
		}
	}

	for y := 220; y < 250; y++ {
		for x := 10; x < 620; x++ {
			// 设置某个点的颜色，依次是 RGBA
			rgba.Set(x, y, color.RGBA{190, 190, 190, 255})
		}
	}

	for y := 250; y < 280; y++ {
		for x := 10; x < 620; x++ {
			// 设置某个点的颜色，依次是 RGBA
			rgba.Set(x, y, color.RGBA{28, 172, 255, 255})
		}
	}

	for y := 280; y < 310; y++ {
		for x := 10; x < 620; x++ {
			// 设置某个点的颜色，依次是 RGBA
			rgba.Set(x, y, color.RGBA{190, 190, 190, 255})
		}
	}

	// Draw the text.
	pt := freetype.Pt(10, 100+int(c.PointToFixed(*size)>>6))
	for _, s := range text {
		_, err = c.DrawString(s, pt)
		if err != nil {
			log.Println(err)
			return
		}
		pt.Y += c.PointToFixed(*size * *spacing)
	}
	pt1 := freetype.Pt(200, 100+int(c.PointToFixed(*size)>>6))
	for _, s := range text1 {
		_, err = c.DrawString(s, pt1)
		if err != nil {
			log.Println(err)
			return
		}
		pt1.Y += c.PointToFixed(*size * *spacing)
	}

	pt2 := freetype.Pt(400, 70+int(c.PointToFixed(*size)>>6))
	for _, s := range text2 {
		_, err = c.DrawString(s, pt2)
		if err != nil {
			log.Println(err)
			return
		}
		pt2.Y += c.PointToFixed(*size * *spacing)
	}

	// Save that RGBA image to disk.
	outFile, err := os.Create("out.png")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, rgba)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

}

func main() {
	//0 今日订单量
	todayorders := "71车"
	//1 今日销售额
	todaysoale := "4万元"
	//2 本月订单量
	monthorders := "1777车"
	//3 本月客单价
	mcp := "5.66万元"
	// 4 本月销售额
	monthsoale := "20000万元"
	// 5 本月完成率
	mcr := "95.35%"
	// 6 季度销售额
	qsoale := "60000万元"
	// 7 季度完成率
	qcr := "75.33%"
	// 8 年度销售额
	ysoale := "180000万元"
	// 9 年度完成率
	ycr := "35.33%"

	Repurchaserate := "28.57%"
	todayregistration := "13家"
	monthregistration := "153家"
	PV := "1261次"
	UV := "100个"

	time_date := "2018-06-28 15:35"

	var text = []string{
		"今日",
		"本月",
		"季度",
		"年度",
		"流量 ",
		"注册量",
		"复购率" + Repurchaserate,
	}

	var text1 = []string{
		todaysoale,
		monthsoale,
		qsoale,
		ysoale,
		"PV" + PV,
		"今日" + todayregistration,
		"今日" + todayorders,
	}
	var text2 = []string{
		time_date,
		"客单价" + mcp,
		"完成率" + mcr,
		"完成率" + qcr,
		"完成率" + ycr,
		"UV" + UV,
		"本月" + monthregistration,
		"本月" + monthorders,
	}

	Drawing(text, text1, text2)
}

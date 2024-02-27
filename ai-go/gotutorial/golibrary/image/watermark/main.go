/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 12:04:41
*/
package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"math/rand"
	"os"
	"time"

	"github.com/golang/freetype"
)

// 水印的位置
const (
	TopLeft int = iota
	TopRight
	BottomLeft
	BottomRight
	Center
)

// 字体路径
var Ttf string

type Water struct {
	Pattern string //增加按时间划分的子目录：默认没有时间划分的子目录
}

func (w *Water) New(SavePath, fileName string, typeface []FontInfo) error {
	subPath := w.Pattern
	dirs, err := createDir(SavePath, subPath)
	if err != nil {
		return err
	}
	imgfile, _ := os.Open(fileName)
	defer imgfile.Close()
	_, str, err := image.DecodeConfig(imgfile)
	if err != nil {
		return err
	}
	newName := fmt.Sprintf("%s%s.%s", dirs, getRandomString(10), str)
	if str == "gif" {
		err = gifFontWater(fileName, newName, typeface)
	} else {
		err = staticFontWater(fileName, newName, str, typeface)
	}
	return err
}

// gif图片水印
func gifFontWater(file, name string, typeface []FontInfo) (err error) {
	imgfile, _ := os.Open(file)
	defer imgfile.Close()
	var err2 error
	gifimg2, _ := gif.DecodeAll(imgfile)
	gifs := make([]*image.Paletted, 0)
	x0 := 0
	y0 := 0
	yuan := 0
	for k, gifimg := range gifimg2.Image {
		img := image.NewNRGBA(gifimg.Bounds())
		if k == 0 {
			x0 = img.Bounds().Dx()
			y0 = img.Bounds().Dy()
		}
		fmt.Printf("%v, %v\n", img.Bounds().Dx(), img.Bounds().Dy())
		if k == 0 && gifimg2.Image[k+1].Bounds().Dx() > x0 && gifimg2.Image[k+1].Bounds().Dy() > y0 {
			yuan = 1
			break
		}
		if x0 == img.Bounds().Dx() && y0 == img.Bounds().Dy() {
			for y := 0; y < img.Bounds().Dy(); y++ {
				for x := 0; x < img.Bounds().Dx(); x++ {
					img.Set(x, y, gifimg.At(x, y))
				}
			}
			img, err2 = common(img, typeface) //添加文字水印
			if err2 != nil {
				break
			}
			//定义一个新的图片调色板img.Bounds()：使用原图的颜色域，gifimg.Palette：使用原图的调色板
			p1 := image.NewPaletted(gifimg.Bounds(), gifimg.Palette)
			//把绘制过文字的图片添加到新的图片调色板上
			draw.Draw(p1, gifimg.Bounds(), img, image.ZP, draw.Src)
			//把添加过文字的新调色板放入调色板slice
			gifs = append(gifs, p1)
		} else {
			gifs = append(gifs, gifimg)
		}
	}
	if yuan == 1 {
		return errors.New("gif: image block is out of bounds")
	} else {
		if err2 != nil {
			return err2
		}
		//保存到新文件中
		newfile, err := os.Create(name)
		if err != nil {
			return err
		}
		defer newfile.Close()
		g1 := &gif.GIF{
			Image:     gifs,
			Delay:     gifimg2.Delay,
			LoopCount: gifimg2.LoopCount,
		}
		err = gif.EncodeAll(newfile, g1)
		return err
	}
}

// png,jpeg图片水印
func staticFontWater(file, name, status string, typeface []FontInfo) (err error) {
	//需要加水印的图片
	imgfile, _ := os.Open(file)
	defer imgfile.Close()
	var staticImg image.Image
	if status == "png" {
		staticImg, _ = png.Decode(imgfile)
	} else {
		staticImg, _ = jpeg.Decode(imgfile)
	}
	img := image.NewNRGBA(staticImg.Bounds())
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			img.Set(x, y, staticImg.At(x, y))
		}
	}
	img, err = common(img, typeface) //添加文字水印
	if err != nil {
		return err
	}
	//保存到新文件中
	newfile, err := os.Create(name)
	if err != nil {
		return err
	}
	defer newfile.Close()
	if status == "png" {
		err = png.Encode(newfile, img)
	} else {
		err = jpeg.Encode(newfile, img, &jpeg.Options{Quality: 100})
	}
	return err
}

// 添加文字水印函数
func common(img *image.NRGBA, typeface []FontInfo) (*image.NRGBA, error) {
	var err2 error
	//拷贝一个字体文件到运行目录
	fontBytes, err := os.ReadFile(Ttf)
	if err != nil {
		err2 = err
		return nil, err2
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		err2 = err
		return nil, err2
	}
	errNum := 1
Loop:
	for _, t := range typeface {
		info := t.Message
		f := freetype.NewContext()
		f.SetDPI(108)
		f.SetFont(font)
		f.SetFontSize(t.Size)
		f.SetClip(img.Bounds())
		f.SetDst(img)
		f.SetSrc(image.NewUniform(color.RGBA{R: t.R, G: t.G, B: t.B, A: t.A}))
		//第一行的文字
		// pt := freetype.Pt(img.Bounds().Dx()-len(info)*4-20, img.Bounds().Dy()-100)
		first := 0
		two := 0
		switch int(t.Position) {
		case 0:
			first = t.Dx
			two = t.Dy + int(f.PointToFixed(t.Size)>>6)
		case 1:
			first = img.Bounds().Dx() - len(info)*4 - t.Dx
			two = t.Dy + int(f.PointToFixed(t.Size)>>6)
		case 2:
			first = t.Dx
			two = img.Bounds().Dy() - t.Dy
		case 3:
			first = img.Bounds().Dx() - len(info)*4 - t.Dx
			two = img.Bounds().Dy() - t.Dy
		case 4:
			first = (img.Bounds().Dx() - len(info)*4) / 2
			two = (img.Bounds().Dy() - t.Dy) / 2
		default:
			errNum = 0
			break Loop
		}
		// fmt.Printf("%v, %v, %v\n", first, two, info)
		pt := freetype.Pt(first, two)
		_, err = f.DrawString(info, pt)
		if err != nil {
			err2 = err
			break
		}
	}
	if errNum == 0 {
		err2 = errors.New("坐标值不对")
	}
	return img, err2
}

// 定义添加的文字信息
type FontInfo struct {
	Size     float64 //文字大小
	Message  string  //文字内容
	Position int     //文字存放位置
	Dx       int     //文字x轴留白距离
	Dy       int     //文字y轴留白距离
	R        uint8   //文字颜色值RGBA中的R值
	G        uint8   //文字颜色值RGBA中的G值
	B        uint8   //文字颜色值RGBA中的B值
	A        uint8   //文字颜色值RGBA中的A值
}

// 生成图片名字
func getRandomString(lenght int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	bytesLen := len(bytes)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lenght; i++ {
		result = append(result, bytes[r.Intn(bytesLen)])
	}
	return string(result)
}

// 检查并生成存放图片的目录
func createDir(SavePath, subPath string) (string, error) {
	var dirs string
	if subPath == "" {
		dirs = fmt.Sprintf("%s/", SavePath)
	} else {
		dirs = fmt.Sprintf("%s/%s/", SavePath, time.Now().Format(subPath))
	}
	_, err := os.Stat(dirs)
	if err != nil {
		err = os.MkdirAll(dirs, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	return dirs, nil
}

func main() {
	SavePath := "./kef"
	str := FontInfo{18, "努力向上", TopLeft, 20, 20, 255, 255, 0, 255}
	arr := make([]FontInfo, 0)
	arr = append(arr, str)
	str2 := FontInfo{Size: 24, Message: "努力向上，涨工资", Position: TopLeft, Dx: 20, Dy: 40, R: 255, G: 255, B: 0, A: 255}
	arr = append(arr, str2)
	// 加水印图片路径
	filename := "19.gif"
	w := new(Water)
	w.Pattern = "2006/01/02"
	Ttf = "./wrzh.ttf" // 字体路径
	err := w.New(SavePath, filename, arr)
	if err != nil {
		fmt.Println(err)
	}
}

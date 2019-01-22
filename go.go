package captcha

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/golang/freetype"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// Captcha struct
type Captcha struct {
	nrgba *image.NRGBA
}

// CreateBackground 创建背景
func (captcha *Captcha) CreateBackground(bgColor color.Color) {
	if bgColor == nil {
		bgColor = randLightColor()
	}
	captchaImage := captcha.nrgba
	draw.Draw(captchaImage, captchaImage.Bounds(), image.NewUniform(bgColor), image.ZP, draw.Src)
}

// DrawText 创建文字
func (captcha *Captcha) DrawText(fontDir string) error {
	fontSize := 26.0
	backImage := captcha.nrgba
	x := 0
	y := 35
	fontArr := randFontArr(4)

	if len(fontDir) == 0 {
		fontDir = "./font"
	}

	fileInfoSilce, err := ioutil.ReadDir(fontDir)
	if err != nil {
		return err
	}

	f := freetype.NewContext()
	//设置分辨率
	f.SetDPI(72)
	//设置尺寸
	f.SetClip(backImage.Bounds())
	f.SetFontSize(fontSize)
	// 设置字体颜色
	f.SetSrc(image.NewUniform(color.RGBA{255, 0, 0, 255}))
	//设置输出的图片
	f.SetDst(backImage)

	for _, font := range fontArr {
		fileInfo := fileInfoSilce[r.Intn(3)]
		fileName := fileInfo.Name()
		fontPath := filepath.Join(fontDir, fileName)

		fontBytes, err := ioutil.ReadFile(fontPath)
		if err != nil {
			return err
		}

		fontFamily, err := freetype.ParseFont(fontBytes)
		if err != nil {
			return err
		}

		//设置字体
		f.SetFont(fontFamily)

		x += int(fontSize)
		pt := freetype.Pt(x, y)

		_, err = f.DrawString(font, pt)
	}

	return err
}

// Save 保存验证码
func (captcha *Captcha) Save(w io.Writer) error {
	return jpeg.Encode(w, captcha.nrgba, &jpeg.Options{Quality: 100})
}

// New 新建验证码
func New(width int, height int) *Captcha {
	rect := image.Rect(0, 0, width, height)
	captchaImage := image.NewNRGBA(rect)

	return &Captcha{captchaImage}
}

// randLightColor 随机生成浅色
func randLightColor() color.RGBA {
	red := r.Intn(55) + 200
	green := r.Intn(55) + 200
	blue := r.Intn(55) + 200

	return color.RGBA{uint8(red), uint8(green), uint8(blue), uint8(255)}
}

// randFontArr 随机生成文字切片
func randFontArr(fontNum int) []string {
	var result []string
	allFont := []string{"y", "j", "t", "o", "p", "g"}

	for i := 0; i < fontNum; i++ {
		index := r.Intn(len(allFont))
		result = append(result, allFont[index])
	}

	return result
}

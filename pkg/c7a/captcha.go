package c7a

import (
	"crypto/rand"
	_ "embed"
	"image"
	"image/color"
	"log"
	"math"
	"math/big"
	"strconv"

	"github.com/golang/freetype"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
)

const (
	chars           = "2345678abcdefhijkmnpqrstuvwxyzABCDEFHIJKMNPQRSTUVWXYZ"
	operator        = "+-*/"
	defaultLen      = 4
	defaultFontSize = 25
	defaultDpi      = 72
)

type Mode int

const (
	DirectString      Mode = iota // 普通字符串
	SimpleMathFormula             // 10以内简单数学公式
)

// Captcha 图形验证码 使用字体默认ttf格式
// w 图片宽度, h图片高度，CodeLen验证码的个数
// FontSize 字体大小, Dpi 清晰度
// mode 验证模式
type Captcha struct {
	W, H, CodeLen int
	FontSize      float64
	Dpi           int
	mode          Mode
}

// NewCaptcha 实例化验证码
func NewCaptcha(w, h, CodeLen int) *Captcha {
	return &Captcha{W: w, H: h, CodeLen: CodeLen}
}

// OutPut 输出
func (c *Captcha) OutPut() (string, *image.RGBA) {
	img := c.initCanvas()
	return c.doImage(img)
}

// RangeRand 获取区间[-m, n]的随机数
func (c *Captcha) RangeRand(min, max int64) int64 {
	if min > max {
		panic("the min is greater than max!")
	}

	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := rand.Int(rand.Reader, big.NewInt(max+1+i64Min))

		return result.Int64() - i64Min
	}

	result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
	return min + result.Int64()
}

// getRandCode 随机字符串
func (c *Captcha) getRandCode() string {
	if c.CodeLen <= 0 {
		c.CodeLen = defaultLen
	}

	code := ""
	for l := 0; l < c.CodeLen; l++ {
		charsPos := c.RangeRand(0, int64(len(chars)-1))
		code += string(chars[charsPos])
	}

	return code
}

// 获取算术运算公式
func (c *Captcha) getFormulaMixData() (string, []string) {
	num1 := int(c.RangeRand(11, 20))
	num2 := int(c.RangeRand(1, 10))
	opArr := []rune(operator)
	opRand := opArr[c.RangeRand(0, 3)]

	strNum1 := strconv.Itoa(num1)
	strNum2 := strconv.Itoa(num2)

	var ret int
	var opRet string
	switch string(opRand) {
	case "+":
		ret = num1 + num2
		opRet = "+"
	case "-":
		ret = num1 - num2
		opRet = "-"
	case "*":
		ret = num1 * num2
		opRet = "×"
	case "/":
		ret = num1 / num2
		opRet = "÷"
	}

	return strconv.Itoa(ret), []string{strNum1, opRet, strNum2, "=", "?"}
}

// initCanvas 初始化画布
func (c *Captcha) initCanvas() *image.RGBA {
	dest := image.NewRGBA(image.Rect(0, 0, c.W, c.H))

	// 随机色
	r := uint8(255) // uint8(c.RangeRand(50, 250))
	g := uint8(255) // uint8(c.RangeRand(50, 250))
	b := uint8(255) // uint8(c.RangeRand(50, 250))

	// 填充背景色
	for x := 0; x < c.W; x++ {
		for y := 0; y < c.H; y++ {
			dest.Set(x, y, color.RGBA{R: r, G: g, B: b, A: 255}) // 设定alpha图片的透明度
		}
	}

	return dest
}

// doImage 处理图像
func (c *Captcha) doImage(dest *image.RGBA) (string, *image.RGBA) {
	gc := draw2dimg.NewGraphicContext(dest)

	defer gc.Close()
	defer gc.FillStroke()

	c.setFont(gc)
	c.doPoint(gc)
	c.doLine(gc)
	c.doSinLine(gc)

	var codeStr string
	if c.mode == 1 {
		ret, formula := c.getFormulaMixData()
		log.Printf("formula: %s", formula)
		codeStr = ret
		c.doFormula(gc, formula)
	} else {
		codeStr = c.getRandCode()
		c.doCode(gc, codeStr)
	}

	return codeStr, dest
}

// 验证码字符设置到图像上
func (c *Captcha) doCode(gc *draw2dimg.GraphicContext, code string) {
	for l := 0; l < len(code); l++ {
		y := c.RangeRand(int64(c.FontSize)-1, int64(c.H)+6)
		x := c.RangeRand(1, 20)

		// 随机色
		r := uint8(c.RangeRand(0, 200))
		g := uint8(c.RangeRand(0, 200))
		b := uint8(c.RangeRand(0, 200))

		gc.SetFillColor(color.RGBA{R: r, G: g, B: b, A: 255})
		gc.FillStringAt(string(code[l]), float64(x)+c.FontSize*float64(l), float64(int64(c.H)-y)+c.FontSize)
		gc.Stroke()
	}
}

// 验证码字符设置到图像上
func (c *Captcha) doFormula(gc *draw2dimg.GraphicContext, formulaArr []string) {
	for l := 0; l < len(formulaArr); l++ {
		y := c.RangeRand(0, 10)
		x := c.RangeRand(5, 10)

		// 随机色
		r := uint8(c.RangeRand(10, 200))
		g := uint8(c.RangeRand(10, 200))
		b := uint8(c.RangeRand(10, 200))

		gc.SetFillColor(color.RGBA{r, g, b, 255})

		gc.FillStringAt(formulaArr[l], float64(x)+c.FontSize*float64(l), c.FontSize+float64(y))
		gc.Stroke()
	}
}

// 增加干扰线
func (c *Captcha) doLine(gc *draw2dimg.GraphicContext) {
	for n := 0; n < 5; n++ {
		// gc.SetLineWidth(float64(c.RangeRand(1, 2)))
		gc.SetLineWidth(1)

		// 随机背景色
		r := uint8(c.RangeRand(0, 255))
		g := uint8(c.RangeRand(0, 255))
		b := uint8(c.RangeRand(0, 255))

		gc.SetStrokeColor(color.RGBA{R: r, G: g, B: b, A: 255})

		// 初始化位置
		gc.MoveTo(float64(c.RangeRand(0, int64(c.W)+10)), float64(c.RangeRand(0, int64(c.H)+5)))
		gc.LineTo(float64(c.RangeRand(0, int64(c.W)+10)), float64(c.RangeRand(0, int64(c.H)+5)))

		gc.Stroke()
	}
}

// doPoint 增加干扰点
func (c *Captcha) doPoint(gc *draw2dimg.GraphicContext) {
	for n := 0; n < 50; n++ {
		gc.SetLineWidth(float64(c.RangeRand(1, 3)))

		// 随机色
		r := uint8(c.RangeRand(0, 255))
		g := uint8(c.RangeRand(0, 255))
		b := uint8(c.RangeRand(0, 255))

		gc.SetStrokeColor(color.RGBA{R: r, G: g, B: b, A: 255})

		x := c.RangeRand(0, int64(c.W)+10) + 1
		y := c.RangeRand(0, int64(c.H)+5) + 1

		gc.MoveTo(float64(x), float64(y))
		gc.LineTo(float64(x+c.RangeRand(1, 2)), float64(y+c.RangeRand(1, 2)))

		gc.Stroke()
	}
}

// doSinLine 增加正弦干扰线
func (c *Captcha) doSinLine(gc *draw2dimg.GraphicContext) {
	h1 := c.RangeRand(-12, 12)
	h2 := c.RangeRand(-1, 1)
	w2 := c.RangeRand(5, 20)
	h3 := c.RangeRand(5, 10)

	h := float64(c.H)
	w := float64(c.W)

	// 随机色
	r := uint8(c.RangeRand(128, 255))
	g := uint8(c.RangeRand(128, 255))
	b := uint8(c.RangeRand(128, 255))

	gc.SetStrokeColor(color.RGBA{R: r, G: g, B: b, A: 255})
	gc.SetLineWidth(float64(c.RangeRand(2, 4)))

	var i float64
	for i = -w / 2; i < w/2; i = i + 0.1 {
		y := h/float64(h3)*math.Sin(i/float64(w2)) + h/2 + float64(h1)

		gc.LineTo(i+w/2, y)

		if h2 == 0 {
			gc.LineTo(i+w/2, y+float64(h2))
		}
	}

	gc.Stroke()
}

// SetMode 设置模式
func (c *Captcha) SetMode(mode Mode) { c.mode = mode }

// SetFontSize 设置字体大小
func (c *Captcha) SetFontSize(fontSize float64) { c.FontSize = fontSize }

//go:embed DENNEthree-dee.ttf
var FontDENNEthreedee []byte

// setFont 设置相关字体
func (c *Captcha) setFont(gc *draw2dimg.GraphicContext) {
	font, err := freetype.ParseFont(FontDENNEthreedee) // 字体文件
	if err != nil {
		log.Println(err)
		return
	}

	// 设置自定义字体相关信息
	gc.FontCache = draw2d.NewSyncFolderFontCache("")
	fontData := draw2d.FontData{Name: `DENNEthree-dee`, Style: draw2d.FontStyleNormal}
	gc.FontCache.Store(fontData, font)
	gc.SetFontData(fontData)

	// 设置清晰度
	if c.Dpi <= 0 {
		c.Dpi = defaultDpi
	}
	gc.SetDPI(c.Dpi)

	// 设置字体大小
	if c.FontSize <= 0 {
		c.FontSize = defaultFontSize
	}
	gc.SetFontSize(c.FontSize)
}

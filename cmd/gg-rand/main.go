package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/bingoohuang/gg-rand/pkg/unsplash"
	"github.com/bingoohuang/gg/pkg/chinaid"
	"github.com/bingoohuang/gg/pkg/uid"
	"github.com/google/uuid"

	"github.com/Pallinder/go-randomdata"
	"github.com/bingoohuang/gg-rand/pkg/art"
	"github.com/bingoohuang/gg-rand/pkg/str"
	"github.com/bingoohuang/gg/pkg/randx"
	"github.com/bingoohuang/jj"
	"github.com/brianvoe/gofakeit/v6"
	pwe "github.com/kuking/go-pwentropy"
)

func main() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	p := func(name, value string) {
		fmt.Fprintf(w, "%s\t: %s\n", name, value)
	}
	defer w.Flush()

	p("SillyName", randomdata.SillyName())
	p("Email", randomdata.Email())
	p("IP v4", randomdata.IpV4Address())
	p("IP v6", randomdata.IpV6Address())
	p("UserAgent", randomdata.UserAgentString())
	p("Password", pwe.PwGen(pwe.FormatComplex, pwe.Strength96))
	p("Password Easy", pwe.PwGen(pwe.FormatEasy, pwe.Strength96))
	p("Numbers(60)", gofakeit.DigitN(60))
	p("Letters(60)", gofakeit.LetterN(60))
	jj.DefaultRandOptions.Pretty = false
	p("JSON", string(jj.Rand()))
	p("String", str.RandStr(60))

	ksuid, _ := uid.NewRandom()
	p("KSUID", ksuid.String()+" "+printInspect(ksuid))
	nid, _ := uuid.NewUUID()
	p("UUID", nid.String())

	p("姓名", chinaid.Name())
	p("性别", chinaid.Sex())
	p("地址", chinaid.Address())
	p("手机", chinaid.Mobile())
	p("身份证", chinaid.ChinaID())
	p("有效期", chinaid.ValidPeriod())
	p("发证机关", chinaid.IssueOrg())
	p("邮箱", chinaid.Email())
	p("银行卡", chinaid.BankNo())
	p("日期", chinaid.RandDate().Format("2006年01月02日"))

	w.Flush()
	arts(p)
	p("Unsplash", unsplash.Random())
}

func arts(p func(name string, value string)) {
	switch randx.IntN(19) {
	case 0:
		p("Generative art Junas", art.Junas())
	case 1:
		p("Generative art Random Shapes", art.RandomShapes())
	case 2:
		p("Generative art Color Circle2", art.ColorCircle2())
	case 3:
		p("Generative art Circle Grid", art.CircleGrid())
	case 4:
		p("Generative art Circle Composes Circle", art.CircleComposesCircle())
	case 5:
		p("Generative art Pixel Hole", art.PixelHole())
	case 6:
		p("Generative art Dots Wave", art.DotsWave())
	case 7:
		p("Generative art Contour Line", art.ContourLine())
	case 8:
		p("Generative art Dot Line", art.DotLine())
	case 9:
		p("Generative art Ocean Fish", art.OceanFish())
	case 10:
		p("Generative art Circle Loop", art.CircleLoop())
	case 11:
		p("Generative art Circle Noise", art.CircleNoise())
	case 12:
		p("Generative art Perlin Perls", art.PerlinPerls())
	case 13:
		p("Generative art Canva", art.ColorCanva())
	case 14:
		p("Generative art Julia Set", art.JuliaSet())
	case 15:
		p("Generative art Black Hole", art.BlackHole())
	case 16:
		p("Generative art Silk Sky", art.SilkSky())
	case 17:
		p("Generative art Circle Move", art.CircleMove())
	case 18:
		p("Generative art Random Circle", art.RandomCircle())
	}
}

func printInspect(id uid.KSUID) string {
	return fmt.Sprintf(`(Time: %v, Timestamp: %v, Payload: %v) `,
		id.Time(), id.Timestamp(), strings.ToUpper(hex.EncodeToString(id.Payload())))
}

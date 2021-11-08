package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/bingoohuang/gg-rand/pkg/c7a"
	"github.com/bingoohuang/gg-rand/pkg/dir"
	"github.com/bingoohuang/gg/pkg/ss"

	"github.com/bingoohuang/gg-rand/pkg/art"

	"github.com/bingoohuang/gg-rand/pkg/unsplash"
	"github.com/bingoohuang/gg/pkg/chinaid"
	"github.com/bingoohuang/gg/pkg/uid"
	"github.com/google/uuid"

	"github.com/Pallinder/go-randomdata"
	"github.com/bingoohuang/gg-rand/pkg/str"
	"github.com/bingoohuang/gg/pkg/randx"
	"github.com/bingoohuang/jj"
	"github.com/brianvoe/gofakeit/v6"
	pwe "github.com/kuking/go-pwentropy"
)

func main() {
	tag := ""
	num := 0
	if len(os.Args) >= 2 {
		tag = os.Args[1]
	}
	if len(os.Args) >= 3 {
		num = ss.ParseInt(os.Args[2])
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	okFn := func(string) bool { return true }
	if tag != "" {
		tag = strings.ToUpper(tag)
		okFn = func(name string) bool {
			return strings.Contains(strings.ToUpper(name), tag)
		}
	}
	if num <= 0 {
		num = 1
	}

	p := func(name string, f func() string) {
		if !okFn(name) {
			return
		}

		for i := 0; i < num; i++ {
			_, _ = fmt.Fprintf(w, "%s\t: %s\n", name, f())
		}
	}

	p("SillyName", randomdata.SillyName)
	p("Email", randomdata.Email)
	p("IP v4", randomdata.IpV4Address)
	p("IP v6", randomdata.IpV6Address)
	p("UserAgent", randomdata.UserAgentString)
	p("Password", func() string { return pwe.PwGen(pwe.FormatComplex, pwe.Strength96) })
	p("Password Easy", func() string { return pwe.PwGen(pwe.FormatEasy, pwe.Strength96) })
	p("Numbers", func() string { return gofakeit.DigitN(60) })
	p("Letters", func() string { return gofakeit.LetterN(60) })
	jj.DefaultRandOptions.Pretty = false
	p("JSON", func() string { return string(jj.Rand()) })
	p("String", func() string { return str.RandStr(60) })

	p("Captcha", func() string {
		cp := c7a.NewCaptcha(150, 40, 5)
		code, img := cp.OutPut()
		return code + " " + dir.ToPng(img, false)
	})

	p("KSUID", func() string { ksuid, _ := uid.NewRandom(); return ksuid.String() + " " + printInspect(ksuid) })
	nid, _ := uuid.NewUUID()
	p("UUID", nid.String)
	p("姓名", chinaid.Name)
	p("性别", chinaid.Sex)
	p("地址", chinaid.Address)
	p("手机", chinaid.Mobile)
	p("身份证", chinaid.ChinaID)
	p("有效期", chinaid.ValidPeriod)
	p("发证机关", chinaid.IssueOrg)
	p("邮箱", chinaid.Email)
	p("银行卡", chinaid.BankNo)
	p("日期", func() string { return chinaid.RandDate().Format("2006年01月02日") })

	_ = w.Flush()
	arts(p)
	p("Unsplash", unsplash.Random)
	_ = w.Flush()
}

type artMap struct {
	Name string
	Fn   func() string
}

var artMaps = []artMap{
	{Name: "Junas", Fn: art.Junas},
	{Name: "Random Shapes", Fn: art.RandomShapes},
	{Name: "Color Circle2", Fn: art.ColorCircle2},
	{Name: "Circle Grid", Fn: art.CircleGrid},
	{Name: "Circle Composes Circle", Fn: art.CircleComposesCircle},
	{Name: "Pixel Hole", Fn: art.PixelHole},
	{Name: "Dots Wave", Fn: art.DotsWave},
	{Name: "Contour Line", Fn: art.ContourLine},
	{Name: "Noise Line", Fn: art.NoiseLine},
	{Name: "Dot Line", Fn: art.DotLine},
	{Name: "Ocean Fish", Fn: art.OceanFish},
	{Name: "Circle Loop", Fn: art.CircleLoop},
	{Name: "Domain Warp", Fn: art.DomainWarp},
	{Name: "Circle Noise", Fn: art.CircleNoise},
	{Name: "Perlin Perls", Fn: art.PerlinPerls},
	{Name: "Color Canve", Fn: art.ColorCanve},
	{Name: "Julia Set", Fn: art.JuliaSet},
	{Name: "Black Hole", Fn: art.BlackHole},
	{Name: "Silk Sky", Fn: art.SilkSky},
	{Name: "Circle Move", Fn: art.CircleMove},
	{Name: "Random Circle", Fn: art.RandomCircle},
}

func arts(p1 func(name string, f func() string)) {
	i := randx.IntN(len(artMaps))
	m := artMaps[i]
	p1("Generative art "+m.Name, m.Fn)
}

func printInspect(id uid.KSUID) string {
	return fmt.Sprintf(`(Time: %v, Timestamp: %v, Payload: %v) `,
		id.Time(), id.Timestamp(), strings.ToUpper(hex.EncodeToString(id.Payload())))
}

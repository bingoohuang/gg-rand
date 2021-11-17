package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/bingoohuang/gg/pkg/rotate"

	"github.com/bingoohuang/gg-rand/pkg/img"
	"github.com/bingoohuang/gg/pkg/fla9"

	"github.com/bingoohuang/gg-rand/pkg/art"
	"github.com/bingoohuang/gg-rand/pkg/c7a"
	"github.com/bingoohuang/gg-rand/pkg/unsplash"
	"github.com/bingoohuang/gg/pkg/chinaid"
	"github.com/bingoohuang/gg/pkg/uid"
	"github.com/google/uuid"

	"github.com/Pallinder/go-randomdata"
	"github.com/bingoohuang/gg-rand/pkg/str"
	"github.com/bingoohuang/jj"
	"github.com/brianvoe/gofakeit/v6"
	pwe "github.com/kuking/go-pwentropy"
)

func main() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	p, isHelpTag := createPrinter(w)

	rand.Seed(time.Now().UnixNano())
	p("Base64", func(int) string {
		token := make([]byte, argLen)
		rand.Read(token)
		return base64.StdEncoding.EncodeToString(token)
	})
	p("SillyName", wrap(randomdata.SillyName))
	p("Email", wrap(randomdata.Email))
	p("IP v4", wrap(randomdata.IpV4Address))
	p("IP v6", wrap(randomdata.IpV6Address))
	p("UserAgent", wrap(randomdata.UserAgentString))
	p("Password", func(int) string { return pwe.PwGen(pwe.FormatComplex, pwe.Strength96) })
	p("Password Easy", func(int) string { return pwe.PwGen(pwe.FormatEasy, pwe.Strength96) })
	p("Numbers", func(int) string { return gofakeit.DigitN(60) })
	p("Letters", func(int) string { return gofakeit.LetterN(60) })
	jj.DefaultRandOptions.Pretty = false
	p("JSON", func(int) string { return string(jj.Rand()) })
	p("String", func(int) string { return str.RandStr(60) })

	p("Captcha", func(int) string {
		cp := c7a.NewCaptcha(150, 40, 5)
		code, pImg := cp.OutPut()
		return code + " " + img.ToPng(pImg, false)
	})

	p("KSUID", func(int) string { ksuid, _ := uid.NewRandom(); return ksuid.String() + " " + printInspect(ksuid) })
	p("UUID", func(int) string { return uuid.New().String() })
	p("姓名", wrap(chinaid.Name))
	p("性别", wrap(chinaid.Sex))
	p("地址", wrap(chinaid.Address))
	p("手机", wrap(chinaid.Mobile))
	p("身份证", wrap(chinaid.ChinaID))
	p("有效期", wrap(chinaid.ValidPeriod))
	p("发证机关", wrap(chinaid.IssueOrg))
	p("邮箱", wrap(chinaid.Email))
	p("银行卡", wrap(chinaid.BankNo))
	p("日期", func(int) string { return chinaid.RandDate().Format("2006年01月02日") })

	_ = w.Flush()
	arts(p, w)
	p("Unsplash", unsplash.Random)
	p("Image", img.RandomImage)

	_ = w.Flush()
	if isHelpTag {
		fmt.Println()
	}
}

func wrap(f func() string) func(int) string { return func(int) string { return f() } }

func createPrinter(w *tabwriter.Writer) (func(name string, f func(int) string), bool) {
	tag = strings.ToUpper(tag)

	if tag == "HELP" {
		fmt.Print("Available tags:")
		firstTag := true
		return func(name string, f func(int) string) {
			if firstTag {
				firstTag = false
				fmt.Print(" " + name)
			} else {
				fmt.Print(", " + name)
			}
		}, true
	}

	okFn := func(string) bool { return true }
	if tag != "" {
		okFn = func(name string) bool { return strings.Contains(strings.ToUpper(name), tag) }
	}

	return func(name string, f func(int) string) {
		if okFn(name) {
			for i := 0; i < num; i++ {
				_, _ = fmt.Fprintf(w, "%s\t: %s\n", name, f(i))
			}
		}
	}, false
}

func arts(p1 func(name string, f func(int) string), flusher rotate.Flusher) {
	p1("Generative art", func(i int) string {
		item := artMaps[i%(len(artMaps))]
		result := item.Fn()
		flusher.Flush()
		return item.Name + ": " + result
	})
}

var (
	tag    string
	num    int
	argLen int
	dir    string
)

const usage = `
  -dir string   in which dir to generate random files. (default temp dir)
  -n   int      how many random values to generate. (default 1)
  -len int      length.
  -tag string   which type to generate, like uuid, art, id, email and etc. (empty for all, help to print all available full tags)
`

func init() {
	fla9.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:%s", os.Args[0], usage)
	}
	fla9.StringVar(&dir, "dir", "", "")
	fla9.StringVar(&tag, "tag", "", "")
	fla9.IntVar(&num, "n", 1, "")
	fla9.IntVar(&argLen, "len", 100, "")
	fla9.Parse()

	if dir != "" {
		img.Dir = dir
	}
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

func printInspect(id uid.KSUID) string {
	return fmt.Sprintf(`(Time: %v, Timestamp: %v, Payload: %v) `,
		id.Time(), id.Timestamp(), strings.ToUpper(hex.EncodeToString(id.Payload())))
}

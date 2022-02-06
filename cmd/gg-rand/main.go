package main

import (
	crand "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/aidarkhanov/nanoid/v2"
	"github.com/spaolacci/murmur3"

	"github.com/bingoohuang/gg-rand/pkg/gist"
	"github.com/bingoohuang/gg-rand/pkg/hash"
	"github.com/bwmarrin/snowflake"
	oid "github.com/coolbed/mgo-oid"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/rs/xid"

	rd "github.com/Pallinder/go-randomdata"
	"github.com/bingoohuang/gg-rand/pkg/art"
	"github.com/bingoohuang/gg-rand/pkg/c7a"
	"github.com/bingoohuang/gg-rand/pkg/img"
	"github.com/bingoohuang/gg-rand/pkg/str"
	"github.com/bingoohuang/gg-rand/pkg/unsplash"
	"github.com/bingoohuang/gg/pkg/chinaid"
	"github.com/bingoohuang/gg/pkg/fla9"
	"github.com/bingoohuang/jj"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/chilts/sid"
	"github.com/google/uuid"
	"github.com/jxskiss/base62"
	"github.com/kjk/betterguid"
	pwe "github.com/kuking/go-pwentropy"
	"github.com/lithammer/shortuuid"
	"github.com/oklog/ulid"
	guuid "github.com/satori/go.uuid"
	"github.com/segmentio/ksuid"
	"github.com/sony/sonyflake"
)

func main() {
	p, isHelpTag := createPrinter()

	p("blake3hash-zeebo", createFileFunc(hash.Blake3Zeebo))
	p("blake3hash-luke", createFileFunc(hash.Blake3Luke))
	p("xxhash", createFileFunc(hash.XXH64File))
	p("md5-hash", createFileFunc(hash.MD5HashFile))
	p("sha256-hash", createFileFunc(func(f string) ([]byte, error) { return hash.HashFile(f, sha256.New()) }))
	p("murmur3-32-hash", createFileFunc(func(f string) ([]byte, error) { return hash.HashFile(f, murmur3.New32()) }))
	p("murmur3-64-hash", createFileFunc(func(f string) ([]byte, error) { return hash.HashFile(f, murmur3.New64()) }))
	p("murmur3-128-hash", createFileFunc(func(f string) ([]byte, error) { return hash.HashFile(f, murmur3.New128()) }))
	p("imo-hash", createFileFunc(hash.IMOHashFile))

	rand.Seed(time.Now().UnixNano())
	p("Base64Std", func(int) string { return base64.StdEncoding.EncodeToString(randToken()) })
	p("Base64RawStd", func(int) string { return base64.RawStdEncoding.EncodeToString(randToken()) })
	p("Base64RawURL", func(int) string { return base64.RawURLEncoding.EncodeToString(randToken()) })
	p("Base64URL", func(int) string { return base64.URLEncoding.EncodeToString(randToken()) })
	p("SillyName", wrap(rd.SillyName))
	p("Email", wrap(rd.Email))
	p("IP v4", wrap(rd.IpV4Address))
	p("IP v6", wrap(rd.IpV6Address))
	p("UserAgent", wrap(rd.UserAgentString))
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

	p("sony/sonyflake", func(int) string {
		flake := sonyflake.NewSonyflake(sonyflake.Settings{})
		v, err := flake.NextID()
		if err != nil {
			log.Fatalf("flake.NextID() failed with %s\n", err)
		}
		// Note: this is base16, could shorten by encoding as base62 string
		return base62.EncodeToString(base62.FormatUint(v))
	})
	p("oklog/ulid", func(int) string {
		t := time.Now().UTC()
		entropy := rand.New(rand.NewSource(t.UnixNano()))
		v := ulid.MustNew(ulid.Timestamp(t), entropy)
		return v.String() + "(48位时间(ms)+64位随机)"
	})
	p("chilts/sid", func(int) string { return sid.IdBase64() + " (32位时间(ns)+64位随机)" })
	p("kjk/betterguid", func(int) string { return betterguid.New() + " (32位时间(ms)+72位随机 )" })
	p("segmentio/ksuid", func(int) string {
		v := ksuid.New()
		return v.String() + " (32位时间(s)+128位随机，20字节，base62固定27位，优选)"
	})
	p("lithammer/shortuuid", func(int) string { return shortuuid.New() + " (UUIDv4 or v5, 紧凑编码)" })
	p("google/uuid v4", func(int) string { return uuid.New().String() + " (128位随机)" })
	p("satori/go.uuid v4", func(int) string { return guuid.NewV4().String() + " (UUIDv4 from RFC 4112 for comparison)" })
	p("aidarkhanov/nanoid/v2", func(int) string { return PickStr(nanoid.New()) })  // "i25_rX9zwDdDn7Sg-ZoaH"
	p("matoous/go-nanoid/v2", func(int) string { return PickStr(gonanoid.New()) }) // "i25_rX9zwDdDn7Sg-ZoaH"
	p("coolbed/mgo-oid Mongodb Object ID", func(int) string {
		v := oid.NewOID()
		return fmt.Sprintf("%s (Timestamp: %d)", v, v.Timestamp())
	})
	p("rs/xid Mongo Object ID", func(int) string {
		v := xid.New()
		m := base64.RawURLEncoding.EncodeToString(v.Machine())
		return fmt.Sprintf("%s (Machine: %s, Pid: %d, Time: %s, Counter: %d) (4B time(s) + 3B machine id + 2B pid + 3Brandom)", v, m, v.Pid(), v.Time(), v.Counter())
	})
	p("BSON Object ID", func(int) string { return gist.NewObjectId().String() })
	p("Snowfake ID", func(int) string {
		n, _ := snowflake.NewNode(1)
		v := n.Generate()
		return fmt.Sprintf("%d (Time: %d, Node: %d, Step:%d) (~6B time(10ms) + 1B # + 2B machine id)", v, v.Time(), v.Node(), v.Step())
	})
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

	arts(p)
	p("Unsplash", unsplash.Random)
	p("Image", img.RandomImage)

	if isHelpTag {
		fmt.Println()
	}
}

func randToken() []byte {
	token := make([]byte, argLen)
	rand.Read(token)
	return token
}

func PickStr(s string, _ interface{}) string {
	return s
}

func wrap(f func() string) func(int) string { return func(int) string { return f() } }

func createPrinter() (func(name string, f func(int) string), bool) {
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
			start0 := time.Now()
			for i := 0; i < num; i++ {
				start := time.Now()
				v := f(i)
				log.Printf("%s: %s (len: %d), cost %s", name, v, len(v), time.Since(start))
			}
			if cost && num > 1 {
				log.Printf("Completed, cost %s", time.Since(start0))
			}
		}
	}, false
}

func arts(p1 func(name string, f func(int) string)) {
	p1("Generative art", func(i int) string {
		item := artMaps[i%(len(artMaps))]
		result := item.Fn()
		return item.Name + ": " + result
	})
}

var (
	tag    string
	input  string
	num    int
	argLen int
	dir    string
	cost   bool
)

const usage = `
  -dir,d   string   In which dir to generate random files. (default temp dir)
  -input,i string   Input file name. 
  -n       int      How many random values to generate. (default 1)
  -len,l   int      Length.
  -cost,c  bool     Time costed.
  -tag     string   Which type to generate, like uuid, art, id, email and etc. (empty for all, help to print all available full tags)
`

func init() {
	fla9.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:%s", os.Args[0], usage)
	}
	fla9.StringVar(&dir, "dir,d", "", "")
	fla9.BoolVar(&cost, "cost,c", false, "")
	fla9.StringVar(&tag, "tag,t", "", "")
	fla9.StringVar(&input, "input,i", "", "")
	fla9.IntVar(&num, "n", 1, "")
	fla9.IntVar(&argLen, "len,l", 100, "")
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

func printInspect(id ksuid.KSUID) string {
	return fmt.Sprintf(`(Time: %v, Timestamp: %v, Payload: %v) `,
		id.Time(), id.Timestamp(), strings.ToUpper(hex.EncodeToString(id.Payload())))
}

func createFileFunc(f func(string) ([]byte, error)) func(int) string {
	return func(int) string {
		file := input
		stat, err := os.Stat(file)
		if err != nil {
			if !os.IsNotExist(err) {
				log.Printf("failed to stat %s error: %v", file, err)
				return ""
			}

			if temp, err := os.CreateTemp("", ""); err != nil {
				log.Printf("failed to create temporary file: %v", err)
				return ""
			} else {
				_, _ = io.CopyN(temp, crand.Reader, 10*1024*1024)
				temp.Close()
				file = temp.Name()
			}
		} else if stat.IsDir() {
			log.Printf("%s is not allowed to be a directory", file)
			return ""
		}

		d, err := f(file)
		if err != nil {
			log.Printf("error creating hash: %v", err)
		}
		return base64.RawURLEncoding.EncodeToString(d)
	}
}

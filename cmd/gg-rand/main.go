package main

import (
	crand "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/bingoohuang/gg-rand/pkg/genpw"
	"github.com/oklog/ulid"

	"github.com/bingoohuang/gg/pkg/ss"
	"github.com/bingoohuang/gou/pbe"
	"github.com/manifoldco/promptui"

	"github.com/bingoohuang/gg-rand/pkg/ksid"

	"github.com/bingoohuang/gg-rand/pkg/snow2"

	"github.com/bingoohuang/gg-rand/pkg/cid"
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
	"github.com/chilts/sid"
	"github.com/google/uuid"
	"github.com/jxskiss/base62"
	"github.com/kjk/betterguid"
	pwe "github.com/kuking/go-pwentropy"
	"github.com/lithammer/shortuuid"
	guuid "github.com/satori/go.uuid"
	"github.com/segmentio/ksuid"
	"github.com/sony/sonyflake"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	p := createPrinter()

	for p != nil {
		defineRandoms(p)
		p = prompt()
	}
}

func prompt() func(name string, f func(int) interface{}) {
	if tag != "" {
		return nil
	}

	prompt := promptui.Select{
		Label: "Select One of the Randoms",
		Items: allTags,
		Searcher: func(input string, index int) bool {
			return ss.ContainsFold(allTags[index], input)
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil
	}

	return printRandom(func(name string) bool { return name == result })
}

func defineRandoms(p func(name string, f func(int) interface{})) {
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
	p("Base64Std", func(int) interface{} { return base64.StdEncoding.EncodeToString(randToken()) })
	p("Base64RawStd", func(int) interface{} { return base64.RawStdEncoding.EncodeToString(randToken()) })
	p("Base64RawURL", func(int) interface{} { return base64.RawURLEncoding.EncodeToString(randToken()) })
	p("Base64URL", func(int) interface{} { return base64.URLEncoding.EncodeToString(randToken()) })
	p("SillyName", wrap(rd.SillyName))
	p("Email", wrap(rd.Email))
	p("IP v4", wrap(rd.IpV4Address))
	p("IP v6", wrap(rd.IpV6Address))
	p("UserAgent", wrap(rd.UserAgentString))
	p("Pwd", func(int) interface{} {
		return string(genpw.Gen(genpw.WithMinCountOfNumbers(1), genpw.WithMinCountOfSymbols(1)))
	})
	p("Password", func(int) interface{} { return pwe.PwGen(pwe.FormatComplex, pwe.Strength96) })
	p("Password Easy", func(int) interface{} { return pwe.PwGen(pwe.FormatEasy, pwe.Strength96) })
	p("Numbers", func(int) interface{} { return gofakeit.DigitN(60) })
	p("Letters", func(int) interface{} { return gofakeit.LetterN(60) })
	jj.DefaultRandOptions.Pretty = false
	p("JSON", func(int) interface{} { return string(jj.Rand()) })
	p("String", func(int) interface{} { return str.RandStr(60) })

	p("Captcha DirectString", func(int) interface{} {
		cp := c7a.NewCaptcha(150, 40, 5)
		cp.SetMode(c7a.DirectString)
		code, pImg := cp.OutPut()
		return code + " " + img.ToPng(pImg, false)
	})

	p("Captcha SimpleMathFormula", func(int) interface{} {
		cp := c7a.NewCaptcha(150, 40, 5)
		cp.SetMode(c7a.SimpleMathFormula)
		code, pImg := cp.OutPut()
		return code + " " + img.ToPng(pImg, false)
	})

	p("sony/sonyflake", func(int) interface{} {
		flake := sonyflake.NewSonyflake(sonyflake.Settings{})
		v, err := flake.NextID()
		if err != nil {
			log.Fatalf("flake.NextID() failed with %s\n", err)
		}
		// Note: this is base16, could shorten by encoding as base62 string
		return base62.EncodeToString(base62.FormatUint(v))
	})
	p("max", func(int) interface{} {
		return []string{
			fmt.Sprintf("\n int64: %d (len: %d), uint64: %d (len: %d)",
				math.MaxInt64, len(fmt.Sprintf("%+v", math.MaxInt64)),
				uint64(math.MaxUint64), len(fmt.Sprintf("%+v", uint64(math.MaxUint64)))),

			fmt.Sprintf("\n int32: %d (len: %d), uint32: %d (len: %d)",
				math.MaxInt32, len(fmt.Sprintf("%+v", math.MaxInt32)),
				math.MaxUint32, len(fmt.Sprintf("%+v", math.MaxUint32))),

			fmt.Sprintf("\n int16: %d (len: %d), uint16: %d (len: %d)",
				math.MaxInt16, len(fmt.Sprintf("%+v", math.MaxInt16)),
				math.MaxUint16, len(fmt.Sprintf("%+v", math.MaxUint16))),

			fmt.Sprintf("\n int8: %d (len: %d), uint8: %d (len: %d)",
				math.MaxInt8, len(fmt.Sprintf("%+v", math.MaxInt8)),
				math.MaxUint8, len(fmt.Sprintf("%+v", math.MaxUint8))),

			fmt.Sprintf("\n float64: %f, float32 %f", math.MaxFloat64, math.MaxFloat32),
		}
	})
	p("oklog/ulid", func(int) interface{} {
		v := ulid.MustNew(ulid.Now(), crand.Reader)
		return []string{v.String(), "base32固定26位，48位时间(ms)+80位随机"}
	})
	p("chilts/sid", func(int) interface{} { return []string{sid.IdBase64(), "32位时间(ns)+64位随机"} })
	p("kjk/betterguid", func(int) interface{} { return []string{betterguid.New(), "32位时间(ms)+72位随机"} })
	p("segmentio/ksuid", func(int) interface{} {
		return []string{ksuid.New().String(), "32位时间(s)+128位随机，20字节，base62固定27位，优选"}
	})

	p("customized/ksuid base64", func(int) interface{} {
		t := time.Now()
		a := ksid.New(ksid.WithValue(ksid.Nil), ksid.WithTime(t)).String()
		b := ksid.New(ksid.WithTime(t)).String()
		c := ksid.New(ksid.WithValue(ksid.Max), ksid.WithTime(t)).String()
		return []string{a, b, c, fmt.Sprintf("a<=b: %v", a <= b), fmt.Sprintf("b<=c: %v", b <= c)}
	})
	p("lithammer/shortuuid", func(int) interface{} { return []string{shortuuid.New(), "UUIDv4 or v5, 紧凑编码"} })
	p("google/uuid v4", func(int) interface{} { return []string{uuid.New().String(), "128位随机"} })
	p("satori/go.uuid v4", func(int) interface{} {
		return []string{guuid.NewV4().String(), "UUIDv4 from RFC 4112 for comparison"}
	})
	p("aidarkhanov/nanoid/v2", func(int) interface{} { return PickStr(nanoid.New()) })  // "i25_rX9zwDdDn7Sg-ZoaH"
	p("matoous/go-nanoid/v2", func(int) interface{} { return PickStr(gonanoid.New()) }) // "i25_rX9zwDdDn7Sg-ZoaH"
	p("coolbed/mgo-oid Mongodb Object ID", func(int) interface{} {
		v := oid.NewOID()
		return []string{v.String(), fmt.Sprintf("Timestamp: %d", v.Timestamp())}
	})
	p("rs/xid Mongo Object ID", func(int) interface{} {
		v := xid.New()
		m := base64.RawURLEncoding.EncodeToString(v.Machine())
		return []string{v.String(), fmt.Sprintf(
			"32 位Time: %s, 24位Machine: %s, Pid: %d, , Counter: %d 4B time(s) + 3B machine id + 2B pid + 3Brandom",
			v.Time(), m, v.Pid(), v.Counter())}
	})
	p("BSON Object ID", func(int) interface{} { return gist.NewObjectId().String() })
	n, _ := snowflake.NewNode(1)
	p("snowflake ID", func(int) interface{} {
		v := n.Generate()
		return []string{v.String(), fmt.Sprintf("41位 Time: %d, 10位 Node: %d, 12位 Step:%d", v.Time(), v.Node(), v.Step())}
	})

	p("Random ID with fixed length 12", func(int) interface{} { return fmt.Sprintf("%d", cid.Cid12()) })
	p("customized snowflake ID with fixed length 12", func(int) interface{} {
		return fmt.Sprintf("%d", snow2.Node12.Next())
	})
	p("customized snowflake ID with uint32", func(int) interface{} {
		return fmt.Sprintf("%d", snow2.NodeUint32.Next())
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
	p("日期", func(int) interface{} { return chinaid.RandDate().Format("2006年01月02日") })

	arts(p)
	p("Unsplash", unsplash.Random)
	p("Image", img.RandomImage)

	p("PBE Encrypt", pbeEncrptDealer)
	p("PBE Decrypt", pbeDecrptDealer)
}

func pbeDecrptDealer(int) interface{} {
	if tag == "ALL" {
		return nil
	}

	validate := func(input string) error {
		if len(input) < 6 {
			return errors.New("password must have more than 6 characters")
		}
		return nil
	}

	plain := promptui.Prompt{
		Label:    "PBE format",
		Validate: validate,
		Default:  pwe.PwGen(pwe.FormatComplex, pwe.Strength96),
	}

	plainResult, err := plain.Run()
	if err != nil {
		return err.Error()
	}

	passwd := promptui.Prompt{
		Label:    "Password",
		Validate: validate,
		Mask:     '*',
	}
	passwdResult, err := passwd.Run()
	if err != nil {
		return err.Error()
	}

	result, err := pbe.Decrypt(plainResult, passwdResult, 19)
	if err != nil {
		return err.Error()
	}

	return []string{result, plainResult}
}

func pbeEncrptDealer(int) interface{} {
	if tag == "ALL" {
		return nil
	}

	validate := func(input string) error {
		if len(input) < 6 {
			return errors.New("password must have more than 6 characters")
		}
		return nil
	}

	plain := promptui.Prompt{
		Label:    "Plain",
		Validate: validate,
		Default:  pwe.PwGen(pwe.FormatComplex, pwe.Strength96),
	}

	plainResult, err := plain.Run()
	if err != nil {
		return err.Error()
	}

	passwd := promptui.Prompt{
		Label:    "Password",
		Validate: validate,
		Mask:     '*',
	}
	passwdResult, err := passwd.Run()
	if err != nil {
		return err.Error()
	}

	result, err := pbe.Encrypt(plainResult, passwdResult, 19)
	if err != nil {
		return err.Error()
	}

	return []string{result, plainResult}
}

func randToken() []byte {
	token := make([]byte, argLen)
	n, err := crand.Read(token)
	if err != nil {
		panic(err)
	}

	return token[:n]
}

func PickStr(s string, _ interface{}) string {
	return s
}

func wrap(f func() string) func(int) interface{} { return func(int) interface{} { return f() } }

var allTags []string

func createPrinter() func(name string, f func(int) interface{}) {
	tag = strings.ToUpper(tag)
	if tag == "" || tag == "HELP" {
		return func(name string, f func(int) interface{}) {
			allTags = append(allTags, name)
		}
	}

	okFn := func(string) bool { return true }

	if tag == "ALL" {
		okFn = func(string) bool { return true }
	} else if tag != "" {
		okFn = func(name string) bool { return strings.Contains(strings.ToUpper(name), tag) }
	}

	return printRandom(okFn)
}

func printRandom(okFn func(string) bool) func(name string, f func(int) interface{}) {
	return func(name string, f func(int) interface{}) {
		if okFn(name) {
			start0 := time.Now()
			for i := 0; i < num; i++ {
				start := time.Now()
				v := f(i)
				if !extra {
					fmt.Printf("%v\n", v)
					continue
				}

				if v2, ok := v.([]string); ok && len(v2) > 0 {
					log.Printf("%s: %s (len: %d) %s, cost %s", name, v2[0], len(v2[0]), v2[1:], time.Since(start))
				} else if v1, ok := v.(string); ok {
					log.Printf("%s: %s (len: %d), cost %s", name, v1, len(v1), time.Since(start))
				} else {
					log.Printf("%s: %v , cost %s", name, v, time.Since(start))
				}
			}
			if cost && num > 1 && extra {
				log.Printf("Completed, cost %s", time.Since(start0))
			}
		}
	}
}

func arts(p1 func(name string, f func(int) interface{})) {
	p1("Generative art", func(i int) interface{} {
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
	extra  bool
)

const usage = `
  -dir,d   string   In which dir to generate random files. (default temp dir)
  -input,i string   Input file name. 
  -n       int      How many random values to generate. (default 1)
  -extra   bool     Print extra information.
  -len,l   int      Length.
  -cost,c  bool     Time costed.
  -tag     string   Which type to generate, like uuid, art, id, email and etc. (empty for prompt, all or all)
`

func init() {
	fla9.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:%s", os.Args[0], usage)
	}
	fla9.StringVar(&dir, "dir,d", "", "")
	fla9.BoolVar(&cost, "cost,c", false, "")
	fla9.BoolVar(&extra, "extra", false, "")
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

func createFileFunc(f func(string) ([]byte, error)) func(int) interface{} {
	return func(int) interface{} {
		file := input
		stat, err := os.Stat(file)
		if err != nil {
			if !os.IsNotExist(err) {
				log.Printf("failed to stat %s error: %v", file, err)
				return nil
			}

			if temp, err := os.CreateTemp("", ""); err != nil {
				log.Printf("failed to create temporary file: %v", err)
				return nil
			} else {
				_, _ = io.CopyN(temp, crand.Reader, 10*1024*1024)
				temp.Close()
				file = temp.Name()
			}
		} else if stat.IsDir() {
			log.Printf("%s is not allowed to be a directory", file)
			return nil
		}

		d, err := f(file)
		if err != nil {
			log.Printf("error creating hash: %v", err)
		}
		return base64.RawURLEncoding.EncodeToString(d)
	}
}

package genpw

import (
	"bytes"
	"crypto/rand"
	"math/big"
)

var (
	numbers = []byte("0123456789")
	symbols = []byte("!\"#$%&'()=~|-^\\`{*}<>?_@[;:],./'")
)

func count(bs []byte, bp []byte) int {
	n := 0
	for _, b := range bs {
		if bytes.Index(bp, []byte{b}) != -1 {
			n++
		}
	}
	return n
}

func Gen(fns ...OptionFn) []byte {
	option := (OptionFns(fns)).createOption()

	nnc := 0
	if option.MinCountOfNums > 0 {
		nnc = option.MinCountOfNums
	}
	nsc := 0
	if option.MinCountOfSymbols > 0 {
		nsc = option.MinCountOfSymbols
	}

	if option.CountOfChars < nnc+nsc {
		option.CountOfChars = nnc + nsc
	}

	var buf bytes.Buffer
	if option.CountOfChars > nnc+nsc {
		for r := 'a'; r < 'z'; r++ {
			buf.WriteRune(r)
		}
		for r := 'A'; r < 'Z'; r++ {
			buf.WriteRune(r)
		}
	}
	if option.MinCountOfNums != 0 && nsc != option.CountOfChars {
		buf.Write(numbers)
	}
	if option.MinCountOfSymbols != 0 && nnc != option.CountOfChars {
		buf.Write(symbols)
	}
	chars := buf.Bytes()

	var pw bytes.Buffer
	for {
		pw.Reset()
		for i := 0; i < option.CountOfChars; i++ {
			r, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
			if err != nil {
				panic(err)
			}
			pw.WriteByte(chars[int(r.Int64())])
		}

		if option.MinCountOfNums > 0 && count(pw.Bytes(), numbers) < option.MinCountOfNums {
			continue
		}
		if option.MinCountOfSymbols > 0 && count(pw.Bytes(), symbols) < option.MinCountOfSymbols {
			continue
		}
		break
	}

	return pw.Bytes()
}

type Option struct {
	CountOfChars      int
	MinCountOfNums    int
	MinCountOfSymbols int
}

type OptionFn func(*Option)

type OptionFns []OptionFn

func (fns OptionFns) createOption() *Option {
	option := &Option{}

	for _, fn := range fns {
		fn(option)
	}

	if option.CountOfChars <= 0 {
		option.CountOfChars = 16
	}

	return option
}

func WithOption(v Option) OptionFn         { return func(o *Option) { *o = v } }
func WithCountOfCharacters(v int) OptionFn { return func(o *Option) { o.CountOfChars = v } }
func WithMinCountOfNumbers(v int) OptionFn { return func(o *Option) { o.MinCountOfNums = v } }
func WithMinCountOfSymbols(v int) OptionFn { return func(o *Option) { o.MinCountOfSymbols = v } }

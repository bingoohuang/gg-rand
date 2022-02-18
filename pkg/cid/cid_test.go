package cid

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"math/big"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/RoaringBitmap/roaring/roaring64"
)

func TestCid12(t *testing.T) {
	bitmap := roaring64.BitmapOf()

	cnt := 0
	not12cnt := 0
	n := 10000
	for i := 0; i < n; i++ {
		v := Cid12()
		if len(fmt.Sprintf("%d", v)) != 12 {
			not12cnt++
		}

		if bitmap.Contains(v) {
			cnt++
		} else {
			bitmap.Add(v)
		}
	}

	t.Log(cnt, "duplicated,", not12cnt, " not 12, in", n)

	//    cid_test.go:28: 39 duplicated, 0  not 12, in 100000
	//--- PASS: TestCid12 (0.19s)

	//     cid_test.go:28: 3741 duplicated, 0  not 12, in 1000000
	//--- PASS: TestCid12 (1.63s)

	//     cid_test.go:28: 363164 duplicated, 0  not 12, in 10000000
	//--- PASS: TestCid12 (18.72s)
}

func TestFindAtLeastBits(t *testing.T) {
	t.Log(((999999999999 >> 12) - (100000000000 >> 12)) / 60 / 6 / 24 / 365)
	t.Log(FormatBinaryUint64(4294967295))
	// (976562500-97656251)/60/60/24/365 = 27.8年
	// min: 97656251 max: 976562500 diff: 878906249 years: 27
	// min12: 100000000000 BINARY: 0b00010111_01001000_01110110_11101000_00000000
	// max12: 999999999999 BINARY: 0b11101000_11010100_10100101_00001111_11111111
	min12, max12 := MinNum(12), MaxNum(12)

	min, max, diff, years := YearsCanUse(12, 10, 365*24*(time.Hour/time.Second))
	t.Log("min:", min, "max:", max, "diff:", diff, "years:", years)

	t.Log("min12:", min12, "BINARY:", FormatBinaryBig(min12))
	t.Log("max12:", max12, "BINARY:", FormatBinaryBig(max12))
	// t.Log(0b0001011101010) // 746, (2^13-746)/365 = 20.4 能用20.4年，每天可用ID数是 2^27 = 134,217,728，13亿

	//  min20: 10000000000000000000 BINARY: 0b10001010_11000111_00100011_00000100_10001001_11101000_00000000_00000000
	//  max20: 99999999999999999999 BINARY: 0b00000101_01101011_11000111_01011110_00101101_01100011_00001111_11111111_11111111
	min20, max20 := MinNum(20), MaxNum(20)
	t.Log("min20:", min20, "BINARY:", FormatBinaryBig(min20))
	t.Log("max20:", max20, "BINARY:", FormatBinaryBig(max20))
}

func MinNum(fixedLength int) (b *big.Int) {
	sv := strings.Repeat("9", fixedLength-1)
	n, _ := new(big.Int).SetString(sv, 10)
	return BigAdd(n, 1)
}

func YearsCanUse(fixedLength int, rsh uint, unit time.Duration) (min, max, diff, years *big.Int) {
	min, max = MinNum(fixedLength), MaxNum(fixedLength)

	a := BigAdd(BigRsh(min, rsh), 1)
	b := BigAdd(BigRsh(max, rsh), 1)
	diff = new(big.Int).Sub(b, a)
	return a, b, diff, new(big.Int).Div(diff, new(big.Int).SetInt64(int64(unit)))
}

func BigRsh(a *big.Int, n uint) (b *big.Int) {
	return new(big.Int).Rsh(a, n)
}

func BigAdd(a *big.Int, add int64) (b *big.Int) {
	return a.Add(a, new(big.Int).SetInt64(add))
}

func MaxNum(fixedLength int) (b *big.Int) {
	sv := strings.Repeat("9", fixedLength)
	n, _ := new(big.Int).SetString(sv, 10)
	return n
}

func TestFormatBig(t *testing.T) {
	min12, max12 := MinNum(12), MaxNum(12)
	t.Log(fmt.Sprintf("%b", min12))
	t.Log(FormatBinaryBig(min12))
	t.Log(fmt.Sprintf("%b", max12))
	t.Log(FormatBinaryBig(max12))
}

func FormatBinaryBig(b *big.Int) string {
	return FormatBinary(fmt.Sprintf("%b", b))
}

func FormatBinaryUint64(b uint64) string {
	return FormatBinaryBig(new(big.Int).SetUint64(b))
}

func FormatBinary(str string) string {
	strLength := len(str)
	n := int(math.Ceil(float64(strLength) / float64(8)))

	n0 := strLength % 8
	sn := make([]string, n)
	if n0 == 0 {
		n0 = 8
	}
	sn[0] = strings.Repeat("0", 8-n0) + str[:n0]
	for i := 1; i < n; i++ {
		sn[i] = str[n0+(i-1)*8 : n0+i*8]
	}

	return "0b" + strings.Join(sn, "_")
}

func ToBinary(s []byte) string {
	res := ""
	for _, c := range s {
		res = fmt.Sprintf("%s%.8b_", res, c)
	}

	if len(s) > 0 {
		return res[:len(res)-1]
	}
	return res
}

func ParseBinary(s string) (b []byte) {
	for s != "" {
		if i, err := strconv.ParseInt(s[:8], 2, 64); err == nil {
			b = append(b, byte(i))
		}
		s = strings.Trim(s[8:], "_ ")
	}
	return b
}

func TestBinary(t *testing.T) {
	assert.Equal(t, "00000001", ToBinary([]byte{1}))
	assert.Equal(t, "00000001_00000010", ToBinary([]byte{1, 2}))
	assert.Equal(t, []byte{1}, ParseBinary("00000001"))
	assert.Equal(t, []byte{1, 2}, ParseBinary("00000001 00000010"))
}

package cid

import (
	"fmt"
	"github.com/RoaringBitmap/roaring/roaring64"
	"math"
	"math/big"
	"strings"
	"testing"
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
	t.Log(97656251 << 10)
	t.Log(0b101110100100001110110111011)
	t.Log(0b11101000_11010100_10100101_000100)
	// (976562500-1525880)/60/24/365 = 1855年
	// 100000000000: 1011101001000011101101110100000000000
	// 999999999999: 1110100011010100101001010000111111111111
	min12, max12 := MinNum(12), MaxNum(12)

	t.Log(BigAdd(new(big.Int).Rsh(min12, 10), 1))
	t.Log(BigAdd(new(big.Int).Rsh(max12, 10), 1))

	t.Log("min12:", min12, "BINARY:", FormatBig(min12))
	t.Log("max12:", max12, "BINARY:", FormatBig(max12))
	//t.Log(0b0001011101010) // 746, (2^13-746)/365 = 20.4 能用20.4年，每天可用ID数是 2^27 = 134,217,728，13亿

	//  10000000000000000000:          0b10001010_11000111_00100011_00000100_10001001_11101000_00000000_00000000
	//  99999999999999999999: 00000101 01101011 11000111 01011110 00101101 01100011 00001111 11111111 11111111
	min20, max20 := MinNum(20), MaxNum(20)
	t.Log("min20:", min20, "BINARY:", FormatBig(min20))
	t.Log("max20:", max20, "BINARY:", FormatBig(max20))
}

func MinNum(fixedLength int) (b *big.Int) {
	sv := strings.Repeat("9", fixedLength-1)
	n, _ := new(big.Int).SetString(sv, 10)
	return BigAdd(n, 1)
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
	t.Log(FormatBig(min12))
	t.Log(fmt.Sprintf("%b", max12))
	t.Log(FormatBig(max12))
}

func FormatBig(b *big.Int) string {
	str := fmt.Sprintf("%b", b)

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

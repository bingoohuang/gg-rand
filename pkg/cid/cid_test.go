package cid

import (
	"fmt"
	"github.com/RoaringBitmap/roaring/roaring64"
	"math/big"
	"strings"
	"testing"
)

func TestCid12(t *testing.T) {
	bitmap := roaring64.BitmapOf()

	cnt := 0
	not12cnt := 0
	n := 100000
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
	// 100000000000: 00010111 01001000 01110110 11101000 00000000
	// 999999999999: 11101000 11010100 10100101 00001111 11111111
	min12, max12 := MinNum(12), MaxNum(12)
	t.Log("min12:", min12)
	t.Log(FormatBig(min12))
	//t.Log(fmt.Sprintf("%b", min12))
	t.Log("max12:", max12)
	t.Log(FormatBig(max12))
	//t.Log(fmt.Sprintf("%b", max12))
	//t.Log(0b0001011101010) // 746, (2^13-746)/365 = 20.4 能用20.4年，每天可用ID数是 2^27 = 134,217,728，13亿

	//  10000000000000000000:          10001010 11000111 00100011 00000100 10001001 11101000 00000000 00000000
	//  99999999999999999999: 00000101 01101011 11000111 01011110 00101101 01100011 00001111 11111111 11111111
	min20, max20 := MinNum(20), MaxNum(20)
	t.Log("min20:", min20)
	t.Log(FormatBig(min20))
	//t.Log(fmt.Sprintf("%b", min20))
	t.Log("max20:", max20)
	t.Log(FormatBig(max20))
	//t.Log(fmt.Sprintf("%b", max20))
}

func MinNum(fixedLength int) (b *big.Int) {
	sv := strings.Repeat("9", fixedLength-1)
	n, _ := new(big.Int).SetString(sv, 10)
	m := new(big.Int)
	m.SetInt64(1)
	return m.Add(n, m)
}

func MaxNum(fixedLength int) (b *big.Int) {
	sv := strings.Repeat("9", fixedLength)
	n, _ := new(big.Int).SetString(sv, 10)
	return n
}

func FormatBig(b *big.Int) string {
	v := ""
	s := fmt.Sprintf("%b", b)
	for i := len(s) - 8; ; i -= 8 {
		if i < 0 {
			v = strings.Repeat("0", -i) + s[:i+8] + " " + v
			break
		} else {
			v = s[i:i+8] + " " + v
			if i == 0 {
				break
			}
		}
	}

	return v
}

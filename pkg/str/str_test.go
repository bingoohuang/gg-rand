package str

import (
	crand "crypto/rand"
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
	"unsafe"

	"github.com/lazybeaver/xorshift"
	"github.com/vpxyz/xorshift/xoroshiro256plus"
	"github.com/vpxyz/xorshift/xorshift1024star"
)

// Implementations
func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var xs1024 = xorshift1024star.NewSource(time.Now().UnixNano())

func RandStringSource(n int, src rand.Source) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func RandString(n int) string {
	return RandStringSource(n, xs1024)
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func RandStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func RandStringBytesMask(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}

func RandStringBytesMaskImpr(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImpRandSrc(n int) string {
	return RandStringSource(n, src)
}

func RandStringBytesMaskImpXorshift1024Src(n int) string {
	return RandStringSource(n, xs1024)
}

var xs256 = xoroshiro256plus.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImpXorshift256Src(n int) string {
	return RandStringSource(n, xs256)
}

// A Source represents a source of uniformly-distributed
// pseudo-random int64 values in the range [0, 1<<63).
type xorshiftSource struct {
	xorshift.XorShift
}

func (x xorshiftSource) Int63() int64 {
	return int64(x.XorShift.Next())
}

func (x xorshiftSource) Seed(seed int64) {
}

var xor64s = xorshift.NewXorShift64Star(uint64(time.Now().UnixNano()))

func RandStringBytesMaskImpXorShift64StarSrc(n int) string {
	return RandStringSource(n, xorshiftSource{xor64s})
}

var xor128p = xorshift.NewXorShift128Plus(uint64(time.Now().UnixNano()))

func RandStringBytesMaskImpXorShift128PlusSrc(n int) string {
	return RandStringSource(n, xorshiftSource{xor128p})
}

var xor1024s = xorshift.NewXorShift1024Star(uint64(time.Now().UnixNano()))

func RandStringBytesMaskImpXorShift1024StarSrc(n int) string {
	return RandStringSource(n, xorshiftSource{xor1024s})
}

// Benchmark functions
const n = 16

func BenchmarkRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringRunes(n)
	}
}

func BenchmarkBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytes(n)
	}
}

func BenchmarkBytesRmndr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesRmndr(n)
	}
}

func BenchmarkBytesMask(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMask(n)
	}
}

func BenchmarkBytesMaskImpr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImpr(n)
	}
}

func BenchmarkBytesMaskImprSrc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImpRandSrc(n)
	}
}

func BenchmarkBytesMaskImprXorshift1024Src(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImpXorshift1024Src(n)
	}
}

func BenchmarkBytesMaskImprXorshift256Src(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImpXorshift256Src(n)
	}
}

func BenchmarkBytesMaskImprXorShift64StarSrc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImpXorShift64StarSrc(n)
	}
}

func BenchmarkBytesMaskImprXorShift128PlusSrc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImpXorShift128PlusSrc(n)
	}
}

func BenchmarkBytesMaskImprXorShift1024StarSrc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImpXorShift1024StarSrc(n)
	}
}

func SecureRandomAlphaString(length int) string {
	result := make([]byte, length)
	bufferSize := int(float64(length) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			randomBytes = SecureRandomBytes(bufferSize)
		}
		if idx := int(randomBytes[j%length] & letterIdxMask); idx < len(letterBytes) {
			result[i] = letterBytes[idx]
			i++
		}
	}

	return string(result)
}

func BenchmarkSecureRandomAlphaString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SecureRandomAlphaString(n)
	}
}

// SecureRandomBytes returns the requested number of bytes using crypto/rand
func SecureRandomBytes(length int) []byte {
	randomBytes := make([]byte, length)
	_, err := crand.Read(randomBytes)
	if err != nil {
		log.Fatal("Unable to generate random bytes")
	}
	return randomBytes
}

func BenchmarkSecureRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SecureRandomString(letterBytes, n)
	}
}

// SecureRandomString returns a string of the requested length,
// made from the byte characters provided (only ASCII allowed).
// Uses crypto/rand for security. Will panic if len(availableCharBytes) > 256.
func SecureRandomString(availableCharBytes string, length int) string {
	// Compute bitMask
	availableCharLength := len(availableCharBytes)
	if availableCharLength == 0 || availableCharLength > 256 {
		panic("availableCharBytes length must be greater than 0 and less than or equal to 256")
	}
	var bitLength byte
	var bitMask byte
	for bits := availableCharLength - 1; bits != 0; {
		bits = bits >> 1
		bitLength++
	}
	bitMask = 1<<bitLength - 1

	// Compute bufferSize
	bufferSize := length + length/3

	// Create random string
	result := make([]byte, length)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			// Random byte buffer is empty, get a new one
			randomBytes = SecureRandomBytes(bufferSize)
		}
		// Mask bytes to get an index into the character slice
		if idx := int(randomBytes[j%length] & bitMask); idx < availableCharLength {
			result[i] = availableCharBytes[idx]
			i++
		}
	}

	return string(result)
}

var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-"

func ShortID(length int) string {
	ll := len(chars)
	b := make([]byte, length)
	crand.Read(b) // generates len(b) random bytes
	for i := 0; i < length; i++ {
		b[i] = chars[int(b[i])%ll]
	}
	return string(b)
}

func BenchmarkShortID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShortID(n)
	}
}

var alphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generate(size int) string {
	b := make([]byte, size)
	rand.Read(b)
	for i := 0; i < size; i++ {
		b[i] = alphabet[b[i]/5]
	}
	return *(*string)(unsafe.Pointer(&b))
}

func TestGenerate(t *testing.T) {
	fmt.Println(generate(10))
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generate(n)
	}
}

func TestRandStr(t *testing.T) {
	fmt.Println(RandStr(10))
}

func BenchmarkRandStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStr(n)
	}
}

// https://github.com/dchest/uniuri/blob/master/uniuri.go

const (
	// StdLen is a standard length of uniuri string to achive ~95 bits of entropy.
	StdLen = 16
	// UUIDLen is a length of uniuri string to achive ~119 bits of entropy, closest
	// to what can be losslessly converted to UUIDv4 (122 bits).
	UUIDLen = 20
)

// StdChars is a set of standard characters allowed in uniuri string.
var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

// New returns a new random string of the standard length, consisting of
// standard characters.
func UniuriNew() string {
	return NewLenChars(StdLen, StdChars)
}

// NewLen returns a new random string of the provided length, consisting of
// standard characters.
func UniuriNewLen(length int) string {
	return NewLenChars(length, StdChars)
}

// NewLenChars returns a new random string of the provided length, consisting
// of the provided byte slice of allowed characters (maximum 256).
func NewLenChars(length int, chars []byte) string {
	if length == 0 {
		return ""
	}
	clen := len(chars)
	if clen < 2 || clen > 256 {
		panic("uniuri: wrong charset length for NewLenChars")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	i := 0
	for {
		if _, err := crand.Read(r); err != nil {
			panic("uniuri: error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				// Skip this number to avoid modulo bias.
				continue
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}

func BenchmarkUniuriNewLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UniuriNewLen(n)
	}
}

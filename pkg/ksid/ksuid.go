package ksid

import (
	"bytes"
	"crypto/rand"
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"time"
)

const (
	// EpochTimestamp KSID's epoch starts more recently so that the 32-bit number space gives a
	// significantly higher useful lifetime of around 136 years from March 2017.
	// This number (14e8) was picked to be easy to remember.
	EpochTimestamp int64 = 1400000000

	// TimestampSize is a uint32
	TimestampSize = 4

	// PayloadSize is 16-bytes
	PayloadSize = 16

	// Size KSUIDs are 20 bytes when binary encoded
	Size = TimestampSize + PayloadSize

	// The length of a KSID when string (base62) encoded
	stringEncodedLength = 27

	// A string-encoded minimum value for a KSID
	minStringEncoded = "000000000000000000000000000"

	// A string-encoded maximum value for a KSID
	maxStringEncoded = "aWgEPTl1tmebfsQzFP4bxwgy80V"
)

// KSID are 20 bytes:
//
//	00-03 byte: uint32 be UTC timestamp with custom epoch
//	04-19 byte: random "payload"
type KSID struct {
	Data  [Size]byte
	Error error
}

var (
	rander = rand.Reader

	errSize        = fmt.Errorf("valid KSUIDs are %v bytes", Size)
	errStrSize     = fmt.Errorf("valid encoded KSUIDs are %v characters", stringEncodedLength)
	errStrValue    = fmt.Errorf("valid encoded KSUIDs are bounded by %s and %s", minStringEncoded, maxStringEncoded)
	errPayloadSize = fmt.Errorf("valid KSID payloads are %v bytes", PayloadSize)

	// Nil Represents a completely empty (invalid) KSID
	Nil KSID
	// Max Represents the highest value a KSID can have
	Max = KSID{Data: [Size]byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}}
)

// Append appends the string representation of i to b, returning a slice to a
// potentially larger memory area.
func (i KSID) Append(b []byte) []byte { return fastAppendEncodeBase62(b, i.Data[:]) }

// Time The timestamp portion of the ID as a Time object
func (i KSID) Time() time.Time { return time.Unix(int64(i.Timestamp())+EpochTimestamp, 0) }

// Timestamp The timestamp portion of the ID as a bare integer which is uncorrected for KSID's special epoch.
func (i KSID) Timestamp() uint32 { return binary.BigEndian.Uint32(i.Data[:TimestampSize]) }

// Payload The 16-byte random payload without the timestamp
func (i KSID) Payload() []byte { return i.Data[TimestampSize:] }

// String-encoded representation that can be passed through Parse()
func (i KSID) String() string { return string(i.Append(make([]byte, 0, stringEncodedLength))) }

// Bytes Raw byte representation of KSID.
func (i KSID) Bytes() []byte { /* Safe because this is by-value*/ return i.Data[:] }

// IsNil returns true if this is a "nil" KSID
func (i KSID) IsNil() bool { return i == Nil }

// Get satisfies the flag.Getter interface, making it possible to use KSUIDs as
// part of the command line options of a program.
func (i KSID) Get() interface{} { return i }

// Set satisfies the flag.Value interface, making it possible to use KSUIDs as
// part of the command line options of a program.
func (i *KSID) Set(s string) error { return i.UnmarshalText([]byte(s)) }

func (i KSID) MarshalText() ([]byte, error)   { return []byte(i.String()), nil }
func (i KSID) MarshalBinary() ([]byte, error) { return i.Bytes(), nil }
func (i *KSID) UnmarshalText(b []byte) error {
	id, err := Parse(string(b))
	if err != nil {
		return err
	}
	*i = id
	return nil
}

func (i *KSID) UnmarshalBinary(b []byte) error {
	id := New(WithBytes(b))
	if id.Error != nil {
		return id.Error
	}
	*i = id
	return nil
}

// Value converts the KSID into a SQL driver value which can be used to
// directly use the KSID as parameter to a SQL query.
func (i KSID) Value() (driver.Value, error) {
	if i.IsNil() {
		return nil, nil
	}
	return i.String(), nil
}

// Scan implements the sql.Scanner interface. It supports converting from
// string, []byte, or nil into a KSID value. Attempting to convert from
// another type will return an error.
func (i *KSID) Scan(src interface{}) error {
	switch v := src.(type) {
	case nil:
		return i.scan(nil)
	case []byte:
		return i.scan(v)
	case string:
		return i.scan([]byte(v))
	default:
		return fmt.Errorf("Scan: unable to scan type %T into KSID", v)
	}
}

func (i *KSID) scan(b []byte) error {
	switch len(b) {
	case 0:
		*i = Nil
		return nil
	case Size:
		return i.UnmarshalBinary(b)
	case stringEncodedLength:
		return i.UnmarshalText(b)
	default:
		return errSize
	}
}

// Parse decodes a string-encoded representation of a KSID object
func Parse(s string) (KSID, error) {
	if len(s) != stringEncodedLength {
		return Nil, errStrSize
	}

	src := [stringEncodedLength]byte{}
	dst := [Size]byte{}

	copy(src[:], s[:])

	if err := fastDecodeBase62(dst[:], src[:]); err != nil {
		return Nil, errStrValue
	}

	k := New(WithBytes(dst[:]))
	return k, k.Error
}

type Config struct {
	KSID
	PayloadSet   bool
	RandomSet    bool
	TimestampSet bool
}

type ConfigFn func(*Config)

func WithTime(t time.Time) ConfigFn {
	return WithTimeStamp(t.Unix())
}

func WithRandom(t []byte) ConfigFn {
	return func(c *Config) {
		copy(c.KSID.Data[TimestampSize:], t)
		c.RandomSet = true
	}
}

func WithTimeStamp(t int64) ConfigFn {
	return func(c *Config) {
		ts := uint32(t - EpochTimestamp)
		binary.BigEndian.PutUint32(c.Data[:TimestampSize], ts)
		c.TimestampSet = true
	}
}

func WithValue(t KSID) ConfigFn { return WithBytes(t.Bytes()) }

func WithBytes(t []byte) ConfigFn {
	return func(c *Config) {
		if len(t) != Size {
			c.Error = errPayloadSize
		}
		copy(c.Data[:], t)
		c.PayloadSet = true
	}
}

// New generates a new KSID.
func New(fns ...ConfigFn) KSID {
	c := Config{}
	for _, f := range fns {
		f(&c)
	}

	if c.PayloadSet {
		return c.KSID
	}

	if !c.TimestampSet {
		ts := uint32(time.Now().Unix() - EpochTimestamp)
		binary.BigEndian.PutUint32(c.Data[:TimestampSize], ts)
	}

	if !c.RandomSet {
		if _, err := io.ReadAtLeast(rander, c.Data[TimestampSize:], PayloadSize); err != nil {
			c.Error = err
		}
	}

	return c.KSID
}

// SetRand Sets the global source of random bytes for KSID generation. This
// should probably only be set once globally. While this is technically
// thread-safe as in it won't cause corruption, there's no guarantee
// on ordering.
func SetRand(r io.Reader) {
	if r == nil {
		rander = rand.Reader
		return
	}
	rander = r
}

// Compare Implements comparison for KSID type
func Compare(a, b KSID) int {
	return bytes.Compare(a.Data[:], b.Data[:])
}

// Sort Sorts the given slice of KSUIDs
func Sort(ids []KSID) {
	quickSort(ids, 0, len(ids)-1)
}

// IsSorted checks whether a slice of KSUIDs is sorted
func IsSorted(ids []KSID) bool {
	if len(ids) != 0 {
		min := ids[0]
		for _, id := range ids[1:] {
			if Compare(min, id) > 0 {
				return false
			}
			min = id
		}
	}
	return true
}

func quickSort(a []KSID, lo int, hi int) {
	if lo < hi {
		pivot := a[hi]
		i := lo - 1

		for j, n := lo, hi; j != n; j++ {
			if Compare(a[j], pivot) < 0 {
				i++
				a[i], a[j] = a[j], a[i]
			}
		}

		i++
		if Compare(a[hi], a[i]) < 0 {
			a[i], a[hi] = a[hi], a[i]
		}

		quickSort(a, lo, i-1)
		quickSort(a, i+1, hi)
	}
}

// Next returns the next KSID after id.
func (i KSID) Next() KSID {
	zero := makeUint128(0, 0)

	t := i.Timestamp()
	u := uint128Payload(i)
	v := add128(u, makeUint128(0, 1))

	if v == zero { // overflow
		t++
	}

	return v.ksuid(t)
}

// Prev returns the previoud KSID before id.
func (i KSID) Prev() KSID {
	max := makeUint128(math.MaxUint64, math.MaxUint64)

	t := i.Timestamp()
	u := uint128Payload(i)
	v := sub128(u, makeUint128(0, 1))

	if v == max { // overflow
		t--
	}

	return v.ksuid(t)
}

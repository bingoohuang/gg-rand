package cid

import (
	"crypto/rand"
	"encoding/binary"
	"io"
	"time"
)

var epoch = func() time.Time {
	t, err := time.Parse("20060102", "20220209")
	if err != nil {
		panic(err)
	}

	return t
}

// Cid12 will return the fixed length of 12 number.
// 13 bits day + 27 bits random
func Cid12() uint64 {
	buf := [8]byte{}
	io.ReadFull(rand.Reader, buf[:])
	u := binary.BigEndian.Uint64(buf[:])
	v := uint64(time.Now().Sub(epoch()).Hours())/24 + 746 // 746 is to keep at string length 12
	return (v << 27) | (u & 0b111111111111111111111111111)
}

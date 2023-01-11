package hash

import (
	"crypto/md5"
	"hash"
	"io"
	"os"

	"github.com/cespare/xxhash/v2"
	"github.com/kalafut/imohash"
	"github.com/zeebo/blake3"
	blake3Luke "lukechampine.com/blake3"
)

// HashFile returns the  64-bit xxHash algorithm, XXH64 of a file.
func HashFile(fname string, h hash.Hash) ([]byte, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if _, err := io.Copy(h, f); err != nil {
		return nil, err
	}

	return h.Sum(nil), err
}

// XXH64File returns the 64-bit xxHash algorithm, XXH64 of a file.
func XXH64File(fname string) ([]byte, error) {
	return HashFile(fname, xxhash.New())
}

// MD5HashFile returns MD5 hash.
func MD5HashFile(fname string) ([]byte, error) {
	return HashFile(fname, md5.New())
}

// IMOHashFile returns imohash.
func IMOHashFile(fname string) ([]byte, error) {
	b, err := imohash.SumFile(fname)
	return b[:], err
}

// Blake3Zeebo return blake3 hash of a file.
// https://github.com/KEINOS/go-blake3-example
func Blake3Zeebo(fname string) ([]byte, error) {
	return HashFile(fname, blake3.New())
}

// Blake3Luke return blake3 hash of a file.
func Blake3Luke(fname string) ([]byte, error) {
	return HashFile(fname, blake3Luke.New(256, nil))
}

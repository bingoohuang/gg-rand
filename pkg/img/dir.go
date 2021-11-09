package img

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/bingoohuang/gg/pkg/uid"
)

var Dir string

func init() {
	d, err := ioutil.TempDir("", "")
	if err != nil {
		panic(err)
	}
	Dir = d
}

// ToPng saves the image to local with PNG format.
func ToPng(img image.Image, appendBase64 bool) string {
	file := filepath.Join(Dir, uid.New().String()+".png")
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0o755)
	if err != nil {
		log.Printf("failed to open file %s, error: %v", file, err)
		return ""
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		log.Printf("failed to png.Encode, error: %v", err)
		return ""
	}

	result := f.Name()

	if appendBase64 {
		var buf bytes.Buffer
		if err := png.Encode(&buf, img); err != nil {
			return ""
		}
		data := buf.Bytes()
		result += " base64: " + base64.StdEncoding.EncodeToString(data)
	}

	return result
}

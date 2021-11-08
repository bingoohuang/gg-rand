package dir

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"io/ioutil"
)

var TempDir string

func init() {
	d, err := ioutil.TempDir("", "")
	if err != nil {
		panic(err)
	}
	TempDir = d
}

// ToPng saves the image to local with PNG format.
func ToPng(img image.Image, appendBase64 bool) string {
	f, err := ioutil.TempFile(TempDir, "rand*.png")
	if err != nil {
		return ""
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
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

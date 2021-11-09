package img

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"io/ioutil"
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
	f, err := os.OpenFile(filepath.Join(Dir, uid.New().String()+".png"), os.O_RDWR|os.O_CREATE, 0o755)
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

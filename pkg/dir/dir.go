package dir

import (
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
func ToPng(m image.Image) string {
	f, err := ioutil.TempFile(TempDir, "rand*.png")
	if err != nil {
		return ""
	}
	defer f.Close()

	png.Encode(f, m)

	return f.Name()
}

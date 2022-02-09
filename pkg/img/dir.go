package img

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bingoohuang/gg/pkg/man"
	"github.com/bingoohuang/gg/pkg/randx"
	"github.com/bingoohuang/gg/pkg/ss"

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

var randInt = randx.Int()

// RandomImage creates a random image.
// Environment variables supported:
// GG_IMG_FAST=Y/N to enable fast mode or not
// GG_IMG_FORMAT=jpg/png to choose the format
// GG_IMG_FILE_SIZE=10M to set image file size
// GG_IMG_SIZE=640x320 to set the {width}x{height} of image
func RandomImage(i int) interface{} {
	imgFormat := parseImageFormat("GG_IMG_FORMAT")
	width, height := parseImageSize("GG_IMG_SIZE")
	fn := fmt.Sprintf("%d_%dx%d%s", randInt+i, width, height, imgFormat)
	c := randx.ImgConfig{
		Width:      width,
		Height:     height,
		RandomText: fmt.Sprintf("%d", randInt+i),
		FastMode:   parseImageFastMode("GG_IMG_FAST"),
	}
	size := c.GenFile(fn, int(parseImageFileSize("GG_IMG_FILE_SIZE")))

	return man.IBytes(uint64(size)) + " " + fn
}

func parseImageFastMode(envName string) bool {
	if val := os.Getenv(envName); val != "" {
		if v, err := ss.ParseBoolE(val); err == nil {
			return v
		}
	}

	return true
}

func parseImageFileSize(envName string) (v uint64) {
	if val := os.Getenv(envName); val != "" {
		v, _ = man.ParseBytes(val)
	}
	return v
}

func parseImageFormat(envName string) string {
	if v := os.Getenv(envName); v != "" {
		switch strings.ToLower(v) {
		case ".jpg", "jpg", ".jpeg", "jpeg":
			return ".jpg"
		case ".png", "png":
			return ".png"
		}
	}
	return ss.If(randx.Bool(), ".jpg", ".png")
}

func parseImageSize(envName string) (width, height int) {
	width, height = 640, 320
	if val := os.Getenv(envName); val != "" {
		val = strings.ToLower(val)
		parts := strings.SplitN(val, "x", 2)
		if len(parts) == 2 {
			if v := ss.ParseInt(parts[0]); v > 0 {
				width = v
			}
			if v := ss.ParseInt(parts[1]); v > 0 {
				height = v
			}
		}
	}
	return width, height
}

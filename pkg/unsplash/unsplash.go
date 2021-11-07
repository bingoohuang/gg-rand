package unsplash

import (
	"fmt"
	"image"
	"net/http"

	"github.com/bingoohuang/gg-rand/pkg/dir"
)

func Random() string {
	img, err := RandomImage(1000, 1000)
	if err != nil {
		panic(err)
	}

	return dir.ToPng(img)
}

// RandomImage https://preslav.me/2021/01/11/an-easy-way-to-get-a-random-stream-of-images-in-your-golang-app/.
func RandomImage(width, height int) (image.Image, error) {
	url := fmt.Sprintf("https://source.unsplash.com/random/%dx%d", width, height)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	img, _, err := image.Decode(res.Body)
	return img, err
}

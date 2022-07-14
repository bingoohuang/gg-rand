package colors

import (
	"fmt"
	"github.com/bingoohuang/gg-rand/pkg/img"
	"github.com/lucasb-eyer/go-colorful"
	"image"
	"image/draw"
	"math/rand"
	"time"
)

func ColorBlend(int) interface{} {
	blocks := 10
	blockw := 40
	rect := image.NewRGBA(image.Rect(0, 0, blocks*blockw, 200))

	c1, _ := colorful.Hex("#fdffcc")
	c2, _ := colorful.Hex("#242a42")

	// Use these colors to get invalid RGB in the gradient.
	//c1, _ := colorful.Hex("#EEEF61")
	//c2, _ := colorful.Hex("#1E3140")

	for i := 0; i < blocks; i++ {
		draw.Draw(rect, image.Rect(i*blockw, 0, (i+1)*blockw, 40),
			&image.Uniform{C: c1.BlendHsv(c2, float64(i)/float64(blocks-1))}, image.Point{}, draw.Src)
		draw.Draw(rect, image.Rect(i*blockw, 40, (i+1)*blockw, 80),
			&image.Uniform{C: c1.BlendLuv(c2, float64(i)/float64(blocks-1))}, image.Point{}, draw.Src)
		draw.Draw(rect, image.Rect(i*blockw, 80, (i+1)*blockw, 120),
			&image.Uniform{C: c1.BlendRgb(c2, float64(i)/float64(blocks-1))}, image.Point{}, draw.Src)
		draw.Draw(rect, image.Rect(i*blockw, 120, (i+1)*blockw, 160),
			&image.Uniform{C: c1.BlendLab(c2, float64(i)/float64(blocks-1))}, image.Point{}, draw.Src)
		draw.Draw(rect, image.Rect(i*blockw, 160, (i+1)*blockw, 200),
			&image.Uniform{C: c1.BlendHcl(c2, float64(i)/float64(blocks-1))}, image.Point{}, draw.Src)

		// This can be used to "fix" invalid colors in the gradient.
		//draw.Draw(img, image.Rect(i*blockw,160,(i+1)*blockw,200), &image.Uniform{c1.BlendHcl(c2, float64(i)/float64(blocks-1)).Clamped()}, image.Point{}, draw.Src)
	}

	return img.ToPng(rect, false)
}
func RandomPalettes(int) interface{} {
	colors := 10
	blockw := 40
	space := 5

	rand.Seed(time.Now().UTC().UnixNano())
	rect := image.NewRGBA(image.Rect(0, 0, colors*blockw+space*(colors-1), 6*blockw+8*space))

	warm, err := colorful.WarmPalette(colors)
	if err != nil {
		return fmt.Sprintf("Error generating warm palette: %v", err)
	}
	fwarm := colorful.FastWarmPalette(colors)
	happy, err := colorful.HappyPalette(colors)
	if err != nil {
		return fmt.Sprintf("Error generating happy palette: %v", err)
	}
	fhappy := colorful.FastHappyPalette(colors)
	soft, err := colorful.SoftPalette(colors)
	if err != nil {
		return fmt.Sprintf("Error generating soft palette: %v", err)
	}
	brownies, err := colorful.SoftPaletteEx(colors,
		colorful.SoftPaletteSettings{CheckColor: isbrowny, Iterations: 50, ManySamples: true})
	if err != nil {
		return fmt.Sprintf("Error generating brownies: %v", err)
	}
	for i := 0; i < colors; i++ {
		draw.Draw(rect, image.Rect(i*(blockw+space), 0, i*(blockw+space)+blockw, blockw),
			&image.Uniform{C: warm[i]}, image.Point{}, draw.Src)
		draw.Draw(rect, image.Rect(i*(blockw+space), 1*blockw+1*space, i*(blockw+space)+blockw, 2*blockw+1*space),
			&image.Uniform{C: fwarm[i]}, image.Point{}, draw.Src)
		draw.Draw(rect, image.Rect(i*(blockw+space), 2*blockw+3*space, i*(blockw+space)+blockw, 3*blockw+3*space),
			&image.Uniform{C: happy[i]}, image.Point{}, draw.Src)
		draw.Draw(rect, image.Rect(i*(blockw+space), 3*blockw+4*space, i*(blockw+space)+blockw, 4*blockw+4*space),
			&image.Uniform{C: fhappy[i]}, image.Point{}, draw.Src)
		draw.Draw(rect, image.Rect(i*(blockw+space), 4*blockw+6*space, i*(blockw+space)+blockw, 5*blockw+6*space),
			&image.Uniform{C: soft[i]}, image.Point{}, draw.Src)
		draw.Draw(rect, image.Rect(i*(blockw+space), 5*blockw+8*space, i*(blockw+space)+blockw, 6*blockw+8*space),
			&image.Uniform{C: brownies[i]}, image.Point{}, draw.Src)
	}

	return img.ToPng(rect, false)
}

func isbrowny(l, a, b float64) bool {
	h, c, L := colorful.LabToHcl(l, a, b)
	return 10.0 < h && h < 50.0 && 0.1 < c && c < 0.5 && L < 0.5
}

func RandomColors(int) interface{} {
	blocks := 10
	blockw := 40
	space := 5

	rand.Seed(time.Now().UTC().UnixNano())
	rect := image.NewRGBA(image.Rect(0, 0, blocks*blockw+space*(blocks-1), 4*(blockw+space)))

	for i := 0; i < blocks; i++ {
		draw.Draw(rect, image.Rect(i*(blockw+space), 0, i*(blockw+space)+blockw, blockw),
			&image.Uniform{C: colorful.WarmColor()}, image.Point{}, draw.Src)
		draw.Draw(rect, image.Rect(i*(blockw+space), blockw+space, i*(blockw+space)+blockw, 2*blockw+space),
			&image.Uniform{C: colorful.FastWarmColor()}, image.Point{}, draw.Src)
		draw.Draw(rect, image.Rect(i*(blockw+space), 2*blockw+3*space, i*(blockw+space)+blockw, 3*blockw+3*space),
			&image.Uniform{C: colorful.HappyColor()}, image.Point{}, draw.Src)
		draw.Draw(rect, image.Rect(i*(blockw+space), 3*blockw+4*space, i*(blockw+space)+blockw, 4*blockw+4*space),
			&image.Uniform{C: colorful.FastHappyColor()}, image.Point{}, draw.Src)
	}

	return img.ToPng(rect, false)
}

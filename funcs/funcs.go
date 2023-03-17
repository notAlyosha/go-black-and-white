package funcs

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

var path string

func BlackAndWhite() {
	fmt.Fscan(os.Stdin, &path)

	imgIn, _ := os.Open(path)
	img, _ := jpeg.Decode(imgIn)
	imgIn.Close()

	imgOut := image.NewGray16(img.Bounds().Bounds())

	for y := 0; y < imgOut.Bounds().Dy(); y += 1 {
		for x := 0; x < imgOut.Bounds().Dx(); x += 1 {
			currentPixel := img.At(x, y)
			r, g, b, _ := currentPixel.RGBA()

			var gray color.Gray16 = color.Gray16{uint16((r + g + b) / 3)}

			imgOut.SetGray16(x, y, gray)
		}
	}

	f, _ := os.Create("black_and_whitte")
	jpeg.Encode(f, imgOut, &jpeg.Options{Quality: 16})
}

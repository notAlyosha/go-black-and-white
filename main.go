package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		panic("I need path to image")
	}

	path := os.Args[1]
	imgIn, _ := os.Open(path)
	img, _ := jpeg.Decode(imgIn)

	imgOut := image.NewGray16(img.Bounds().Bounds())

	for y := 0; y < imgOut.Bounds().Dy(); y += 1 {
		for x := 0; x < imgOut.Bounds().Dx(); x += 1 {
			currentPixel := img.At(x, y)
			r, g, b, _ := currentPixel.RGBA()

			gray := color.Gray16{uint16((r + g + b) / 3)}
			imgOut.SetGray16(x, y, gray)
		}
	}

	f, _ := os.Create("b")
	jpeg.Encode(f, imgOut, &jpeg.Options{Quality: 16})
}

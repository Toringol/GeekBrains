package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	col := color.RGBA{0, 255, 0, 255} // green
	img := image.NewRGBA(image.Rect(0, 0, 200, 200))

	draw.Draw(img, img.Bounds(), &image.Uniform{col}, image.ZP, draw.Src)

	col = color.RGBA{255, 0, 0, 255} // Red
	HLine(0, 100, 200, img, col)
	VLine(100, 0, 200, img, col)

	file, err := os.Create("draw.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, img)
}

// HLine draws a horizontal line
func HLine(x1, y, x2 int, img *image.RGBA, col color.Color) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

// VLine draws a veritcal line
func VLine(x, y1, y2 int, img *image.RGBA, col color.Color) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

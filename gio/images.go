package gio

import (
	"image"
	"image/draw"
	"image/png"
	"os"
)

func LoadPNG(path string) (data []uint8, width, height int32) {
	reader, loadErr := os.Open(path)
	if loadErr != nil {
		panic(loadErr)
	}
	img, decodeErr := png.Decode(reader)
	if decodeErr != nil {
		panic(decodeErr)
	}
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
	return rgba.Pix, int32(rgba.Rect.Size().X), int32(rgba.Rect.Size().Y)
}

func LoadIcon(path string) (image image.Image) {
	reader, loadErr := os.Open(path)
	if loadErr != nil {
		panic(loadErr)
	}
	img, decodeErr := png.Decode(reader)
	if decodeErr != nil {
		panic(decodeErr)
	}
	return img
}

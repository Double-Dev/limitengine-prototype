package gio

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

type Image struct {
	rgba *image.RGBA
}

func (image *Image) Data() []uint8           { return image.rgba.Pix }
func (image *Image) Width() int32            { return int32(image.rgba.Rect.Size().X) }
func (image *Image) Height() int32           { return int32(image.rgba.Rect.Size().Y) }
func (image *Image) ColorModel() color.Model { return image.rgba.ColorModel() }
func (image *Image) Bounds() image.Rectangle { return image.rgba.Bounds() }
func (image *Image) At(x, y int) color.Color { return image.rgba.At(x, y) }

func LoadPNG(path string) *Image {
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
	return &Image{rgba}
}

func LoadJPEG(path string) *Image {
	reader, loadErr := os.Open(path)
	if loadErr != nil {
		panic(loadErr)
	}
	img, decodeErr := jpeg.Decode(reader)
	if decodeErr != nil {
		panic(decodeErr)
	}
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
	return &Image{rgba}
}

package gio

import (
	"strings"

	"github.com/double-dev/limitengine/gmath"
)

const (
	paddingTop    = 0
	paddingLeft   = 1
	paddingBottom = 2
	paddingRight  = 3
)

type Char struct {
	bounds  gmath.Vector4
	offset  gmath.Vector2
	advance float32
	page    int32
}

func (char *Char) Bounds() gmath.Vector4 { return char.bounds }

type Font struct {
	pages      []*Image
	atlas      map[string]*Char
	padding    []float32
	lineHeight float32
}

func (font *Font) Pages() []*Image          { return font.pages }
func (font *Font) GetChar(key string) *Char { return font.atlas[key] }
func (font *Font) Padding() []float32       { return font.padding }
func (font *Font) LineHeight() float32      { return font.lineHeight }

func LoadFNT(directoryPath, fileName string) *Font {
	fileStr := LoadAsString(directoryPath + "/" + fileName)
	fileStr = strings.ReplaceAll(fileStr, "\r", "")
	fileStrs := strings.Split(fileStr, "\n")

	var scaleW, scaleH float32
	font := &Font{
		atlas: make(map[string]*Char),
	}

	for _, lineStr := range fileStrs {
		lineStrs := strings.Split(lineStr, " ")

		switch lineStrs[0] {
		case "info":
			for _, varStr := range lineStrs[1:] {
				varStrs := strings.Split(varStr, "=")
				switch varStrs[0] {
				case "padding":
					values := strings.Split(varStrs[1], ",")
					for _, value := range values {
						font.padding = append(font.padding, float32(parseInt(value)))
					}
					break
				}
			}
			break
		case "common":
			for _, varStr := range lineStrs[1:] {
				varStrs := strings.Split(varStr, "=")
				switch varStrs[0] {
				case "lineHeight":
					font.lineHeight = float32(parseInt(varStrs[1]))
					break
				case "scaleW":
					scaleW = float32(parseInt(varStrs[1]))
					font.padding[paddingLeft] /= scaleW
					font.padding[paddingRight] /= scaleW
					break
				case "scaleH":
					scaleH = float32(parseInt(varStrs[1]))
					font.padding[paddingTop] /= scaleH
					font.padding[paddingBottom] /= scaleH
					font.lineHeight /= scaleH
					break
				}
			}
			break
		case "page":
			for _, varStr := range lineStrs[1:] {
				varStrs := strings.Split(varStr, "=")
				switch varStrs[0] {
				case "file":
					font.pages = append(font.pages, LoadPNG(directoryPath+"/"+strings.ReplaceAll(varStrs[1], "\"", "")))
					break
				}
			}
			break
		case "char":
			var key string
			char := &Char{
				bounds: gmath.NewZeroVector4(),
				offset: gmath.NewZeroVector2(),
			}
			for _, varStr := range lineStrs[1:] {
				varStrs := strings.Split(varStr, "=")

				switch varStrs[0] {
				case "id":
					key = string(parseInt(varStrs[1]))
					break
				case "x":
					char.bounds[0] = float32(parseInt(varStrs[1])) / scaleW
					break
				case "y":
					char.bounds[1] = float32(parseInt(varStrs[1])) / scaleH
					break
				case "width":
					char.bounds[2] = float32(parseInt(varStrs[1])) / scaleW
					break
				case "height":
					char.bounds[3] = float32(parseInt(varStrs[1])) / scaleH
					break
				case "xoffset":
					char.offset[0] = float32(parseInt(varStrs[1])) / scaleW
					break
				case "yoffset":
					char.offset[1] = float32(parseInt(varStrs[1])) / scaleH
					break
				case "xadvance":
					char.advance = float32(parseInt(varStrs[1])) / scaleW
					break
				case "page":
					char.page = parseInt(varStrs[1])
					break
				}
			}
			font.atlas[key] = char
			break
		}
	}
	return font
}

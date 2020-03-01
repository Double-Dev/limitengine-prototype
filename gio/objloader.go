package gio

import (
	"strconv"
	"strings"
)

func LoadOBJ(path string) ([]uint32, []float32, []float32, []float32) {
	fileStr := LoadAsString(path)
	fileStr = strings.ReplaceAll(fileStr, "\r", "")
	fileStrs := strings.Split(fileStr, "\n")

	lineIndex := 0
	var line string
	var indices []uint32
	var rawVertices, rawTextureCoords, rawNormals []float32
	var outVertices, outTextureCoords, outNormals []float32

	for lineIndex < len(fileStrs)-1 {
		line = fileStrs[lineIndex]
		currentLine := strings.Split(line, " ")
		if currentLine[0] == "v" {
			x, err := strconv.ParseFloat(currentLine[1], 32)
			if err != nil {
				panic(err)
			}
			y, err := strconv.ParseFloat(currentLine[2], 32)
			if err != nil {
				panic(err)
			}
			z, err := strconv.ParseFloat(currentLine[3], 32)
			if err != nil {
				panic(err)
			}
			rawVertices = append(rawVertices, float32(x), float32(y), float32(z))
		} else if currentLine[0] == "vt" {
			x, err := strconv.ParseFloat(currentLine[1], 32)
			if err != nil {
				panic(err)
			}
			y, err := strconv.ParseFloat(currentLine[2], 32)
			if err != nil {
				panic(err)
			}
			rawTextureCoords = append(rawTextureCoords, float32(x), float32(y))
		} else if currentLine[0] == "vn" {
			x, err := strconv.ParseFloat(currentLine[1], 32)
			if err != nil {
				panic(err)
			}
			y, err := strconv.ParseFloat(currentLine[2], 32)
			if err != nil {
				panic(err)
			}
			z, err := strconv.ParseFloat(currentLine[3], 32)
			if err != nil {
				panic(err)
			}
			rawNormals = append(rawNormals, float32(x), float32(y), float32(z))
		} else if currentLine[0] == "f" {
			for i := 1; i < 4; i++ {
				vertex := strings.Split(currentLine[i], "/")
				vertIndex, err := strconv.ParseUint(vertex[0], 10, 32)
				if err != nil {
					panic(err)
				}
				texIndex, err := strconv.ParseUint(vertex[1], 10, 32)
				if err != nil {
					panic(err)
				}
				normIndex, err := strconv.ParseUint(vertex[2], 10, 32)
				if err != nil {
					panic(err)
				}
				offsetVertIndex := (vertIndex - 1) * 3
				outVertices = append(outVertices, rawVertices[offsetVertIndex], rawVertices[offsetVertIndex+1], rawVertices[offsetVertIndex+2])
				offsetTexIndex := (texIndex - 1) * 2
				outTextureCoords = append(outTextureCoords, rawTextureCoords[offsetTexIndex], rawTextureCoords[offsetTexIndex+1])
				offsetNormIndex := (normIndex - 1) * 3
				outNormals = append(outNormals, rawNormals[offsetNormIndex], rawNormals[offsetNormIndex+1], rawNormals[offsetNormIndex+2])
				indices = append(indices, uint32(len(indices)))
			}
		}
		lineIndex++
	}
	// TODO: Optimize loaded mesh.
	// fmt.Println(indices)
	// fmt.Println(outVertices)
	// fmt.Println(outTextureCoords)
	// fmt.Println(outNormals)
	return indices, outVertices, outTextureCoords, outNormals
}

func parseOBJIndex(token string) {
	// values := strings.Split(token, "/")

}

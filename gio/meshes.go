package gio

import (
	"fmt"
	"strconv"
	"strings"
)

func LoadOBJ(path string) (outIndices []uint32, outVertices, outTextureCoords, outNormals []float32) {
	fileStr := LoadAsString(path)
	fileStrs := strings.Split(fileStr, "\n")

	lineIndex := 0
	var line string
	var indices []uint32
	var vertices, textureCoords, normals []float32

	for {
		line = fileStrs[lineIndex]
		currentLine := strings.Split(line, " ")
		if currentLine[0] == "v" {
			x, _ := strconv.ParseFloat(currentLine[1], 32)
			y, _ := strconv.ParseFloat(currentLine[2], 32)
			z, _ := strconv.ParseFloat(currentLine[3], 32)
			vertices = append(vertices, float32(x), float32(y), float32(z))
		} else if currentLine[0] == "vt" {
			x, _ := strconv.ParseFloat(currentLine[1], 32)
			y, _ := strconv.ParseFloat(currentLine[2], 32)
			textureCoords = append(textureCoords, float32(x), float32(y))
		} else if currentLine[0] == "vn" {
			x, _ := strconv.ParseFloat(currentLine[1], 32)
			y, _ := strconv.ParseFloat(currentLine[2], 32)
			z, _ := strconv.ParseFloat(currentLine[3], 32)
			normals = append(normals, float32(x), float32(y), float32(z))
		} else if currentLine[0] == "f" {
			fmt.Println("break")
			break
		}
		lineIndex++
	}
	for lineIndex < len(fileStrs)-1 {
		line = fileStrs[lineIndex]
		currentLine := strings.Split(line, " ")
		for i := 1; i < 4; i++ {
			vertex := strings.Split(currentLine[i], "/")
			indice, _ := strconv.ParseUint(vertex[0], 10, 32)
			indices = append(indices, uint32(indice))
		}
		lineIndex++
	}
	fmt.Println("hello")
	return indices, vertices, textureCoords, normals
}

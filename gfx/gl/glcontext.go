package gl

import (
	"doubledev/limitengine/gfx/framework"
	"fmt"

	"github.com/go-gl/gl/v3.2-core/gl"
)

type glContext struct{}

func NewGLContext() (glContext, error) {
	var err error
	if err = gl.Init(); err != nil {
		err = fmt.Errorf("Error starting OpenGL instance: %s", err)
	}
	// TODO: Add options for opengl features.
	gl.Enable(gl.MULTISAMPLE)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	return glContext{}, err
}

func (glContext glContext) Resize(width, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}

func (glContext glContext) ClearScreen(r, g, b, a float32) {
	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(r, g, b, a)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (glContext glContext) CreateFrameBuffer() {

}

func (glContext glContext) CreateRenderBuffer() {

}

func (glContext glContext) CreateShader(vertSrc, fragSrc string) framework.IShader {
	shader := newShader(vertSrc, fragSrc)
	return shader
}

func (glContext glContext) CreateModel(indices []uint32, vertices, texCoords, normals []float32) framework.IModel {
	model := newVAO()
	model.bind()
	model.addIndicesArr(indices)
	model.addFloatAttribArr(vertices, 0, 3, false)
	model.addFloatAttribArr(texCoords, 1, 2, false)
	model.addFloatAttribArr(normals, 2, 3, false)
	model.unbind()
	return model
}

func (glContext glContext) CreateTexture(image []uint8, width, height int32) framework.ITexture {
	texture := newTexture(image, width, height)
	return texture
}

// var testTextureID uint32

// func CreateTextureTest() {
// 	reader, _ := os.Open("testIcon.png")
// 	img, _ := png.Decode(reader)
// 	rgba := image.NewRGBA(img.Bounds())
// 	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

// 	gl.GenTextures(1, &testTextureID)
// 	gl.ActiveTexture(gl.TEXTURE0)
// 	gl.BindTexture(gl.TEXTURE_2D, testTextureID)
// 	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(rgba.Rect.Size().X), int32(rgba.Rect.Size().Y), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))
// }

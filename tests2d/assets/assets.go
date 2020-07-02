package assets

import (
	"github.com/double-dev/limitengine/gfx"
	"github.com/double-dev/limitengine/gio"
	"github.com/double-dev/limitengine/gmath"
	"github.com/double-dev/limitengine/utils2d"
)

var (
	// Cameras
	SceneCamera  *gfx.Camera
	PostTexture  *gfx.Texture
	PostMaterial *gfx.TextureMaterial
	PostCamera   *gfx.Camera

	// Shaders
	SceneShader *utils2d.SpriteShader
	TextShader  *gfx.TextShader
	PostShader  *gfx.GenericShader

	// Level
	LevelMaterial *gfx.ColorMaterial

	// Player
	PlayerTexture     *gfx.Texture
	PlayerMaterial    *gfx.TextureMaterial
	PlayerSpriteSheet *gfx.SpriteSheet

	PlayerRightIdle      *gfx.FrameAnimation
	PlayerRightWalk      *gfx.FrameAnimation
	PlayerRightWallSlide *gfx.FrameAnimation

	PlayerLeftIdle      *gfx.FrameAnimation
	PlayerLeftWalk      *gfx.FrameAnimation
	PlayerLeftWallSlide *gfx.FrameAnimation

	// Font
	SegoeFont     *gio.Font
	SegoeTexture  *gfx.Texture
	SegoeMaterial *gfx.TextureMaterial
)

func LoadAssets() {
	// Cameras
	SceneCamera = gfx.NewCamera2D(gfx.NewRenderbuffer(true), gfx.NewRenderbuffer(true))
	SceneCamera.SetClearColor(0.9, 0.9, 0.9, 1.0)
	PostTexture = gfx.NewEmptyTexture()
	PostMaterial = gfx.NewTextureMaterial(PostTexture)
	PostCamera = gfx.NewCamera2D(PostTexture, gfx.NewRenderbuffer(false))
	SceneCamera.AddBlitCamera(PostCamera)

	// Shaders
	TestLESL := gfx.CreateLESLPlugin(gio.LoadAsString("assets/testshader.lesl"))
	FBOLESL := gfx.CreateLESLPlugin(gio.LoadAsString("assets/fboshader.lesl"))

	SceneShader = utils2d.NewSpriteShader(TestLESL)
	TextShader = gfx.NewTextShader()
	PostShader = gfx.NewGenericShader(gfx.NewRenderProgram(FBOLESL))

	// Level
	LevelMaterial = gfx.NewColorMaterial(gmath.NewVector3(0.4, 0.4, 0.45))

	// Player
	PlayerTexture = gfx.NewTexture(gio.LoadPNG("assets/slime.png"))
	PlayerTexture.SetPointFilter(true, false)

	PlayerMaterial = gfx.NewTextureMaterial(PlayerTexture)
	PlayerMaterial.SetTint(gmath.NewVector3(0.5, 0.5, 0.0), 0.5)

	PlayerSpriteSheet = gfx.NewSpriteSheet(0.25, 0.25, 0.003)

	PlayerRightIdle = gfx.NewFrameAnimation(gfx.NewDurationFrame(PlayerSpriteSheet.GetBounds(0), 1.25), gfx.NewDurationFrame(PlayerSpriteSheet.GetBounds(3), 1.25))
	PlayerRightWalk = gfx.NewFrameAnimation(gfx.NewDurationFrame(PlayerSpriteSheet.GetBounds(0), 0.25), gfx.NewDurationFrame(PlayerSpriteSheet.GetBounds(1), 0.25), gfx.NewDurationFrame(PlayerSpriteSheet.GetBounds(2), 0.25), gfx.NewDurationFrame(PlayerSpriteSheet.GetBounds(3), 0.25))
	PlayerRightWallSlide = gfx.NewFrameAnimation(gfx.NewDurationFrame(PlayerSpriteSheet.GetBounds(8), 0.5), gfx.NewDurationFrame(PlayerSpriteSheet.GetBounds(10), 0.5))

	PlayerLeftIdle = gfx.NewFrameAnimation(gfx.NewDurationFrame(PlayerSpriteSheet.GetBounds(4), 1.25), gfx.NewDurationFrame(PlayerSpriteSheet.GetBounds(7), 1.25))
	PlayerLeftWalk = gfx.NewFrameAnimation(gfx.NewDurationFrame(PlayerSpriteSheet.GetBounds(4), 0.25), gfx.NewDurationFrame(PlayerSpriteSheet.GetBounds(5), 0.25), gfx.NewDurationFrame(PlayerSpriteSheet.GetBounds(6), 0.25), gfx.NewDurationFrame(PlayerSpriteSheet.GetBounds(7), 0.25))
	PlayerLeftWallSlide = gfx.NewFrameAnimation(gfx.NewDurationFrame(PlayerSpriteSheet.GetBounds(12), 0.5), gfx.NewDurationFrame(PlayerSpriteSheet.GetBounds(14), 0.5))

	// Font
	SegoeFont = gio.LoadFNT("assets", "segoe.fnt")
	SegoeTexture = gfx.NewTexture(gio.LoadPNG("assets/segoe.png"))
	SegoeTexture.SetLinearFilter(true, false)
	SegoeMaterial = gfx.NewTextureMaterial(SegoeTexture)
	SegoeMaterial.SetTint(gmath.NewVector3(0.0, 0.0, 0.0), 1.0)
}

package utils2d

type SpriteSheetAnimationComponent struct {
}

func (spriteSheetAnimationComponent *SpriteSheetAnimationComponent) Delete() {}

// type SpriteSheetAnimation struct {
// 	spriteSheet           *SpriteSheet
// 	animations            map[string][]uint32
// 	currentAnimation      string
// 	currentAnimationIndex int
// }

// func CreateSpriteSheetAnimation(spriteSheet *SpriteSheet) *SpriteSheetAnimation {
// 	material := &SpriteSheetAnimation{
// 		spriteSheet:           spriteSheet,
// 		animations:            make(map[string][]uint32),
// 		currentAnimation:      "idleRight",
// 		currentAnimationIndex: 0,
// 	}
// var anim func()
// anim = func() {
// 	material.currentAnimationIndex++
// 	if material.currentAnimationIndex >= len(material.animations[material.currentAnimation]) {
// 		material.currentAnimationIndex = 0
// 	}
// 	if animation, ok := material.animations[material.currentAnimation]; ok {
// 		material.spriteSheet.SetIndex(animation[material.currentAnimationIndex])
// 	}
// 	fmt.Println(material.currentAnimationIndex)
// 	limitengine.DelayFunc(anim, 0.25)
// }
// limitengine.DelayFunc(anim, 0.25)
// 	return material
// }

// func (material *SpriteSheetAnimation) AddAnimationSequence(name string, frames ...uint32) {
// 	material.animations[name] = frames
// }

// func (material *SpriteSheetAnimation) SetAnimation(name string) {
// 	if material.currentAnimation != name {
// 		material.currentAnimationIndex = 0
// 		material.currentAnimation = name
// 	}
// }

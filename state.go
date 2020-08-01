package limitengine

type State interface {
	OnActive()
	Update(delta float32)
	OnInactive()
}

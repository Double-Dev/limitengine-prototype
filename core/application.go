package core

var _ int64 = run()

var (
	Running bool
)

func run() int64 {
	Running = true
	return 1
}

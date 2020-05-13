# limitengine
A custom GoLang cross-platform game engine.

# Dependencies
This engine uses the GLFW, OpenGL, OpenAL, and Vorbis C libraries.

External Packages:
- golang.org/x/mobile
- github.com/vulkan-go/glfw/v3.3/glfw
- github.com/go-gl/gl/v3.3-core/gl

# Profiling
Put the following line of code at the top of your main method:
`defer profile.Start().Stop()`

Note the file path now output at the end of every run.

Run the following command to enter the profiling console after running the engine:
`go tool pprof <insert path of output file here>`

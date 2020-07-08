# limitengine
A custom GoLang cross-platform *application* engine. This engine has been designed with a wide variety of uses in mind (not just game development) to provide both the interface and tools necessary for developers to write any kind of program.
Please note that Limitengine is still in alpha development and some features are subject to change. 

## Requirements
- A cgo compiler.
- On Ubuntu/Debian-based systems, the `libgl1-mesa-dev` package.

## Dependencies
This engine contains bindings for the GLFW, OpenGL, OpenAL Soft, and Vorbis libraries for low-level graphical and audio interface.

Please note that the following external Go packages are used for mobile and web builds:
- [golang.org/x/mobile]
- [github.com/gopherjs/gopherjs]

## Profiling
Put the following line of code at the top of your main method:
`defer profile.Start().Stop()`

Note the file path now output at the end of every run.

Run the following command to enter the profiling console after running the engine:
`go tool pprof <insert path of output file here>`

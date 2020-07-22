# LimitEngine Prototype
NOTICE: This repository is a prototype for LimitEngine written in GoLang.
LimitEngine's will be written in C++ and can be found in my LimitEngine repository. 

A custom GoLang cross-platform *application* engine. This engine has been designed with a wide variety of uses in mind (not just game development) to provide both the interface and tools necessary for developers to write any kind of program.
Please note that Limitengine is still in alpha development and some features are subject to change.

## Requirements
- A cgo compiler.
- On Ubuntu/Debian-based systems, the `libgl1-mesa-dev` package.

## Dependencies
This engine contains bindings for the GLFW, OpenGL, OpenAL Soft, and Vorbis libraries for low-level graphical and audio interface.

## Profiling
Put the following line of code at the top of your main method:
`defer profile.Start().Stop()`

Note the file path now output at the end of every run.

Run the following command to enter the profiling console after running the engine:
`go tool pprof <insert path of output file here>`

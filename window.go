package rtech

import (
	"log"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type RWindow struct {
	window *glfw.Window
}

type Window interface {
	Present()
	Destroy()
	ShouldClose() bool
}

func NewWindow() *RWindow {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	window, err := glfw.CreateWindow(1280, 720, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		log.Fatalln(err)
	}
	return &RWindow{
		window: window,
	}
}
func (w *RWindow) ShouldClose() bool {
	return w.window.ShouldClose()
}
func (w *RWindow) Present() {
	w.window.SwapBuffers()
	glfw.PollEvents()
}

func (w *RWindow) Destroy() {
	glfw.Terminate()
}

package rtech

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	mgl "github.com/go-gl/mathgl/mgl64"
)

type RWindow struct {
	window     *glfw.Window
	width      int
	height     int
	x          int
	y          int
	fullscreen bool
}

type Window interface {
	Present()
	Destroy()
	Position() mgl.Vec2
	Size() mgl.Vec2
	ToggleFullscreen() bool
	ShouldClose() bool
}

func NewWindow() *RWindow {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Hint(glfw.ContextVersionMajor), 4)
	glfw.WindowHint(glfw.Hint(glfw.ContextVersionMinor), 1)
	glfw.WindowHint(glfw.Hint(glfw.OpenGLForwardCompatible), 1)
	glfw.WindowHint(glfw.Hint(glfw.OpenGLProfile), glfw.OpenGLCoreProfile)
	if runtime.GOOS == "darwin" {
		glfw.WindowHint(glfw.Hint(glfw.CocoaRetinaFramebuffer), 1)
	}
	window, err := glfw.CreateWindow(1280, 720, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		log.Fatalln(err)
	}
	rwindow := &RWindow{
		window:     window,
		x:          0,
		y:          0,
		width:      1280,
		height:     720,
		fullscreen: false,
	}

	window.SetSizeCallback(rwindow.resize)
	window.SetPosCallback(rwindow.move)
	return rwindow
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

func (w *RWindow) resize(win *glfw.Window, width, height int) {
	w.width = width
	w.height = height
}
func (w *RWindow) move(win *glfw.Window, x, y int) {
	w.x = x
	w.y = y
}

func (w *RWindow) Position() mgl.Vec2 {
	return mgl.Vec2{float64(w.x), float64(w.y)}
}
func (w *RWindow) Size() mgl.Vec2 {
	return mgl.Vec2{float64(w.width), float64(w.height)}
}
func (w *RWindow) ToggleFullscreen() bool {
	monitor := glfw.GetPrimaryMonitor()
	if !w.fullscreen {
		w.fullscreen = true
		w.window.SetMonitor(monitor, 0, 0, monitor.GetVideoMode().Width, monitor.GetVideoMode().Height, monitor.GetVideoMode().RefreshRate)
	} else {
		w.fullscreen = false
		w.window.SetMonitor(nil, w.x, w.y, w.width, w.height, monitor.GetVideoMode().RefreshRate)
	}
	return w.fullscreen
}

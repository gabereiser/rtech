package rtech

import (
	"runtime"
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type REngine struct {
	window   Window
	running  bool
	updatefn func(time time.Duration)
	renderfn func(time time.Duration)
}

var __engine *REngine

func EngineInit(updatefn, renderfn func(time time.Duration)) *REngine {
	if __engine == nil {
		runtime.LockOSThread()
		__engine = &REngine{
			window:   NewWindow(),
			running:  true,
			updatefn: updatefn,
			renderfn: renderfn,
		}
	}
	return __engine
}

var t time.Time
var fps int64

func (e *REngine) Run() error {
	for {
		if !e.running {
			break
		}
		gameTime := time.Since(t)
		t = time.Now()
		e.update(gameTime)
		e.render(gameTime)
		if gameTime.Milliseconds() > int64(0) {
			fps = (1000 / gameTime.Milliseconds())
		}

	}
	return nil
}
func (e *REngine) clear() {
	gl.Scissor(0, 0, 1280, 720)
	gl.ClearColor(0.2, 0.2, 0.2, 1.0)
	gl.Clear(gl.DEPTH_BUFFER_BIT | gl.COLOR_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)
}
func (e *REngine) render(time time.Duration) {
	// main render pass.
	e.clear()
	e.renderfn(time)
	e.window.Present()
}

func (e *REngine) update(time time.Duration) {
	if e.window.ShouldClose() {
		e.running = false
		return
	}
	e.updatefn(time)
}

func (e *REngine) Destroy() {
	e.window.Destroy()
}

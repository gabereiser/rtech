package rtech

import (
	"fmt"
	"runtime"
	"time"

	fs "github.com/gabereiser/rtech/fs"
	"github.com/gabereiser/rtech/gl"
	"github.com/gabereiser/rtech/types"
)

// REngine is the main object that handles the core services of the game engine.
// Window init, calling update and render, shutting down.
// Call [EngineInit] to get an REngine instance to begin the game engine.
type REngine struct {
	fs         fs.FileSystem
	window     gl.Window
	running    bool
	game       RGame
	clearColor types.RColor
	viewport   types.RViewport
}

var __engine *REngine

// EngineInit
// Pass an RGame interface and get back an initialized REngine ready to run.
func EngineInit(game RGame) *REngine {
	if __engine == nil {
		runtime.LockOSThread()
		window := gl.NewWindow()
		__engine = &REngine{
			fs:         fs.NewFilesystem(),
			window:     window,
			running:    false,
			game:       game,
			clearColor: types.RColor{255, 255, 255, 255},
			viewport:   types.RViewport{0, 0, int(window.Size().X()), int(window.Size().Y())},
		}
	}
	return __engine
}

var t time.Time
var fps int64

// Run
// runs the REngine main loop. If you have a pointer to an REngine, it's safe to call.
func (e *REngine) Run() error {
	if e.running {
		return fmt.Errorf("engine is already running")
	}
	e.running = true
	for {
		if !e.running {
			break
		}
		gameTime := time.Since(t)
		t = time.Now()
		e.updateFps(gameTime)
		e.update(gameTime)
		e.render(gameTime)
	}
	return nil
}

// updateFps will update the current tracked fps based on duration of last frame.
func (e *REngine) updateFps(gameTime time.Duration) {
	if gameTime.Milliseconds() > int64(0) {
		fps = (1000 / gameTime.Milliseconds())
	} else {
		fps = 0
	}
}

func (e *REngine) Clear() {
	gl.ClearAll(e.clearColor)
}

func (e *REngine) render(time time.Duration) {
	// main render pass.
	e.viewport.Bind()
	e.Clear()
	e.game.Render(time)
	e.window.Present()
}

func (e *REngine) update(time time.Duration) {
	if e.window.ShouldClose() {
		e.running = false
		return
	}
	e.game.Update(time)
}

func (e *REngine) Destroy() {
	e.window.Destroy()
}

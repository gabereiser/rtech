package rtech

import (
	"fmt"
	"runtime"
	"time"

	fs "github.com/gabereiser/rtech/fs"
	"github.com/gabereiser/rtech/gl"
)

// REngine is the main object that handles the core services of the game engine.
// Window init, calling update and render, shutting down.
// Call [EngineInit] to get an REngine instance to begin the game engine.
type REngine struct {
	fs         fs.FileSystem
	window     gl.Window
	running    bool
	game       RGame
	clearColor RColor
	viewport   RViewport
	scene      RScene
}

var __engine *REngine

// EngineInit
// Pass an RGame interface and get back an initialized REngine ready to run.
func EngineInit(game RGame) *REngine {
	if __engine == nil {
		runtime.LockOSThread()
		window := gl.NewWindow()
		scene := RScene{}
		root := scene.CreateNode()
		scene.SetRootNode(root)
		__engine = &REngine{
			fs:         fs.NewFilesystem(),
			window:     window,
			running:    false,
			game:       game,
			clearColor: RColor{255, 255, 255, 255},
			viewport:   RViewport{0, 0, int(window.Size().X()), int(window.Size().Y())},
			scene:      scene,
		}
	}
	return __engine
}

var t time.Time
var fps int64

// updateFps will update the current tracked fps based on duration of last frame.
func (e *REngine) updateFps(gameTime time.Duration) {
	if gameTime.Milliseconds() > int64(0) {
		fps = (1000 / gameTime.Milliseconds())
	} else {
		fps = 0
	}
}

// render will bind the viewport, call clear to clear the screen, call [RGame.Render] on your game instance so you can draw, then present the results to the screen by flipping the swap chain.
func (e *REngine) render(time time.Duration) {
	// main render pass.
	e.viewport.Bind()
	e.Clear()
	e.game.Render(time)
	e.window.Present()
}

// update will check if the engine should shutdown then it passes to [RGame.Update] so you can perform update logic to your game.
func (e *REngine) update(time time.Duration) {
	if e.window.ShouldClose() {
		e.running = false
		return
	}
	e.game.Update(time)
}

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

// Clear will clear the current Color, Depth, and Stencil Buffers.
func (e *REngine) Clear() {
	gl.ClearAll(e.clearColor.RedF(), e.clearColor.GreenF(), e.clearColor.BlueF(), e.clearColor.AlphaF())
}

// ClearBit will allow you to clear a specific buffer using [gl.COLOR_BUFFER_BIT] [gl.DEPTH_BUFFER_BIT] or [gl.STENCIL_BUFFER_BIT]
func (e *REngine) ClearBit(bit int) {
	gl.Clear(e.clearColor.RedF(), e.clearColor.GreenF(), e.clearColor.BlueF(), e.clearColor.AlphaF(), bit)
}

// Destroy will destroy the game window. You still need to call Destroy on your owned objects.
func (e *REngine) Destroy() {
	e.window.Destroy()
}

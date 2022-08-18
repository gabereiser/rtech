package input

import "github.com/go-gl/glfw/v3.3/glfw"

func charCallback(w *glfw.Window, char rune) {

}
func kcb(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if __keycallback != nil {
		__keycallback(int(key), scancode, int(action), int(mods))
	}

}
func mpcb(w *glfw.Window, x float64, y float64) {
	if __mousecallback != nil {
		__mousecallback(x, y)
	}
}
func mbcb(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if __mousebuttoncallback != nil {
		__mousebuttoncallback(int(button), int(action), int(mods))
	}
}

type KeyCallBack func(key int, scancode int, action int, mods int)
type MousePositionCallBack func(x float64, y float64)
type MouseButtonCallBack func(button int, action int, mods int)

var __keycallback KeyCallBack
var __mousecallback MousePositionCallBack
var __mousebuttoncallback MouseButtonCallBack

func Init() {
	window := glfw.GetCurrentContext()
	window.SetInputMode(glfw.LockKeyMods, glfw.True)
	window.SetInputMode(glfw.StickyKeysMode, glfw.True)
	window.SetInputMode(glfw.StickyMouseButtonsMode, glfw.True)
	window.SetCharCallback(charCallback)
	window.SetCursorPosCallback(mpcb)
	window.SetMouseButtonCallback(mbcb)
	window.SetKeyCallback(kcb)
}

func SetKeyCallback(kcb KeyCallBack) {
	__keycallback = kcb
}

func SetMousePositionCallback(mcb MousePositionCallBack) {
	__mousecallback = mcb
}

func PollEvents() {
	glfw.PollEvents()
}

func IsKeyDown(key int) bool {
	action := glfw.GetCurrentContext().GetKey(glfw.Key(key))
	return action == glfw.Press
}
func IsKeyUp(key int) bool {
	action := glfw.GetCurrentContext().GetKey(glfw.Key(key))
	return action == glfw.Release
}
func IsMouseDown(button int) bool {
	action := glfw.GetCurrentContext().GetMouseButton(glfw.MouseButton(button))
	return action == glfw.Press
}
func IsMouseUp(button int) bool {
	action := glfw.GetCurrentContext().GetMouseButton(glfw.MouseButton(button))
	return action == glfw.Release
}

var __cursorLock = false

func LockCursor() {
	if __cursorLock {
		return
	}
	window := glfw.GetCurrentContext()
	if glfw.RawMouseMotionSupported() {
		window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
		window.SetInputMode(glfw.RawMouseMotion, glfw.True)
		__cursorLock = true
	} else {
		window.SetInputMode(glfw.CursorMode, glfw.CursorHidden)
		__cursorLock = true
	}
}
func UnlockCursor() {
	if !__cursorLock {
		return
	}
	window := glfw.GetCurrentContext()
	if glfw.RawMouseMotionSupported() {
		window.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
		window.SetInputMode(glfw.RawMouseMotion, glfw.False)
		__cursorLock = false
	} else {
		window.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
		__cursorLock = false
	}
}
func ShowCursor() {
	window := glfw.GetCurrentContext()
	window.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
}
func HideCursor() {
	window := glfw.GetCurrentContext()
	window.SetInputMode(glfw.CursorMode, glfw.CursorHidden)
}

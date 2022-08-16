package rtech

import "time"

type RGame interface {
	Init()
	Render(gameTime time.Duration)
	Update(gameTime time.Duration)
}

package main

import (
	"time"

	"github.com/gabereiser/rtech"
)

var engine *rtech.REngine

type Game struct{}

func main() {
	game := &Game{}
	engine = rtech.EngineInit(game)
	engine.Run()
	engine.Destroy()
}

func (g *Game) Init() {

}

func (g *Game) Update(time time.Duration) {

}

func (g *Game) Render(time time.Duration) {

}

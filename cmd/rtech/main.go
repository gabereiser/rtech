package main

import (
	"time"

	"github.com/gabereiser/rtech"
)

var engine *rtech.REngine

func main() {
	engine = rtech.EngineInit(update, render)
	engine.Run()
	engine.Destroy()
}

func update(time time.Duration) {

}

func render(time time.Duration) {

}

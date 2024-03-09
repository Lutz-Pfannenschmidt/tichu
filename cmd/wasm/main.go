//go:build js && wasm

package main

import (
	"syscall/js"
	"time"

	"github.com/Lutz-Pfannenschmidt/tichu/internal/tichu"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/markfarnan/go-canvas/canvas"
)

var cvs *canvas.Canvas2d
var width float64
var height float64

var deck = tichu.NewDeck()

var renderDelay time.Duration = 20 * time.Millisecond

func main() {

	FrameRate := time.Second / renderDelay
	println("FPS: ", FrameRate)

	cvs, _ = canvas.NewCanvas2d(false)
	cvs.Create(int(js.Global().Get("innerWidth").Float()), int(js.Global().Get("innerHeight").Float()))

	height = float64(cvs.Height())
	width = float64(cvs.Width())

	cvs.Start(60, Render)

	<-make(chan struct{})
}

func Render(gc *draw2dimg.GraphicContext) bool {

	for i, card := range deck.Cards {
		card.Render(gc, float64(i*30), 0)
	}

	return true
}

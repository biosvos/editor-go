package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"log"
)

func main() {
	go func() {
		err := Run()
		if err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func Run() error {
	var ops op.Ops
	window := app.NewWindow()
	th := material.NewTheme(gofont.Collection())
	var editor widget.Editor

	for ev := range window.Events() {
		switch ev := ev.(type) {
		case system.DestroyEvent:
			return ev.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, ev)
			mEditor := material.Editor(th, &editor, "힌트다")
			mEditor.Layout(gtx)
			ev.Frame(gtx.Ops)
		default:
			log.Println(ev)
		}
	}
	return nil
}

package main

import (
	"fmt"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		// create new window
		w := app.NewWindow(
			app.Title("MySQL Spider"),
		)

		var startButton widget.Clickable

		var fileButton widget.Clickable

		// var progressIncrementer chan float32
		var progress float32

		var ops op.Ops

		th := material.NewTheme(gofont.Collection())

		// listen for events in the window.
		for e := range w.Events() {
			// detect what type of event
			switch e := e.(type) {

			// this is sent when the application should re-render.
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				layout.Flex{
					// Vertical alignment, from top to bottom
					Axis: layout.Vertical,
					// Empty space is left at the start, i.e. at the top
				}.Layout(gtx,
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							label := material.H1(th, "MySQL Spider")
							return label.Layout(gtx)
						},
					),
				)

				layout.Flex{
					// Vertical alignment, from top to bottom
					Axis: layout.Vertical,
					// Empty space is left at the start, i.e. at the top
					Spacing: layout.SpaceStart,
				}.Layout(gtx,
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							bar := material.ProgressBar(th, progress) // Here progress is used
							return bar.Layout(gtx)
						},
					),

					layout.Rigid(
						// The height of the spacer is 25 Device independent pixels
						layout.Spacer{Height: unit.Dp(25)}.Layout,
					),
					// First a button ...
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							btn := material.Button(th, &startButton, "Start")
							return btn.Layout(gtx)
						},
					),
					layout.Rigid(
						// The height of the spacer is 25 Device independent pixels
						layout.Spacer{Height: unit.Dp(5)}.Layout,
					),
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							btn := material.Button(th, &fileButton, "Open File (Not working yet")
							return btn.Layout(gtx)
						},
					),
					layout.Rigid(
						// The height of the spacer is 25 Device independent pixels
						layout.Spacer{Height: unit.Dp(25)}.Layout,
					),
				)

				e.Frame(gtx.Ops)
				if startButton.Clicked() {
					fmt.Println("tsest")
					progress += 0.04
					w.Invalidate()
				}

			}

		}
	}()
	app.Main()
}

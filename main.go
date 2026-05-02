package main

import (
	"os"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main() {
	app := gtk.NewApplication("com.example.myapp", 0)

	app.ConnectActivate(func() {
		win := gtk.NewApplicationWindow(app)
		win.SetTitle("Testing")
		win.SetDefaultSize(800, 600)

		vbox := gtk.NewBox(gtk.OrientationVertical, 8)
		win.SetChild(vbox)

		topBar := gtk.NewBox(gtk.OrientationHorizontal, 8)
		topBar.SetHExpand(true)

		leftGroup := gtk.NewBox(gtk.OrientationHorizontal, 4)
		leftGroup.Append(gtk.NewButtonWithLabel("--1--"))
		leftGroup.Append(gtk.NewButtonWithLabel("--2--"))

		rightGroup := gtk.NewBox(gtk.OrientationHorizontal, 4)
		rightGroup.Append(gtk.NewButtonWithLabel("--3--"))
		rightGroup.Append(gtk.NewButtonWithLabel("--4--"))

		spacer := gtk.NewBox(gtk.OrientationHorizontal, 0)
		spacer.SetHExpand(true)

		topBar.Append(leftGroup)
		topBar.Append(spacer)
		topBar.Append(rightGroup)

		vbox.Append(topBar)

		win.Show()
	})

	os.Exit(app.Run(os.Args))
}

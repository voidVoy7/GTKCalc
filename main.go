package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func calculate(expr string) string {
	for _, op := range []string{"+", "-", "*", "/"} {
		idx := strings.LastIndex(expr, op)
		if idx <= 0 {
			continue
		}
		a, err1 := strconv.ParseFloat(expr[:idx], 64)
		b, err2 := strconv.ParseFloat(expr[idx+1:], 64)
		if err1 != nil || err2 != nil {
			continue
		}
		switch op {
		case "+":
			return fmt.Sprintf("%g", a+b)
		case "-":
			return fmt.Sprintf("%g", a-b)
		case "*":
			return fmt.Sprintf("%g", a*b)
		case "/":
			if b == 0 {
				return "DO NOT DIVIDE BY 0!!"
			}
			return fmt.Sprintf("%g", a/b)
		}
	}
	return expr
}

func main() {
	app := gtk.NewApplication("com.example.myapp", 0)

	app.ConnectActivate(func() {
		win := gtk.NewApplicationWindow(app)
		win.SetTitle("Testing")
		// win.SetDefaultSize(400, 400) - Not needed unless dont want automatic scaling.

		vbox := gtk.NewBox(gtk.OrientationVertical, 0)
		win.SetChild(vbox)

		top := gtk.NewBox(gtk.OrientationHorizontal, 8)
		top.SetHExpand(true)

		display := gtk.NewEntry()
		display.SetEditable(false)
		display.SetHExpand(true)
		vbox.Append(display)

		numbers := gtk.NewGrid()
		numbers.SetColumnSpacing(4)
		numbers.SetRowSpacing(4)

		labels := []string{
			"1", "2", "3", "/",
			"4", "5", "6", "*",
			"7", "8", "9", "-",
			"0", "C", "=", "+",
		}
		cols := 4

		for i, label := range labels {
			btn := gtk.NewButtonWithLabel(label)
			lbl := label
			btn.ConnectClicked(func() {
				current := display.Text()
				switch lbl {
				case "=":
					display.SetText(calculate(current))
				case "C":
					display.SetText("")
				default:
					display.SetText(current + lbl)
				}
			})
			numbers.Attach(btn, i%cols, i/cols, 1, 1)
		}

		top.Append(numbers)
		vbox.Append(top)
		win.Show()

	})

	os.Exit(app.Run(os.Args))
}

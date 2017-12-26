package tree

import (
	"math/rand"

	"github.com/fatih/color"
)

var (
	yellowColor = color.New(color.FgYellow)
	redColor    = color.New(color.FgRed)

	colors = []*color.Color{
		color.New(color.FgGreen),
		yellowColor,
		redColor,
		color.New(color.FgBlue),
		color.New(color.FgCyan),
		color.New(color.FgMagenta),
	}

	colorSize = len(colors)
)

func randomColor() *color.Color {
	return colors[rand.Intn(colorSize)]
}

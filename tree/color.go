package tree

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var (
	colors = []*color.Color{
		color.New(color.FgGreen),
		color.New(color.FgYellow),
		color.New(color.FgRed),
		color.New(color.FgBlue),
		color.New(color.FgCyan),
		color.New(color.FgMagenta),
	}

	colorSize = len(colors)
)

func randomColor() *color.Color {
	rand.Seed(time.Now().UnixNano())
	return colors[rand.Intn(colorSize)]
}

func yellowColor() *color.Color {
	return colors[1]
}

func redColor() *color.Color {
	return colors[2]
}

package main

import (
	"github.com/fogleman/gg"
	"math"
)

var COLORS = []string {
	"69D2E7",
	"A7DBD8",
	"E0E4CC",
	"F38630",
	"FA6900",
}

func drawTriangle(
	ctx *gg.Context,
	x float64,
	y float64,
	radius float64,
	rotation float64,
	color string,
) {
	ctx.Push()
	ctx.SetHexColor(color)
	// Translate to centre
	ctx.Translate(x, y)
	ctx.Rotate(rotation)
	// Move to the radius up
	ctx.MoveTo(0, -radius)
	// For each point:
	for i := 0; i < 3; i++ {
		// Rotate 2/3 rad
		ctx.Rotate(gg.Radians(360.0 / 3.0))
		// Line to radius up
		ctx.LineTo(0, -radius)
	}
	ctx.Fill()
	ctx.Pop()
}

func triBaseWidth(radius float64) float64 {
	return math.Sin(gg.Radians(60.0)) * radius
}

func triHeight(radius float64) float64 {
	return (math.Sin(gg.Radians(30.0)) * radius) + radius
}

func main() {
	ctx := gg.NewContext(500, 500)
	r := 40.0
	width := triBaseWidth(r)
	height := triHeight(r)

	colorCount := 0

	for col := 0; col < 100; col++ {
		for row := 0; row < 100; row++ {
			var (
				color string
				rot float64
				x float64
				y float64
			)

			if row % 3 == 0 {
				colorCount = (colorCount + 1) % len(COLORS)
			}
			color = COLORS[colorCount]

			x = float64(row) * width
			y = float64(col) * height
			if col % 2 == 0 {
				x -= width
			}
			if row % 2 == 0 {
				rot = 0
				y += r
			} else {
				rot = 3.141
				y += r - (r / 2.0)
			}
			drawTriangle(ctx, x, y, r, rot, color)
		}
	}

	ctx.SavePNG("out.png")
}

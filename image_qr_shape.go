package qrcode

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
)

var (
	_shapeRectangle IShape = rectangle{}
	_shapeCircle    IShape = circle{}
)

type IShape interface {
	// draw to fill the IShape
	Draw(ctx *DrawContext)
}

// DrawContext is a rectangle area
type DrawContext struct {
	*gg.Context

	UpperLeft image.Point // (x1, y1)
	W, H      int

	Color color.Color
}

// rectangle IShape
type rectangle struct{}

func (r rectangle) Draw(c *DrawContext) {
	// FIXED(@yeqown): miss parameter of DrawRectangle
	c.DrawRectangle(float64(c.UpperLeft.X), float64(c.UpperLeft.Y),
		float64(c.W), float64(c.H))
	c.SetColor(c.Color)
	c.Fill()
}

// circle IShape
type circle struct{}

// FIXED: Draw could not draw circle
func (r circle) Draw(c *DrawContext) {
	// choose a proper radius values
	radius := c.W / 2
	r2 := c.H / 2
	if r2 <= radius {
		radius = r2
	}

	cx, cy := c.UpperLeft.X+c.W/2, c.UpperLeft.Y+c.H/2 // get center point
	c.DrawCircle(float64(cx), float64(cy), float64(radius))
	c.SetColor(c.Color)
	c.Fill()
}

package qrcode

import (
	"image"
	"image/color"
	"testing"

	"github.com/fogleman/gg"
)

func Test_rectangle_Draw(t *testing.T) {
	rect := image.Rect(0, 0, 100, 100)
	rgba := image.NewRGBA(rect)
	dc := gg.NewContextForRGBA(rgba)

	dc.DrawRectangle(0, 0, 100, 100)
	dc.SetColor(color.White)
	dc.Fill()

	ctx := &DrawContext{
		Context:   dc,
		UpperLeft: image.Point{X: 0, Y: 0},
		W:         50,
		H:         50,
		Color:     color.Black,
	}
	_shapeRectangle.Draw(ctx)

	err := dc.SavePNG("./testdata/rectangle.png")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func Test_circle_Draw(t *testing.T) {
	rect := image.Rect(0, 0, 100, 100)
	rgba := image.NewRGBA(rect)
	dc := gg.NewContextForRGBA(rgba)

	dc.DrawRectangle(0, 0, 100, 100)
	dc.SetColor(color.White)
	dc.Fill()

	ctx := &DrawContext{
		Context:   dc,
		UpperLeft: image.Point{X: 0, Y: 0},
		W:         50,
		H:         50,
		Color:     color.Black,
	}
	_shapeCircle.Draw(ctx)

	err := dc.SavePNG("./testdata/circle.png")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func Test_gg(t *testing.T) {
	rect := image.Rect(0, 0, 100, 100)
	rgba := image.NewRGBA(rect)
	dc := gg.NewContextForRGBA(rgba)

	dc.DrawRectangle(0, 0, 100, 100)
	dc.SetColor(color.White)
	dc.Fill()

	dc.DrawCircle(50, 50, 40)
	dc.SetColor(color.Black)
	dc.Fill()
	_ = dc.SavePNG("out.png")
}

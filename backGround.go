package text2picture

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
)

func NewColorBackGround(width, height int, color color.Color) *image.RGBA {
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(rgba, rgba.Bounds(), image.NewUniform(color), image.Point{}, draw.Src)
	return rgba
}

func LoadBackGround(file io.Reader) *image.RGBA {
	bac, err := png.Decode(file)
	if err != nil {
		return nil
	}
	rgba := image.NewRGBA(image.Rect(0, 0, bac.Bounds().Max.X, bac.Bounds().Max.Y))
	draw.Src.Draw(rgba, rgba.Rect, bac, bac.Bounds().Min)
	return rgba
}

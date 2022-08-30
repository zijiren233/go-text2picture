package text2picture

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func NewColorBackGround(width, height int, color color.Color) *image.RGBA {
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(rgba, rgba.Bounds(), image.NewUniform(color), image.Point{}, draw.Src)
	return rgba
}

func LoadBackGround(filePath string) *image.RGBA {
	f, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	bac, err := png.Decode(f)
	if err != nil {
		return nil
	}
	return bac.(*image.RGBA)
}

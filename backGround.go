package text2picture

import (
	"image"
	"image/color"
	"image/draw"
)

func NewBackGroundWithColor(width, height int, _16bit_color uint16) *image.RGBA {
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(rgba, rgba.Bounds(), image.NewUniform(color.Gray16{_16bit_color}), image.Point{}, draw.Src)
	return rgba
}

func NewWhiteBackGround(width, height int) *image.RGBA {
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(rgba, rgba.Bounds(), image.White, image.Point{}, draw.Src)
	return rgba
}

func NewBlackBackGround(width, height int) *image.RGBA {
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(rgba, rgba.Bounds(), image.Black, image.Point{}, draw.Src)
	return rgba
}

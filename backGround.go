package text2picture

import (
	"image"
	"image/draw"
)

func NewBackGroundWithColor(width, height int, color *image.Uniform) *image.RGBA {
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(rgba, rgba.Bounds(), color, image.Point{}, draw.Src)
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

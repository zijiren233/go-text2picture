package text2picture

import (
	"image"
	"image/color"
)

var (
	Transparent = image.Transparent
	White       = image.White
	Black       = image.Black
	Red         = image.NewUniform(color.RGBA{255, 0, 0, 255})
	Green       = image.NewUniform(color.RGBA{0, 255, 0, 255})
	Blue        = image.NewUniform(color.RGBA{0, 0, 255, 255})
)

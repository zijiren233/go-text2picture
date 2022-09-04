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

func NewColor(R, G, B, A uint8) *image.Uniform {
	return RGBA(R, G, B, A)
}

func RGBA(R, G, B, A uint8) *image.Uniform {
	return image.NewUniform(color.RGBA{R, G, B, A})
}

package text2picture

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"os"
)

// Color can be generated using text2picture.NewColor()
func NewColorPicture(width, height int, color color.Color) *image.RGBA {
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(rgba, rgba.Rect, image.NewUniform(color), image.Point{}, draw.Src)
	return rgba
}

func LoadPicture(filepath string) (*image.RGBA, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	bac, err := png.Decode(f)
	if err != nil {
		return nil, err
	}
	rgba := image.NewRGBA(image.Rect(0, 0, bac.Bounds().Max.X, bac.Bounds().Max.Y))
	draw.Src.Draw(rgba, rgba.Rect, bac, bac.Bounds().Min)
	return rgba, nil
}

func ReadPicture(file io.Reader) (*image.RGBA, error) {
	bac, err := png.Decode(file)
	if err != nil {
		return nil, err
	}
	rgba := image.NewRGBA(image.Rect(0, 0, bac.Bounds().Max.X, bac.Bounds().Max.Y))
	draw.Src.Draw(rgba, rgba.Rect, bac, bac.Bounds().Min)
	return rgba, nil
}

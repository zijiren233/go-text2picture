package text2picture

import (
	"bufio"
	"bytes"
	"image"
	"image/png"
	"io"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type picture struct {
	dpi float64

	font *truetype.Font

	// font size in points
	fontSize float64

	// text left and right padding
	padding int

	// line spacing
	spacing float64

	rgba *image.RGBA

	pt fixed.Point26_6
	c  *freetype.Context
}

func (p *picture) GetPoint() (fixed.Int26_6, fixed.Int26_6) {
	return p.pt.X, p.pt.Y
}

func (p *picture) SetPoint(x, y float64) {
	p.pt.X = p.c.PointToFixed(x)
	p.pt.Y = p.c.PointToFixed(y)
}

func (p *picture) PointOffset(x, y float64) {
	p.pt.X += p.c.PointToFixed(x)
	p.pt.Y += p.c.PointToFixed(y)
}

func (p *picture) SetSpacing(spacing float64) {
	p.spacing = spacing
}

func (p *picture) SetFontSize(fontSize float64) {
	p.c.SetFontSize(fontSize)
	p.fontSize = fontSize
}

func (p *picture) SetPadding(padding int) {
	p.padding = padding
}

func (p *picture) SetFont(font *truetype.Font) {
	p.c.SetFont(font)
	p.font = font
}

func (p *picture) GetRGBA() *image.RGBA {
	return p.rgba
}

// padding: text left and right padding
func NewPictureWithBackGroundFile(file io.Reader, dpi float64, padding int, fontSize float64) *picture {
	p := picture{dpi: dpi, padding: padding, fontSize: fontSize, spacing: 1, font: DefaultFont}
	bac, err := png.Decode(file)
	if err != nil {
		return nil
	}

	p.rgba = bac.(*image.RGBA)
	p.c = freetype.NewContext()
	p.c.SetDPI(p.dpi)
	p.c.SetClip(p.rgba.Bounds())
	p.c.SetDst(p.rgba)
	p.c.SetHinting(font.HintingFull)
	p.c.SetFont(p.font)
	p.c.SetFontSize(p.fontSize)

	p.pt = freetype.Pt(padding, p.c.PointToFixed(fontSize).Round())
	return &p
}

func NewPictureWithBackGround(png *image.RGBA, dpi float64, padding int, fontSize float64) *picture {
	p := picture{dpi: dpi, padding: padding, fontSize: fontSize, spacing: 1, font: DefaultFont}
	p.rgba = png
	p.c = freetype.NewContext()
	p.c.SetDPI(p.dpi)
	p.c.SetClip(p.rgba.Bounds())
	p.c.SetDst(p.rgba)
	p.c.SetHinting(font.HintingFull)
	p.c.SetFont(p.font)
	p.c.SetFontSize(p.fontSize)

	p.pt = freetype.Pt(padding, p.c.PointToFixed(fontSize).Round())
	return &p
}

func (p *picture) GeneratePicture() []byte {
	return saveImage(p.rgba)
}

func saveImage(rgba *image.RGBA) []byte {
	b := bytes.NewBuffer(nil)
	bf := bufio.NewWriter(b)
	if err := png.Encode(bf, rgba); err != nil {
		return nil
	}
	if err := bf.Flush(); err != nil {
		return nil
	}
	return b.Bytes()
}

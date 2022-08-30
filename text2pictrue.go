package text2picture

import (
	"bufio"
	"bytes"
	"image"
	"image/png"

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

	face font.Face

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
	p.pt.X = fixed.Int26_6(x) << 6
	p.pt.Y = fixed.Int26_6(y) << 6
}

func (p *picture) PointOffset(x, y float64) {
	p.pt.X += fixed.Int26_6(x) << 6
	p.pt.Y += fixed.Int26_6(y) << 6
}

func (p *picture) SetSpacing(spacing float64) {
	p.spacing = spacing
}

func (p *picture) SetFontSize(fontSize float64) {
	p.c.SetFontSize(fontSize)
	p.face = truetype.NewFace(p.font, &truetype.Options{Size: fontSize, DPI: p.dpi})
	p.fontSize = fontSize
}

func (p *picture) SetPadding(padding int) {
	p.padding = padding
}

func (p *picture) SetFont(font *truetype.Font) {
	p.c.SetFont(font)
	p.face = truetype.NewFace(font, &truetype.Options{Size: p.fontSize, DPI: p.dpi})
	p.font = font
}

func (p *picture) GetRGBA() *image.RGBA {
	return p.rgba
}

func NewPictureWithBackGround(png *image.RGBA, dpi float64, padding int, fontSize float64) *picture {
	p := picture{dpi: dpi, padding: padding, fontSize: fontSize, spacing: 1, font: defaultFont}
	p.rgba = png
	p.c = freetype.NewContext()
	p.c.SetDPI(p.dpi)
	p.c.SetClip(p.rgba.Bounds())
	p.c.SetDst(p.rgba)
	p.c.SetHinting(font.HintingFull)
	p.c.SetFont(p.font)
	p.c.SetFontSize(p.fontSize)

	p.face = truetype.NewFace(p.font, &truetype.Options{Size: p.fontSize, DPI: p.dpi})

	p.pt = freetype.Pt(padding, p.c.PointToFixed(fontSize).Round())
	return &p
}

func (p *picture) GeneratePicture() *bytes.Buffer {
	return saveImage(p.rgba)
}

func saveImage(rgba *image.RGBA) *bytes.Buffer {
	b := bytes.NewBuffer(nil)
	bf := bufio.NewWriter(b)
	if err := png.Encode(bf, rgba); err != nil {
		return nil
	}
	if err := bf.Flush(); err != nil {
		return nil
	}
	return b
}

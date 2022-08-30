package text2picture

import (
	"bufio"
	"bytes"
	"image"
	"image/draw"
	"image/png"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type autoPicture struct {
	width int
	dpi   float64

	font *truetype.Font

	// font size in points
	fontSize float64

	face font.Face

	// text left and right padding
	padding int

	rgba *image.RGBA

	pt fixed.Point26_6
	c  *freetype.Context
}

func AutoNewPicture(width int, dpi float64, padding int, fontSize float64) *autoPicture {
	p := autoPicture{width: width, dpi: dpi, padding: padding, fontSize: fontSize, font: defaultFont}
	p.rgba = NewColorPicture(width, fixed.Int26_6(fontSize*dpi*(64.0/72.0)*1.2).Round(), Transparent)
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

func (p *autoPicture) NextLineDistance() int {
	return p.c.PointToFixed(p.fontSize).Round()
}

func (p *autoPicture) DrawWithWhite(text string) *autoPicture {
	// font color
	p.c.SetSrc(Black)

	p.handleText(&text)

	return p
}

func (p *autoPicture) handleText(text *string) {
	for _, x := range *text {
		w, _ := p.face.GlyphAdvance(x)
		if x == '\n' {
			p.newline()
			continue
		} else if x == '\t' {
			x = ' '
		} else if p.font.Index(x) == 0 {
			continue
		} else if p.pt.X.Round()+w.Round() > p.rgba.Bounds().Max.X-p.padding {
			p.newline()
		}

		p.pt, _ = p.c.DrawString(string(x), p.pt)
	}
}

func (p *autoPicture) newline() {
	if p.pt.Y.Round()+p.NextLineDistance() > p.rgba.Bounds().Max.Y {
		p.addnewline()
	}
	p.pt.X = fixed.Int26_6(p.padding) << 6
	p.pt.Y += p.c.PointToFixed(p.fontSize)
}

func (p *autoPicture) addnewline() {
	rgba := NewColorPicture(p.width, p.pt.Y.Round()+fixed.Int26_6(p.fontSize*p.dpi*(64.0/72.0)*1.2).Round(), Transparent)
	draw.Draw(rgba, p.rgba.Rect, p.rgba, p.rgba.Rect.Min, draw.Src)
	p.rgba = rgba

	p.c.SetClip(p.rgba.Bounds())
	p.c.SetDst(p.rgba)
}

func (p *autoPicture) GeneratePicture() *bytes.Buffer {
	src := NewColorPicture(p.rgba.Bounds().Max.X, p.rgba.Bounds().Max.Y, image.White)
	draw.Draw(src, p.rgba.Rect, p.rgba, p.rgba.Rect.Min, draw.Over)
	b := bytes.NewBuffer(nil)
	bf := bufio.NewWriter(b)
	if err := png.Encode(bf, src); err != nil {
		return nil
	}
	if err := bf.Flush(); err != nil {
		return nil
	}
	return b
}

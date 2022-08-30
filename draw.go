package text2picture

import (
	"image"
	"image/color"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/math/fixed"
)

func (p *picture) DrawWithColor(color color.RGBA, text string) *picture {
	// font color
	p.c.SetSrc(image.NewUniform(color))

	face := truetype.NewFace(p.font, &truetype.Options{Size: p.fontSize, DPI: p.dpi})

	for _, x := range text {
		w, _ := face.GlyphAdvance(x)
		if x == '\n' {
			p.newline(&p.pt, p.c)
			continue
		} else if x == '\t' {
			x = ' '
		} else if p.font.Index(x) == 0 {
			continue
		} else if p.pt.X.Round()+w.Round() > p.rgba.Bounds().Max.X-p.padding {
			p.newline(&p.pt, p.c)
		}

		p.pt, _ = p.c.DrawString(string(x), p.pt)
	}
	return p
}

func (p *picture) DrawWithBlack(text string) *picture {
	// font color
	p.c.SetSrc(image.Black)

	face := truetype.NewFace(p.font, &truetype.Options{Size: p.fontSize, DPI: p.dpi})

	for _, x := range text {
		w, _ := face.GlyphAdvance(x)
		if x == '\n' {
			p.newline(&p.pt, p.c)
			continue
		} else if x == '\t' {
			x = ' '
		} else if p.font.Index(x) == 0 {
			continue
		} else if p.pt.X.Round()+w.Round() > p.rgba.Bounds().Max.X-p.padding {
			p.newline(&p.pt, p.c)
		}

		p.pt, _ = p.c.DrawString(string(x), p.pt)
	}
	return p
}

func (p *picture) DrawWithWhite(text string) *picture {
	// font color
	p.c.SetSrc(image.White)

	face := truetype.NewFace(p.font, &truetype.Options{Size: p.fontSize, DPI: p.dpi})

	for _, x := range text {
		w, _ := face.GlyphAdvance(x)
		if x == '\n' {
			p.newline(&p.pt, p.c)
			continue
		} else if x == '\t' {
			x = ' '
		} else if p.font.Index(x) == 0 {
			continue
		} else if p.pt.X.Round()+w.Round() > p.rgba.Bounds().Max.X-p.padding {
			p.newline(&p.pt, p.c)
		}

		p.pt, _ = p.c.DrawString(string(x), p.pt)
	}
	return p
}

func (p *picture) newline(pt *fixed.Point26_6, c *freetype.Context) {
	pt.X = fixed.Int26_6(p.padding) << 6
	pt.Y += c.PointToFixed(p.fontSize * p.spacing)
}

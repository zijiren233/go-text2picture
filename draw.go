package text2picture

import (
	"image"
	"image/color"

	"golang.org/x/image/math/fixed"
)

func (p *picture) DrawWithColor(color color.Color, text string) *picture {
	// font color
	p.c.SetSrc(image.NewUniform(color))

	p.handleText(&text)

	return p
}

func (p *picture) DrawWithBlack(text string) *picture {
	// font color
	p.c.SetSrc(Black)

	p.handleText(&text)

	return p
}

func (p *picture) DrawWithWhite(text string) *picture {
	// font color
	p.c.SetSrc(White)

	p.handleText(&text)

	return p
}

func (p *picture) handleText(text *string) {
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

func (p *picture) newline() {
	p.pt.X = fixed.Int26_6(p.padding) << 6
	p.pt.Y += p.c.PointToFixed(p.fontSize)
}

func (p *picture) NextLineDistance() int {
	return p.c.PointToFixed(p.fontSize).Round()
}

package text2picture

import (
	"bytes"
	_ "embed"
	"io"

	"github.com/golang/freetype/truetype"
)

//go:embed font.ttf
var fontFile []byte

var defaultFont, _ = ReadFont(bytes.NewReader(fontFile))

func SetDefaultFont(font *truetype.Font) {
	defaultFont = font
}

func ReadFont(fontfile io.Reader) (*truetype.Font, error) {
	b, err := io.ReadAll(fontfile)
	if err != nil {
		return nil, err
	}
	f, err := truetype.Parse(b)
	if err != nil {
		return nil, err
	}

	return f, nil
}

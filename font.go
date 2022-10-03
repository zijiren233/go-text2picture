package text2picture

import (
	"bytes"
	_ "embed"
	"io"
	"os"

	"github.com/golang/freetype/truetype"
)

//go:embed font.ttf
var fontFile []byte

var defaultFont, _ = ReadFont(bytes.NewReader(fontFile))

// Load local fonts
func LoadFont(fontPath string) (*truetype.Font, error) {
	fontfile, err := os.Open(fontPath)
	if err != nil {
		return nil, err
	}
	defer fontfile.Close()
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

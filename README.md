# go-text2picture

```go
package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"io"
	"os"

	"github.com/zijiren233/go-text2picture"
)

func main() {
	pic := text2picture.NewPictureWithBackGround(text2picture.NewWhiteBackGround(500, 100), 320, 10, 5)
	pic.DrowWithBlack("test 123\n4321...")
	pic.DrowWithColor(&image.Uniform{color.Gray16{0x60f2}}, "test color ... ")
	pic.PointOffset(0, 2000)
	pic.DrowWithBlack("test offset\n")
	f, err := os.OpenFile("./test.png", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, bytes.NewReader(pic.GeneratePicture()))
}
```
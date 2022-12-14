# go-text2picture

```go
package main

import (
	"fmt"
	"image/color"
	"io"
	"os"

	"github.com/zijiren233/go-text2picture"
)

func main() {
	pic := text2picture.NewPictureWithBackGround(text2picture.NewColorPicture(500, 100, text2picture.White), 320, 10, 5)
	pic.DrawWithBlack("test 123\n4321...")
	pic.DrawWithColor(text2picture.Red, "test color ... ")
	pic.PointOffset(0, float64(pic.NextLineDistance()))
	pic.DrawWithBlack("test offset")
	pic.DrawWithBlack("   test auto newline")
	f, err := os.OpenFile("./test.png", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, pic.GeneratePicture())
}
```

<img src="https://github.com/zijiren233/go-text2picture/blob/main/example/example.png" />
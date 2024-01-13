package modul

import (
	"image/png"
	"os"
	"time"
	"image"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/gofont/goregular"
)

//####GAME THINGS########

// napis
func Przypal(win *pixelgl.Window, content string, done chan<- bool) {

	ttfFont, _ := truetype.Parse(goregular.TTF)
	fontFace := truetype.NewFace(ttfFont, &truetype.Options{
		Size: 32,
	})

	txt := text.New(pixel.V(300, 200), text.NewAtlas(fontFace, text.ASCII))
	txt.Color = colornames.White
	txt.WriteString(content)
	txt.Draw(win, pixel.IM)
	win.Update()
	time.Sleep(2 * time.Second)
	done <- true
}

//otwieranie obrazkow
func OpeningFile(path string) image.Image {

	File, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer File.Close()

	img, err := png.Decode(File)
	if err != nil {
		panic(err)
	}
	return img
}
//inicjalizowanie obrazkow
func FileInit(img image.Image) *pixel.Sprite {
	bounds := img.Bounds()
	frame := pixel.R(
		float64(bounds.Min.X), float64(bounds.Min.Y),
		float64(bounds.Max.X), float64(bounds.Max.Y),
	)
	sprite := pixel.NewSprite(pixel.PictureDataFromImage(img), frame)
	return sprite
}

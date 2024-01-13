package modul

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// #####Bird#####
// BIRD Movement
// ruch
func Movement(win *pixelgl.Window, birdPosition pixel.Vec, vector, velacity float64) pixel.Vec {
	boost := 0.0
	if birdPosition.X >= 1000.0 {
		boost = 20.0
	}

	if win.Pressed(pixelgl.KeySpace) {
		birdPosition.X += vector
		birdPosition.Y += velacity + float64(boost)
	} else {
		birdPosition.X += vector
		birdPosition.Y -= 10
	}

	return birdPosition
}

// BIRD DIZAJN
// jakie zdjecie ptaka ma pokazywac
func BirdPhoto(win *pixelgl.Window, birdPosition pixel.Vec, vector float64, bird, bird2, goralewo, goraprawo *pixel.Sprite) {
	if vector > 0 {
		if win.Pressed(pixelgl.KeySpace) {
			goraprawo.Draw(win, pixel.IM.Moved(birdPosition))
		} else {
			bird.Draw(win, pixel.IM.Moved(birdPosition))
		}
	} else if vector < 0 {
		if win.Pressed(pixelgl.KeySpace) {
			goralewo.Draw(win, pixel.IM.Moved(birdPosition))
		} else {
			bird2.Draw(win, pixel.IM.Moved(birdPosition))
		}
	}
}

// zmiana kierunku
func Positioning(birdPosition pixel.Vec, vector float64) float64 {

	if birdPosition.X >= 1000.0 {
		vector *= -1
	} else if birdPosition.X <= 0.0 {
		vector *= -1
	}
	return vector

}

package main

import (
	"Gierka/modul"
	"Gierka/modul2"
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	game()

}

func main() {
	pixelgl.Run(run)
}
func game() {

	cfg := pixelgl.WindowConfig{
		Title:  "DONT TOUCH THE SPIKES",
		Bounds: pixel.R(0, 0, 1000, 1000),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	//ladowanie obrazkow
	img2 := modul.OpeningFile("birdleft.png")
	img := modul.OpeningFile("BirdUp.png")
	img3 := modul.OpeningFile("bubbles.png")
	img4 := modul.OpeningFile("goralewo.png")
	img5 := modul.OpeningFile("goraprawo.png")

	// obrazki zrobione
	bird := modul.FileInit(img)
	bird2 := modul.FileInit(img2)
	tlo := modul.FileInit(img3)
	goralewo := modul.FileInit(img4)
	goraprawo := modul.FileInit(img5)

	//pozycje ptaszka i jak ma latac
	x := 500.0
	y := 500.0
	birdPosition := pixel.V(x, y)
	vector := 10.0 //predkosc
	skik := 10.0   //skik
	//kolor
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1, 0, 1)
	done := make(chan bool) // to do napisu kanal zeby czekal

	//struktury kordow z kolcami lewo prawo do tablicy
	var lewo []modul2.Spike
	lewo = modul.KolecLewo((lewo))
	var prawo []modul2.Spike
	prawo = modul.KolecPrawo((prawo))

	trudnosc := 3      //trudnosc bokow ile naraz ma byc
	czyByltam := false //sprawdzanie czy ptak juz byl na sciance zeby moc usunac boczne kolce
	odbicie := 0       //odbicie od sciany
	czyKoniec := false //sprawdzanie czy juz uderzyl

	for !win.Closed() {
		win.Clear(pixel.RGB(0, 0, 0)) // czyszczenie okna
		tlo.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
		modul.SpikeLines(imd) // gora dol kolce

		//tworzenie bocznych kolcow
		if (birdPosition.X != 0.0 && birdPosition.X != 1000.0) && czyByltam == false {
			for k := 0; k < trudnosc; k++ {
				modul.StronaKolca(imd, birdPosition, lewo, prawo)
			}
			czyByltam = true
		} else if (birdPosition.X <= 0.0 || birdPosition.X >= 1000.0) && czyByltam == true {
			imd.Clear()
			modul.SpikeLines(imd)
			imd.Draw(win)
			win.Update()
			czyByltam = false
			odbicie++

		}

		vector = modul.Positioning(birdPosition, vector) // kierunek poruszania po X

		birdPosition = modul.Movement(win, birdPosition, vector, skik)
		// w ktora ma patrzec strone
		go modul.BirdPhoto(win, birdPosition, vector, bird, bird2, goralewo, goraprawo)
		// uderzenie
		if modul.CzyPrzypal(birdPosition) {
			go modul.Przypal(win, "YOU TOUCH THE SPIKE!", done)
			break
		}

		if odbicie > 0 {
			for _, sideSpike := range modul2.SideSpikes {
				if modul.CzyUderzyl(birdPosition, sideSpike) {
					go modul.Przypal(win, "YOU TOUCH THE SPIKE!", done)
					czyKoniec = true

					fmt.Println("SPIKE", sideSpike)
				}
			}
		}
		if czyKoniec == true {
			break
		}

		imd.Draw(win)
		win.Update()
		modul.Usuwanie(birdPosition)
		// Usuwanie poprzednich kolcow na boku
	}
	<-done // do napisu zeby czekal na niego 'przypal'
}

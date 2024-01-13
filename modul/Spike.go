package modul

import (
	"math/rand"
	"Gierka/modul2"
	"time"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

// WYGLAD
// gora dol kolce
func SpikeLines(imd *imdraw.IMDraw) {

	X1 := 0.0
	X2 := 100.0
	X3 := 50.0
	Y1 := 0.0
	Y2 := 0.0
	Y3 := 50.0
	tX1 := 0.0
	tX2 := 100.0
	tX3 := 50.0
	tY1 := 1000.0
	tY2 := 1000.0
	tY3 := 950.0
	for i := 0; i < 10; i++ {
		Kolec(imd, X1, X2, X3, Y1, Y2, Y3)
		Kolec(imd, tX1, tX2, tX3, tY1, tY2, tY3)
		X1 += 100
		X2 += 100
		X3 += 100
		tX1 += 100
		tX2 += 100
		tX3 += 100

	}

}

// #####Spikes#####

// kolce boczne
func KolecLewo(lewo []modul2.Spike) []modul2.Spike {
	spike := modul2.Spike{
		X1: 0.0,
		X2: 50.0,
		X3: 0.0,
		Y1: 0.0,
		Y2: 50,
		Y3: 100,
	}
	for i := 0; i < 10; i++ {

		lewo = append(lewo, spike)
		spike.Y1 += 100
		spike.Y2 += 100
		spike.Y3 += 100
	}
	return lewo
}
func KolecPrawo(prawo []modul2.Spike) []modul2.Spike {
	spike := modul2.Spike{
		X1: 1000.0,
		X2: 950.0,
		X3: 1000.0,
		Y1: 0.0,
		Y2: 50,
		Y3: 100,
	}
	for i := 0; i < 10; i++ {

		prawo = append(prawo, spike)
		spike.Y1 += 100
		spike.Y2 += 100
		spike.Y3 += 100
	}
	return prawo
}

// rysowanie kolca
func Kolec(imd *imdraw.IMDraw, X1, X2, X3, Y1, Y2, Y3 float64) {

	imd.Push(pixel.V(X1, Y1))
	imd.Push(pixel.V(X2, Y2))
	imd.Push(pixel.V(X3, Y3))
	imd.Polygon(0)

}



// ZACHOWANIE KOLCOW
// z ktorej strony ma sie pojawic Kolec 
func StronaKolca(imd *imdraw.IMDraw, birdPosition pixel.Vec, lewo []modul2.Spike, prawo []modul2.Spike) {
	if birdPosition.X > 500.0 {
		RandomSpike(imd, lewo)
	} else if birdPosition.X < 500.0 {
		RandomSpike(imd, prawo)
	}
}

// losowanie z talicy kordow kolca po boku i wypisanie + do tablicy wrzucamy kordy z wylosowanym kolcem
func RandomSpike(imd *imdraw.IMDraw, strona []modul2.Spike) modul2.Spike {
	rand.Seed(time.Now().UnixNano())

	choice := rand.Intn(len(strona))
	selectedSpike := strona[choice]

	Kolec(imd, selectedSpike.X1, selectedSpike.X2, selectedSpike.X3, selectedSpike.Y1, selectedSpike.Y2, selectedSpike.Y3)

	modul2.SideSpikes = append(modul2.SideSpikes, selectedSpike)

	return selectedSpike
}

// usuwa z listy kolcow kolce ktore byly i zniknely zeby ptak nie dostal za niewinnosc
func Usuwanie(birdPosition pixel.Vec) {
	if birdPosition.X >= 1000.0 || birdPosition.X <= 0.0 {
		for i := 0; i < len(modul2.SideSpikes); i++ {
			modul2.SideSpikes = append(modul2.SideSpikes[:i], modul2.SideSpikes[i+1:]...)
			i -= 1
		}
	}

}

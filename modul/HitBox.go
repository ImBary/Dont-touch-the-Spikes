package modul

import (
	"Gierka/modul2"
	"github.com/faiface/pixel"
)

//######hit box#######

// hitbox bocznych kolcow
func CzyUderzyl(birdPosition pixel.Vec, hit modul2.Spike) bool {
	Xhit := 30.0
	Yhit := 0.0
	if birdPosition.X >= 950.0 {
		if (((hit.X1 - birdPosition.X) <= Xhit) && ((hit.X1 - birdPosition.X) >= Yhit)) && (((hit.Y1 - birdPosition.Y) <= Xhit) && ((hit.Y1 - birdPosition.Y) >= Yhit)) {
			return true
		} else if (((hit.X2 - birdPosition.X) <= Xhit+50.0) && ((hit.X2 - birdPosition.X) >= Yhit)) && (((hit.Y2 - birdPosition.Y) <= Xhit) && ((hit.Y2 - birdPosition.Y) >= Yhit)) {
			return true
		} else if (((hit.X3 - birdPosition.X) <= Xhit) && ((hit.X3 - birdPosition.X) >= Yhit)) && (((hit.Y3 - birdPosition.Y) <= Xhit) && ((hit.Y3 - birdPosition.Y) >= Yhit)) {
			return true
		}
	}
	if birdPosition.X <= 50.0 {
		if (((hit.X1 + birdPosition.X) <= (Xhit)) && ((hit.X1 + birdPosition.X) >= Yhit)) && (((hit.Y1 - birdPosition.Y) <= Xhit) && ((hit.Y1 - birdPosition.Y) >= Yhit)) {
			return true
		} else if (((hit.X2 + birdPosition.X) <= (Xhit+50.0)) && ((hit.X2 + birdPosition.X) >= Yhit)) && (((hit.Y2 - birdPosition.Y) <= Xhit) && ((hit.Y2 - birdPosition.Y) >= Yhit)) {
			return true
		} else if (((hit.X3 + birdPosition.X) <= Xhit) && ((hit.X3 + birdPosition.X) >= Yhit)) && (((hit.Y3 - birdPosition.Y) <= Xhit) && ((hit.Y3 - birdPosition.Y) >= Yhit)) {
			return true
		}
	}
	return false
}

// hitbox gora dol
func CzyPrzypal(birdPosition pixel.Vec) bool {
	if birdPosition.Y >= 910.0 || birdPosition.Y <= 60.0 {
		return true
	}
	return false
}

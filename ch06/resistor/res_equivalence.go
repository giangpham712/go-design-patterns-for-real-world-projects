package resistor

func Rser(resists ...float64) (Rtotal float64) {
	for _, r := range resists {
		Rtotal = Rtotal + r
	}
	return
}

func Rpara(resists ...float64) (Rtotal float64) {
	for _, r := range resists {
		Rtotal = Rtotal + recip(r)
	}
	return
}

package volt

func V(i, r float64) float64 {
	return i * r
}

func Vser(volts ...float64) (Vtotal float64) {
	for _, v := range volts {
		Vtotal = Vtotal + v
	}
	return
}

func Vpi(p, i float64) float64 {
	return p / i
}

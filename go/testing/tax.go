package tax

func Tax(price float64) float64 {
	tax := 0.08

	if price > 1000 {
		return price * tax * 0.9
	}

	return price * tax
}

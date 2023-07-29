package core

import "math/rand"

type Double = float64

func randomDoubleW(min, max Double) Double {
	return rand.Float64()*(max-min) + min
}

func randomDouble() Double {
	return rand.Float64()
}

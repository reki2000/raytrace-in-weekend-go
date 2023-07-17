package core

import "math/rand"

type Color struct {
	R, G, B double
}

func NewColor(r, g, b double) Color {
	return Color{r, g, b}
}

func NewColorRandom(min, max double) Color {
	return Color{rand.Float64()*(max-min) + min,
		rand.Float64()*(max-min) + min,
		rand.Float64()*(max-min) + min}
}

func (c Color) MulVec(x Color) Color {
	return Color{c.R * x.R, c.G * x.G, c.B * x.B}
}

func (c Color) Mul(n double) Color {
	return Color{c.R * n, c.G * n, c.B * n}
}

func (c Color) Add(x Color) Color {
	return Color{c.R + x.R, c.G + x.G, c.B + x.B}
}

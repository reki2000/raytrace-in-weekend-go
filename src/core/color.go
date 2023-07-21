package core

var black = Color{0, 0, 0}
var white = Color{1.0, 1.0, 1.0}

type Color struct {
	R, G, B Double
}

func NewColor(r, g, b Double) Color {
	return Color{r, g, b}
}

func NewColorRandom(min, max Double) Color {
	return Color{
		randomDoubleW(min, max),
		randomDoubleW(min, max),
		randomDoubleW(min, max),
	}
}

func (c Color) MulVec(x Color) Color {
	return Color{c.R * x.R, c.G * x.G, c.B * x.B}
}

func (c Color) Mul(n Double) Color {
	return Color{c.R * n, c.G * n, c.B * n}
}

func (c Color) Add(x Color) Color {
	return Color{c.R + x.R, c.G + x.G, c.B + x.B}
}

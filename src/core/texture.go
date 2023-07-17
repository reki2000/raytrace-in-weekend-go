package core

import "math"

type Texture interface {
	Value(u, v float64, p *Vec3) *Vec3
}

type SolidColor struct {
	Color *Vec3
}

func NewSolidColor(rgb *Vec3) *SolidColor {
	return &SolidColor{rgb}
}

func NewSolidColorRGB(r, g, b double) *SolidColor {
	return &SolidColor{NewVec3(r, g, b)}
}

func (sc *SolidColor) Value(u, v float64, p *Vec3) *Vec3 {
	return sc.Color
}

type CheckerTexture struct {
	odd, even Texture
}

func NewCheckerTexture(odd, even Texture) *CheckerTexture {
	return &CheckerTexture{odd, even}
}

func (ct *CheckerTexture) Value(u, v float64, p *Vec3) *Vec3 {
	sines := math.Sin(10*p.X) * math.Sin(10*p.Y) * math.Sin(10*p.Z)
	if sines < 0 {
		return ct.odd.Value(u, v, p)
	} else {
		return ct.even.Value(u, v, p)
	}
}
